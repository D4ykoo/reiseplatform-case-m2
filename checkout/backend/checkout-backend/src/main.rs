mod dto;
mod request_helper;
mod models;

use std::env;
use dotenvy::dotenv;

use actix_cors::Cors;
use actix_web::{
    delete, get,
    http::StatusCode,
    post, put,
    web::{self, Data},
    App, HttpRequest, HttpResponse, HttpServer,
};
use checkout_db::PostgresPool;
use message_service::MessageProducer;

use crate::dto::CombinedCartResponse;


#[post("/create")]
async fn create_cart() -> HttpResponse {
    HttpResponse::Ok().json("Hello world!")
}
//     pool: web::Data<PostgresPool>,
//     producer: web::Data<MessageProducer>,
//     new_cart: web::Json<CombinedCartResponse>,
// ) -> HttpResponse {
//     let mut conn = pool.get().expect("Could not connect to db from pool");
//     let res = checkout_db::create_cart(&mut conn, new_cart.into_new_cart());

//     match res {
//         Ok(r) => {
//             producer
//                 .send_message(&format!("Created cart {:?}", new_cart))
//                 .await;
//             return HttpResponse::Ok()
//                 .status(StatusCode::OK)
//                 .json(format!("affected: {}", r));
//         }
//         Err(e) => {
//             producer
//                 .send_message(&format!("Could not create cart {:?}", new_cart))
//                 .await;
//             return HttpResponse::NotFound()
//                 .status(StatusCode::NOT_FOUND)
//                 .json(e.to_string());
//         }
//     }


#[get("/cart/{cart_id}")]
async fn get_cart(
    pool: Data<PostgresPool>,
    producer: web::Data<MessageProducer>,
    req: HttpRequest,
) -> HttpResponse {
    let cart_id: i32 = req.match_info().query("cart_id").parse().unwrap();

    let mut conn = pool.get().expect("Could not connect to db from pool");
    let res = checkout_db::get_cart(&mut conn, &cart_id);
    match res {
        Some(cart) => {
            producer
                .send_message(&format!("Get cart {}", cart_id))
                .await;
            return HttpResponse::Ok().status(StatusCode::OK).json(cart);
        }
        None => {
            producer
                .send_message(&format!("Could not get cart {}", cart_id))
                .await;
            return HttpResponse::NotFound().into();
        }
    }
}

// #[put("/cart/{cart_id}")]
// async fn change_cart(
//     pool: web::Data<PostgresPool>,
//     new_cart: web::Json<dto::CartRequest>,
//     producer: web::Data<MessageProducer>,
//     req: HttpRequest,
// ) -> HttpResponse {
//     let cart_id: i32 = req.match_info().query("cart_id").parse().unwrap();
//     let mut conn = pool.get().expect("Could not connect to db from pool");

//     let res = checkout_db::update_card(&mut conn, &cart_id, new_cart.into_new_cart());

//     match res {
//         Ok(_) => {
//             producer
//                 .send_message(&format!("Updated cart {}", cart_id))
//                 .await;
//             return HttpResponse::Ok().status(StatusCode::OK).finish();
//         }
//         Err(e) => {
//             producer
//                 .send_message(&format!(
//                     "Could not updated cart {} due to error: {}",
//                     cart_id, e
//                 ))
//                 .await;
//             return HttpResponse::InternalServerError()
//                 .status(StatusCode::INTERNAL_SERVER_ERROR)
//                 .json(e);
//         }
//     }
// }

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
    req: HttpRequest,
) -> HttpResponse {
    let user_id: i32 = req.match_info().query("user_id").parse().unwrap();
    let hotel_id: i32 = req.match_info().query("hotel_id").parse().unwrap();
    let travel_id: i32 = req.match_info().query("travel_id").parse().unwrap();

    // actix get cookie
    let cookie = req.cookie("authTravel");

    if cookie.is_none() {
        return HttpResponse::Unauthorized().status(StatusCode::UNAUTHORIZED).finish();
    }
    
    let cookie = cookie.unwrap();

    let mut conn = pool.get().expect("Could not connect to db from pool");

    let mut cart_id = checkout_db::get_cart_id(&mut conn, &user_id);
    
    if !cart_id.is_ok(){
        let _ = checkout_db::create_cart(&mut conn, checkout_db::models::NewCart::create(Some(user_id), None, None));
        cart_id = checkout_db::get_cart_id(&mut conn, &user_id);
    }
    
    let cart_id = cart_id.unwrap();
    
    // cart id is now available, get hotels and travel slices from api
    let hoteltravel = request_helper::get_hotel_travel_request_with_cookie(hotel_id, travel_id, cookie.value()).await.unwrap();
    

    let hotel = hoteltravel.to_db_hotel(cart_id);
    let travel_slice = hoteltravel.to_db_travel_slice(hotel_id);

    let res = checkout_db::add_to_cart(&mut conn, &cart_id,  &hotel, &travel_slice);

    match res {
        Ok(_) => {
            return HttpResponse::Ok().status(StatusCode::OK).finish();
        }
        Err(e) => {
            return HttpResponse::NotFound()
                .status(StatusCode::INTERNAL_SERVER_ERROR)
                .json(e.to_string());
        }
    }
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenv().ok();
    let api_url = env::var("API_URL").expect("API_URL must be set");
    let srpai = api_url.as_str();

    let api_port: u16 = env::var("API_PORT").expect("API_PORT must be set").parse().unwrap();
    let pool = checkout_db::get_pool();
    // let mut producer = MessageProducer { producer: None };
    // let _ = producer.init_message_producer();
    // async send message thath is not blocking  using producer

    // producer.send_message("Starting checkout webserver");
    env_logger::init_from_env(env_logger::Env::new().default_filter_or("info"));
    HttpServer::new(move || {
        let cors = Cors::permissive();

        App::new()
            .wrap(cors)
            .wrap(actix_web::middleware::Logger::default())
            .app_data(web::Data::new(pool.clone()))
            // .app_data(web::Data::new(producer.clone()))
            .service(
                web::scope("/api/v1/cart")
                    .service(get_cart)
                    // .service(change_cart)
                    .service(delete_cart)
                    .service(create_cart)
                    .service(add_to_cart)
                    ,
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
