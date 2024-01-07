mod dto;
mod models;
mod request_helper;

use dotenvy::dotenv;
use std::env;

use actix_cors::Cors;
use actix_web::{
    delete, get,
    http::StatusCode,
    put,
    web::{self, Data},
    App, HttpRequest, HttpResponse, HttpServer,
};
use checkout_db::PostgresPool;
use message_service::MessageProducer;

use crate::dto::CombinedCartResponse;

#[get("/{user_id}")]
async fn get_cart(
    pool: Data<PostgresPool>,
    producer: web::Data<MessageProducer>,
    req: HttpRequest,
) -> HttpResponse {
    let user_id: i32 = req.match_info().query("user_id").parse().unwrap();

    let mut conn = pool.get().expect("Could not connect to db from pool");

    let cart_id = checkout_db::get_cart_id(&mut conn, &user_id);

    if cart_id.is_err() {
        producer
            .send_message(&format!(
                "Could not find cart for user {}",
                user_id
            ))
            .await;

        return HttpResponse::NotFound()
            .status(StatusCode::NOT_FOUND)
            .json("Could not find cart");
    }

    let res = checkout_db::get_cart_content(&mut conn, &cart_id.as_ref().unwrap());

    match res {
        Some(result) => {
            let travel_slice = result.travel_slice.clone();
            let mut combined_cart_response: CombinedCartResponse =
                CombinedCartResponse::from_db_combined_cart(result);

            // may need the index in the future
            for (_, hotel) in combined_cart_response
                .hotel
                .as_mut()
                .unwrap()
                .iter_mut()
                .enumerate()
            {
                hotel.add_travel_slice(travel_slice.clone().unwrap());
            }

            producer
                .send_message(&format!("Get cart {}", cart_id.unwrap()))
                .await;

            return HttpResponse::Ok()
                .status(StatusCode::OK)
                .json(combined_cart_response);
        }
        None => {
            producer
                .send_message(&format!("Could not find cart {}", cart_id.unwrap()))
                .await;

            return HttpResponse::NotFound()
                .status(StatusCode::NOT_FOUND)
                .json("Could not find cart");
        }
    }
}

#[delete("/cart/{cart_id}")]
async fn delete_cart(
    pool: web::Data<PostgresPool>,
    producer: web::Data<MessageProducer>,
    req: HttpRequest,
) -> HttpResponse {
    let cart_id: i32 = req.match_info().query("cart_id").parse().unwrap();
    let mut conn = pool.get().expect("Could not connect to db from pool");

    let res = checkout_db::remove_cart(&mut conn, &cart_id);

    match res {
        Ok(_) => {
            producer
                .send_message(&format!("Deleted cart {}", cart_id))
                .await;
            return HttpResponse::Ok().status(StatusCode::OK).finish();
        }
        Err(e) => {
            producer
                .send_message(&format!("Could not delete cart {}", cart_id))
                .await;
            return HttpResponse::NotFound()
                .status(StatusCode::NOT_FOUND)
                .json(e.to_string());
        }
    }
}

