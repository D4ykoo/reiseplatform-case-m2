pub mod model;

use axum::extract::{Query, State};
use axum::Json;
use axum::{http::StatusCode, routing::get, Router};
use chrono::{DateTime, Utc};
use model::Period;
use monitoring_db::model::{CheckoutEvent, HotelEvent, UserEvent};
use monitoring_db::get_connection_pool;

use tower_http::services::ServeDir;

#[tokio::main]
async fn main() {

    let pool = get_connection_pool();

    let app = Router::new()
        .nest_service(
            "/",
            ServeDir::new("../../frontend/test-app/dist/test-app/browser/"),
        )
        .route("/api/v1/user-events", get(get_user_events))
        .route("/api/v1/hotel-events", get(get_hotel_events))
        .route("/api/v1/checkout-events", get(get_checkout_events))
        .with_state(pool);

    // run our app with hyper, listening globally on port 3000
    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}

async fn get_user_events(
    State(pool): State<deadpool_diesel::postgres::Pool>,
    querry: Query<Period>,
) -> Result<Json<Vec<UserEvent>>, (StatusCode, String)> {
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
    State(pool): State<deadpool_diesel::postgres::Pool>,
    querry: Query<Period>,
) -> Result<Json<Vec<HotelEvent>>, (StatusCode, String)> {
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
    State(pool): State<deadpool_diesel::postgres::Pool>,
    querry: Query<Period>,
) -> Result<Json<Vec<CheckoutEvent>>, (StatusCode, String)> {
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

fn internal_error<E>(err: E) -> (StatusCode, String)
where
    E: std::error::Error,
{
    (StatusCode::INTERNAL_SERVER_ERROR, err.to_string())
}
