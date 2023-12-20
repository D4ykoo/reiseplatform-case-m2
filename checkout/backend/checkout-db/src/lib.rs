use std::env;
use diesel::{prelude::*, update};
use dotenvy::dotenv;
use diesel::r2d2::ConnectionManager;
use r2d2::Pool;

use crate::models::{NewCart, Cart};

mod models;
mod schema;

pub type PostgresPool = Pool<ConnectionManager<PgConnection>>;

pub fn get_pool() -> PostgresPool {
    dotenv().ok();

    let url = env::var("DATABASE_URL").expect("DATABASE_URL must be set!");
    let mgr = ConnectionManager::<PgConnection>::new(url);

    r2d2::Pool::builder().build(mgr).expect("Could not build connection pool")
}

pub fn connect_db() -> PgConnection{
    dotenv().ok();

    let url = env::var("DATABASE_URL").expect("DATABASE_URL must be set!");

    PgConnection::establish(&url)
        .unwrap_or_else(|e| panic!("Error connecting to {} \n{e}", url))
}

pub fn create_cart(conn: &mut PgConnection, new_cart: NewCart) -> Result<usize, diesel::result::Error> {
    use self::schema::cart;

    let inserted_rows = diesel::insert_into(cart::table)
        .values(&new_cart)
        .execute(conn)
        .expect("Error can not create a new cart");

    Ok(inserted_rows)
}

pub fn remove_cart(conn: &mut PgConnection, cart_id: &i32) -> Result<usize, diesel::result::Error>{
    use self::schema::cart::dsl::*;

    let res = diesel::delete(cart.filter(id.eq(cart_id))).execute(conn);

    res
}

pub fn get_cart(conn: &mut PgConnection, cart_id: &i32) -> Option<Cart> {
    use self::schema::cart::dsl::*;

    let res = cart
        .filter(id.eq(cart_id))
        .select(Cart::as_select())
        .first(conn)
        .expect(&format!("Error receiving cart with id {}", cart_id));

    return Some(res)
}

pub fn update_card(conn: &mut PgConnection, cart_id: &i32, new_cart: NewCart) -> Result<bool, &'static str>{
    use self::schema::cart::dsl::*;
    let cart;
    
    if new_cart.offers != None {
        let _ = update(cart.filter(id.eq(cart_id)))
            .set(offers.eq(new_cart.offers)).execute(conn);
    }
    if new_cart.paid != None {
        let _ = update(cart.filter(id.eq(cart_id)))
        .set(paid.eq(new_cart.paid)).execute(conn);
    }
    if new_cart.payment_method != None {
        let _ = update(cart.filter(id.eq(cart_id)))
        .set(payment_method.eq(new_cart.payment_method)).execute(conn);
    }

    let single_cart = get_cart(conn, cart_id).unwrap();
    let updated_cart = NewCart::create(single_cart.paid, single_cart.payment_method.as_deref(), single_cart.offers.as_ref());
    
    if updated_cart.eq(&new_cart) {
        return Ok(true)
    } else {
        return Err("Error updating cart");
    }
}