#[put("/addtocart/{user_id}/{hotel_id}/{travel_id}")]
async fn add_to_cart(
    pool: web::Data<PostgresPool>,
    producer: web::Data<MessageProducer>,
    req: HttpRequest,
) -> HttpResponse {
    let user_id: i32 = req.match_info().query("user_id").parse().unwrap();
    let hotel_id: i32 = req.match_info().query("hotel_id").parse().unwrap();
    let travel_id: i32 = req.match_info().query("travel_id").parse().unwrap();

    // actix get cookie
    let cookie = req.cookie("authTravel");

    if cookie.is_none() {
        producer
            .send_message(&format!(
                "Unauthorized user {} with cookie {}",
                user_id,
                &cookie.unwrap()
            ))
            .await;

        return HttpResponse::Unauthorized()
            .status(StatusCode::UNAUTHORIZED)
            .finish();
    }

    let cookie = cookie.unwrap();

    let mut conn = pool.get().expect("Could not connect to db from pool");

    let mut cart_id = checkout_db::get_cart_id(&mut conn, &user_id);

    if cart_id.is_err() {
        let _ = checkout_db::create_cart(
            &mut conn,
            checkout_db::models::NewCart::create(Some(user_id), None, None),
        );
        cart_id = checkout_db::get_cart_id(&mut conn, &user_id);

        producer
            .send_message(&format!(
                "Created cart for user {} with cart id {}",
                user_id,
                &cart_id.as_ref().unwrap()
            ))
            .await;
    }

    let cart_id = cart_id.unwrap();

    // cart id is now available, get hotels and travel slices from api
    let hoteltravel =
        request_helper::get_hotel_travel_request_with_cookie(hotel_id, travel_id, cookie.value())
            .await
            .unwrap();

    let hotel = hoteltravel.to_db_hotel(cart_id);
    let travel_slice = hoteltravel.to_db_travel_slice(hotel_id);

    let res = checkout_db::add_to_cart(&mut conn, &cart_id, &hotel, &travel_slice);

    match res {
        Ok(_) => {
            producer
                .send_message(&format!(
                    "Added hotel {} and travel slice {} to cart {}",
                    hotel_id, travel_id, cart_id
                ))
                .await;
            return HttpResponse::Ok().status(StatusCode::OK).finish();
        }
        Err(e) => {
            producer
                .send_message(&format!(
                    "Could not add hotel {} and travel slice {} to cart {}",
                    hotel_id, travel_id, cart_id
                ))
                .await;

            return HttpResponse::NotFound()
                .status(StatusCode::INTERNAL_SERVER_ERROR)
                .json(e.to_string());
        }
    }
}

#[delete("/entry/{cart_id}/{hotel_id}/{travel_id}")]
async fn delete_cart_entry(
    pool: web::Data<PostgresPool>,
    producer: web::Data<MessageProducer>,
    req: HttpRequest,
) -> HttpResponse {
    let cart_id: i32 = req.match_info().query("cart_id").parse().unwrap();
    let hotel_id: i32 = req.match_info().query("hotel_id").parse().unwrap();
    let travel_id: i32 = req.match_info().query("travel_id").parse().unwrap();

    let mut conn = pool.get().expect("Could not connect to db from pool");

    let res =
        checkout_db::remove_hotel_and_travel_slice(&mut conn, &cart_id, &hotel_id, &travel_id);

    match res {
        Ok(_) => {
            producer
                .send_message(&format!(
                    "Deleted hotel {} and travel slice {} from cart {}",
                    hotel_id, travel_id, cart_id
                ))
                .await;
            return HttpResponse::Ok().status(StatusCode::OK).finish();
        }
        Err(e) => {
            producer
                .send_message(&format!(
                    "Could not delete hotel {} and travel slice {} from cart {}",
                    hotel_id, travel_id, cart_id
                ))
                .await;
            return HttpResponse::NotFound()
                .status(StatusCode::INTERNAL_SERVER_ERROR)
                .json(e.to_string());
        }
    }
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenv().ok();
    env_logger::init_from_env(env_logger::Env::new().default_filter_or("info"));

    let api_url = env::var("API_URL").expect("API_URL must be set");
    let srpai = api_url.as_str();

    let api_port: u16 = env::var("API_PORT")
        .expect("API_PORT must be set")
        .parse()
        .unwrap();

    let pool = checkout_db::get_pool();
    let migrations_res = checkout_db::init_migrations(&mut pool.get().unwrap());

    if migrations_res.is_err() {
        panic!("Could not run migrations");
    }

    let mut producer = MessageProducer { producer: None };
    let kafka_url = env::var("KAFKA_URL").unwrap_or("".into());

    println!("Kafka url: {}", kafka_url.as_str());

    producer.init_message_producer(kafka_url.as_str());

    producer.send_message("Starting checkout webserver").await;

    HttpServer::new(move || {
        let cors = Cors::permissive();

        App::new()
            .wrap(cors)
            .wrap(actix_web::middleware::Logger::default())
            .app_data(web::Data::new(pool.clone()))
            .app_data(web::Data::new(producer.clone()))
            .service(
                web::scope("/api/v1/cart")
                    .service(get_cart)
                    .service(delete_cart)
                    .service(add_to_cart)
                    .service(delete_cart_entry),
                // if more versions of the api are needed, they can be added here
                // web::scope("/api/v2/checkout")
                //     .service(get_cart)
                //     .service(change_cart)
                //     .service(delete_cart)
                //     .service(create_cart),
            )
    })
    .bind((srpai, api_port))?
    .run()
    .await
}
