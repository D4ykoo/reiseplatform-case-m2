pub mod model;
pub mod schema;

use chrono::{DateTime, Utc};
use deadpool_diesel::postgres::Manager;
use deadpool_diesel::Pool;
use diesel::prelude::*;
use diesel::PgConnection;
use diesel_migrations::{embed_migrations, EmbeddedMigrations, MigrationHarness};
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
use std::error::Error;

const MIGRATIONS: EmbeddedMigrations = embed_migrations!("migrations/");

pub fn get_connection_pool() -> Pool<Manager> {
    dotenv().ok();
    let host = env::var("DB_HOST").expect("DATABASE_URL must be set");
    let user = env::var("DB_USER").expect("DATABASE_URL must be set");
    let passwd = env::var("DB_PASSWORD").expect("DATABASE_URL must be set");
    let name = env::var("DB_NAME").expect("DATABASE_URL must be set");

    let db_url = format!("postgres://{}:{}@{}/{}", user, passwd, host, name);
    // set up connection pool
    let manager =
        deadpool_diesel::postgres::Manager::new(&db_url, deadpool_diesel::Runtime::Tokio1);

    deadpool_diesel::postgres::Pool::builder(manager)
        .build()
        .unwrap_or_else(|_| panic!("Error connecting to {db_url}"))
}

pub fn establish_connection() -> PgConnection {
    dotenv().ok();
    let host = env::var("DB_HOST").expect("DATABASE_URL must be set");
    let user = env::var("DB_USER").expect("DATABASE_URL must be set");
    let passwd = env::var("DB_PASSWORD").expect("DATABASE_URL must be set");
    let name = env::var("DB_NAME").expect("DATABASE_URL must be set");

    let db_url = format!("postgres://{}:{}@{}/{}", user, passwd, host, name);

    PgConnection::establish(&db_url).unwrap_or_else(|_| panic!("Error connecting to {db_url}"))
}

pub fn run_migrations(
    conn: &mut PgConnection,
) -> Result<(), Box<dyn Error + Send + Sync + 'static>> {
    // This will run the necessary migrations.
    conn.run_pending_migrations(MIGRATIONS)?;

    Ok(())
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
    from: &DateTime<Utc>,
) -> Result<Vec<UserEvent>, diesel::result::Error> {
    user_event::table
        .filter(user_event::time.gt(from).or(user_event::time.eq(from)))
        .select(UserEvent::as_select())
        .get_results(conn)
}

pub fn get_hotel_events(
    conn: &mut PgConnection,
    from: &DateTime<Utc>,
) -> Result<Vec<HotelEvent>, diesel::result::Error> {
    hotel_event::table
        .filter(hotel_event::time.gt(from).or(hotel_event::time.eq(from)))
        .select(HotelEvent::as_select())
        .load(conn)
}

pub fn get_checkout_events(
    conn: &mut PgConnection,
    from: &DateTime<Utc>,
) -> Result<Vec<CheckoutEvent>, diesel::result::Error> {
    checkout_event::table
        .filter(
            checkout_event::time
                .gt(from)
                .or(checkout_event::time.eq(from)),
        )
        .select(CheckoutEvent::as_select())
        .load(conn)
}

#[cfg(test)]
mod tests {
    use chrono::Utc;

    use super::*;

    #[test]
    fn it_works() {
        let pool = &mut establish_connection();

        let binding = Utc::now();
        let new_event = NewUserEvent::new("login".into(), Some("bad message".into()), binding);
        let res: Result<usize, diesel::result::Error> = add_user_event(pool, new_event);

        let date_str = "2023-12-30T12:53:29.260266Z";
        let datetime: DateTime<Utc> = DateTime::parse_from_rfc3339(date_str).unwrap().into();

        let result = get_user_events(pool, &datetime);

        println!("{result:?}");
        println!("{res:?}");

        assert_eq!(result.unwrap()[0].type_, "logind");
    }
}
