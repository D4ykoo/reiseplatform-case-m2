pub mod dto;

use actix_cors::Cors;
use actix_web::{
    get,
    http::StatusCode,
    put,
    web::{self, Data},
    App, HttpRequest, HttpResponse, HttpServer, delete, post,
};
use checkout_db::PostgresPool;

#[post("/create")]
async fn create_cart(
    pool: web::Data<PostgresPool>,
    new_cart: web::Json<dto::CartRequest>,
) -> HttpResponse {
    let mut conn = pool.get().expect("Could not connect to db from pool");
    let res = checkout_db::create_cart(&mut conn, new_cart.into_new_cart());

    match res {
        Ok(r) => {
            return HttpResponse::Ok()
                .status(StatusCode::OK)
                .json(format!("affected: {}", r))
        }
        Err(e) => {
            return HttpResponse::NotFound()
                .status(StatusCode::NOT_FOUND)
                .json(e.to_string())
        }
    }
}


#[get("/get/{cart_id}")]
async fn get_cart(pool: Data<PostgresPool>, req: HttpRequest) -> HttpResponse {
    let cart_id: i32 = req.match_info().query("cart_id").parse().unwrap();

    let mut conn = pool.get().expect("Could not connect to db from pool");
    let res = checkout_db::get_cart(&mut conn, &cart_id);
    match res {
        Some(cart) => return HttpResponse::Ok().status(StatusCode::OK).json(cart),
        None => return HttpResponse::NotFound().into(),
    }
}

#[put("/update/{cart_id}")]
async fn change_cart(
    pool: web::Data<PostgresPool>,
    new_cart: web::Json<dto::CartRequest>,
    req: HttpRequest,
) -> HttpResponse {
    let cart_id: i32 = req.match_info().query("cart_id").parse().unwrap();
    let mut conn = pool.get().expect("Could not connect to db from pool");

    let res = checkout_db::update_card(&mut conn, &cart_id, new_cart.into_new_cart());

    match res {
        Ok(_) => return HttpResponse::Ok().status(StatusCode::OK).finish(),
        Err(e) => {
            return HttpResponse::InternalServerError()
                .status(StatusCode::INTERNAL_SERVER_ERROR)
                .json(e)
        }
    }
}

#[delete("/delete/{cart_id}")]
async fn delete_cart(
    pool: web::Data<PostgresPool>,
    req: HttpRequest,
) -> HttpResponse {
    let cart_id: i32 = req.match_info().query("cart_id").parse().unwrap();
    let mut conn = pool.get().expect("Could not connect to db from pool");

    let res = checkout_db::remove_cart(&mut conn, &cart_id);

    match res {
        Ok(_) => return HttpResponse::Ok().status(StatusCode::OK).finish(),
        Err(e) => {
            return HttpResponse::NotFound()
                .status(StatusCode::NOT_FOUND)
                .json(e.to_string())
        }
    }
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let pool = checkout_db::get_pool();
    HttpServer::new(move || {
        let cors = Cors::permissive();

        App::new()
            .wrap(cors)
            .app_data(web::Data::new(pool.clone()))
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
