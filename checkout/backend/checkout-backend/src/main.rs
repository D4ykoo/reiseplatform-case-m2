pub mod dto;

use std::fmt::format;

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

#[post("/create")]
async fn create_cart(
    pool: web::Data<PostgresPool>,
    producer: web::Data<MessageProducer>,
    new_cart: web::Json<dto::CartRequest>,
) -> HttpResponse {
    let mut conn = pool.get().expect("Could not connect to db from pool");
    let res = checkout_db::create_cart(&mut conn, new_cart.into_new_cart());

    match res {
        Ok(r) => {
            producer
                .send_message(&format!("Created cart {:?}", new_cart))
                .await;
            return HttpResponse::Ok()
                .status(StatusCode::OK)
                .json(format!("affected: {}", r));
        }
        Err(e) => {
            producer
                .send_message(&format!("Could not create cart {:?}", new_cart))
                .await;
            return HttpResponse::NotFound()
                .status(StatusCode::NOT_FOUND)
                .json(e.to_string());
        }
    }
}

#[get("/get/{cart_id}")]
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

#[put("/update/{cart_id}")]
async fn change_cart(
    pool: web::Data<PostgresPool>,
    new_cart: web::Json<dto::CartRequest>,
    producer: web::Data<MessageProducer>,
    req: HttpRequest,
) -> HttpResponse {
    let cart_id: i32 = req.match_info().query("cart_id").parse().unwrap();
    let mut conn = pool.get().expect("Could not connect to db from pool");

    let res = checkout_db::update_card(&mut conn, &cart_id, new_cart.into_new_cart());

    match res {
        Ok(_) => {
            producer
                .send_message(&format!("Updated cart {}", cart_id))
                .await;
            return HttpResponse::Ok().status(StatusCode::OK).finish();
        }
        Err(e) => {
            producer
                .send_message(&format!(
                    "Could not updated cart {} due to error: {}",
                    cart_id, e
                ))
                .await;
            return HttpResponse::InternalServerError()
                .status(StatusCode::INTERNAL_SERVER_ERROR)
                .json(e);
        }
    }
}

#[delete("/delete/{cart_id}")]
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

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let pool = checkout_db::get_pool();
    let mut producer = MessageProducer { producer: None };
    producer.init_message_producer();
    producer.send_message("Starting checkout webserver").await;
    HttpServer::new(move || {
        let cors = Cors::permissive();

        App::new()
            .wrap(cors)
            .app_data(web::Data::new(pool.clone()))
            .app_data(web::Data::new(producer.clone()))
            .service(
                web::scope("/c")
                    .service(get_cart)
                    .service(change_cart)
                    .service(delete_cart)
                    .service(create_cart),
            )
    })
    .bind(("127.0.0.1", 8071))?
    .run()
    .await
}
