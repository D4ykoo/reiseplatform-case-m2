pub mod model;

use axum::extract::{Query, State};
use axum::Json;
use axum::{http::StatusCode, routing::get, Router};
use chrono::{DateTime, Utc};
use dotenvy::dotenv;
use kafka_consumer::{create_consumer, subscribe, KafkaMsg};
use model::Period;
use monitoring_db::model::{
    CheckoutEvent, HotelEvent, NewCheckoutEvent, NewHotelEvent, NewUserEvent, UserEvent,
};
use monitoring_db::{
    add_checkout_event, add_hotel_event, add_user_event, establish_connection, get_connection_pool,
    run_migrations,
};

use jwt_auth::validate_jwt;
use std::env;
use std::sync::mpsc::channel;
use tower_cookies::{CookieManagerLayer, Cookies};
use tower_http::services::ServeDir;
use tracing::warn;
use tower_http::cors::{Any, CorsLayer};

use crate::model::TokenError;

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt::init();
    let cors = CorsLayer::new().allow_origin(Any);
    let (tx, rx) = channel();

    let pool = get_connection_pool();

    // run the migrations on server startup
    {
        let conn = pool.get().await.unwrap();
        conn.interact(move |pg| run_migrations(pg))
            .await
            .unwrap()
            .unwrap();

        dotenv().ok();
        let server = env::var("BROKERS").unwrap_or("".into());

        tokio::spawn(async move {
            let topics = ["usermanagement", "travelmanagement", "checkout"];
            let consumer =
                create_consumer(server.as_str(), &topics).expect("Canot create Consumer");
            subscribe(tx, consumer).await;
        });

        tokio::spawn(async move {
            let conn = &mut establish_connection();
            loop {
                let event = rx.recv().unwrap_or(KafkaMsg {
                    topic: "".to_string(),
                    payload: "".to_string(),
                });
                match event.topic.as_str() {
                    "usermanagement" => {
                        let event: NewUserEvent = serde_json::from_str(&event.payload).unwrap();
                        add_user_event(conn, event).expect("Error");
                    }
                    "travelmanagement" => {
                        let event: NewHotelEvent = serde_json::from_str(&event.payload).unwrap();
                        add_hotel_event(conn, event).expect("Error");
                    }
                    "checkout" => {
                        let event: NewCheckoutEvent = serde_json::from_str(&event.payload).unwrap();
                        add_checkout_event(conn, event).expect("Error");
                    }
                    &_ => warn!("Received not supported Topic {:?}", event),
                }
            }
        });
    }

    let app = Router::new()
        .nest_service(
            "/",
            ServeDir::new("../../frontend/test-app/dist/test-app/browser/"),
        )
        .route("/api/v1/user-events", get(get_user_events))
        .route("/api/v1/hotel-events", get(get_hotel_events))
        .route("/api/v1/checkout-events", get(get_checkout_events))
        .layer(CookieManagerLayer::new())
        .layer(cors)
        .with_state(pool);

    // run our app with hyper, listening globally on port 3000
    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}

async fn get_user_events(
    cookies: Cookies,
    State(pool): State<deadpool_diesel::postgres::Pool>,
    querry: Query<Period>,
) -> Result<Json<Vec<UserEvent>>, (StatusCode, String)> {
    //validate_auth(cookies)?;

    let period = querry.0;

    let from = parse_time_querry(&period.from, "1970-01-01T01:00:01+01:00");
    let _to = parse_time_querry(&period.to, "2100-01-01T01:00:01+01:00");

    let conn = pool.get().await.map_err(internal_error)?;

    let res = conn
        .interact(move |pg| monitoring_db::get_user_events(pg, &from))
        .await
        .map_err(internal_error)?;

    Ok(Json(res.unwrap()))
}

async fn get_hotel_events(
    cookies: Cookies,
    State(pool): State<deadpool_diesel::postgres::Pool>,
    querry: Query<Period>,
) -> Result<Json<Vec<HotelEvent>>, (StatusCode, String)> {
    validate_auth(cookies)?;

    let period = querry.0;

    let from = parse_time_querry(&period.from, "1970-01-01T01:00:01+01:00");
    let _to = parse_time_querry(&period.to, "2100-01-01T01:00:01+01:00");

    let conn = pool.get().await.map_err(internal_error)?;

    let res = conn
        .interact(move |pg| monitoring_db::get_hotel_events(pg, &from))
        .await
        .map_err(internal_error)?;

    Ok(Json(res.unwrap()))
}

async fn get_checkout_events(
    cookies: Cookies,
    State(pool): State<deadpool_diesel::postgres::Pool>,
    querry: Query<Period>,
) -> Result<Json<Vec<CheckoutEvent>>, (StatusCode, String)> {
    validate_auth(cookies)?;

    let period = querry.0;

    let from = parse_time_querry(&period.from, "1970-01-01T01:00:01+01:00");
    let _to = parse_time_querry(&period.to, "2100-01-01T01:00:01+01:00");

    let conn = pool.get().await.map_err(internal_error)?;

    let res = conn
        .interact(move |pg| monitoring_db::get_checkout_events(pg, &from))
        .await
        .map_err(internal_error)?;

    Ok(Json(res.unwrap()))
}

fn validate_auth(cookies: Cookies) -> Result<bool, (StatusCode, String)> {
    let cookie = cookies.get("authTravel");

    let token = cookie.ok_or(TokenError::new("Missing Token"));

    if token.is_err() {
        return Err(auth_error(token.err().unwrap()));
    }

    let res = validate_jwt(token.unwrap().value());
    if res.is_ok() {
        Ok(true)
    } else {
        Err(auth_error(res.err().unwrap()))
    }
}

fn parse_time_querry(time: &Option<String>, default: &str) -> DateTime<Utc> {
    match time {
        None => chrono::DateTime::parse_from_rfc3339(default)
            .unwrap()
            .into(),
        Some(s) => {
            let result = chrono::DateTime::parse_from_rfc3339(s)
                .unwrap_or(chrono::DateTime::parse_from_rfc3339(default).unwrap());
            result.into()
        }
    }
}

fn auth_error<E>(err: E) -> (StatusCode, String)
where
    E: std::error::Error,
{
    (StatusCode::UNAUTHORIZED, err.to_string())
}

fn internal_error<E>(err: E) -> (StatusCode, String)
where
    E: std::error::Error,
{
    (StatusCode::INTERNAL_SERVER_ERROR, err.to_string())
}
