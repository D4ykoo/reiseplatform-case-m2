pub mod model;
pub mod schema;

use chrono::{DateTime, Utc};
use diesel::prelude::*;
use diesel::r2d2::ConnectionManager;
use diesel::r2d2::Pool;
use diesel::PgConnection;
use dotenvy::dotenv;
use model::CheckoutEvent;
use model::HotelEvent;
use model::NewCheckoutEvent;
use model::NewHotelEvent;
use model::NewUserEvent;
use model::UserEvent;
use schema::checkout_event;
use schema::hotel_event;
use schema::user_event;
use std::env;

pub fn get_connection_pool() -> Pool<ConnectionManager<PgConnection>> {
    dotenv().ok();
    let database_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set");

    let manager = ConnectionManager::<PgConnection>::new(database_url);

    Pool::builder()
        .test_on_check_out(true)
        .build(manager)
        .expect("Could not build connection pool")
}

pub fn establish_connection() -> PgConnection {
    dotenv().ok();

    let database_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set");
    PgConnection::establish(&database_url)
        .unwrap_or_else(|_| panic!("Error connecting to {database_url}"))
}

pub fn add_user_event(
    conn: &mut PgConnection,
    new_event: NewUserEvent,
) -> Result<usize, diesel::result::Error> {
    let inserted_rows = diesel::insert_into(user_event::table)
        .values(&new_event)
        .execute(conn)
        .expect("Error can not add Event");

    Ok(inserted_rows)
}

pub fn add_hotel_event(
    conn: &mut PgConnection,
    new_event: NewHotelEvent,
) -> Result<usize, diesel::result::Error> {
    let inserted_rows = diesel::insert_into(hotel_event::table)
        .values(&new_event)
        .execute(conn)
        .expect("Error can not add Event");

    Ok(inserted_rows)
}

pub fn add_checkout_event(
    conn: &mut PgConnection,
    new_event: NewCheckoutEvent,
) -> Result<usize, diesel::result::Error> {
    let inserted_rows = diesel::insert_into(checkout_event::table)
        .values(&new_event)
        .execute(conn)
        .expect("Error can not add Event");

    Ok(inserted_rows)
}

pub fn get_user_events(
    conn: &mut PgConnection,
    from: &Option<DateTime<Utc>>,
) -> Result<Vec<UserEvent>, diesel::result::Error> {
    let default = "2000-01-01T00:00:00+00:00";
    let default_time = DateTime::parse_from_rfc3339(default).unwrap().into();
    let start_time = from.unwrap_or(default_time);
    user_event::table
        .filter(
            user_event::time
                .gt(start_time)
                .or(user_event::time.eq(start_time)),
        )
        .select(UserEvent::as_select())
        .load(conn)
}

pub fn get_hotel_events(
    conn: &mut PgConnection,
    from: &Option<DateTime<Utc>>,
) -> Result<Vec<HotelEvent>, diesel::result::Error> {
    let default = "2000-01-01T00:00:00+00:00";
    let default_time = DateTime::parse_from_rfc3339(default).unwrap().into();
    let start_time = from.unwrap_or(default_time);
    hotel_event::table
        .filter(
            hotel_event::time
                .gt(start_time)
                .or(hotel_event::time.eq(start_time)),
        )
        .select(HotelEvent::as_select())
        .load(conn)
}

pub fn get_checkout_events(
    conn: &mut PgConnection,
    from: &Option<DateTime<Utc>>,
) -> Result<Vec<CheckoutEvent>, diesel::result::Error> {
    let default = "2000-01-01T00:00:00+00:00";
    let default_time = DateTime::parse_from_rfc3339(default).unwrap().into();
    let start_time = from.unwrap_or(default_time);
    checkout_event::table
        .filter(
            checkout_event::time
                .gt(start_time)
                .or(checkout_event::time.eq(start_time)),
        )
        .select(CheckoutEvent::as_select())
        .load(conn)
}

#[cfg(test)]
mod tests {
    use crate::{model::UserEvent, schema::user_event};
    use chrono::Utc;

    use super::*;

    #[test]
    fn it_works() {
        let pool = &mut get_connection_pool().get().unwrap();

        let binding = Utc::now();
        let new_event = NewUserEvent::new("login", Some("bad message"), &binding);
        let res: Result<usize, diesel::result::Error> = add_user_event(pool, new_event);

        let date_str = "2023-12-30T12:53:29.260266Z";
        let datetime: DateTime<Utc> = DateTime::parse_from_rfc3339(date_str).unwrap().into();

        let result = get_user_events(pool, &Some(datetime));

        println!("{result:?}");
        println!("{res:?}");

        assert_eq!(result.unwrap()[0].type_, "logind");
    }
}
