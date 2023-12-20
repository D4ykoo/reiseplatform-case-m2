use std::env;
use diesel::pg::PgConnection;
use dotenvy::dotenv;
use diesel::r2d2::ConnectionManager;
use r2d2::Pool;

mod models;
mod schema;

pub type PostgresPool = Pool<ConnectionManager<PgConnection>>;

pub fn get_pool() -> PostgresPool {
    dotenv().ok();

    let url = env::var("DATABASE_URL").expect("DATABASE_URL must be set!");
    let mgr = ConnectionManager::<PgConnection>::new(url);

    r2d2::Pool::builder().build(mgr).expect("Could not build connection pool")
}