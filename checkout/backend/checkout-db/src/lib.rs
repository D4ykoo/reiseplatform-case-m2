use diesel::r2d2::ConnectionManager;
use diesel::{prelude::*, PgConnection};
use dotenvy::dotenv;
use models::{Hotel, NewHotel, NewTravelSlice, TravelSlice, CombinedCart};
use r2d2::Pool;
use std::env;

use crate::models::{Cart, NewCart};

pub mod models;
mod schema;

pub type PostgresPool = Pool<ConnectionManager<PgConnection>>;

pub fn get_pool() -> PostgresPool {
    dotenv().ok();

    let url = env::var("DATABASE_URL").expect("DATABASE_URL must be set!");
    let mgr = ConnectionManager::<PgConnection>::new(url);

    r2d2::Pool::builder()
        .build(mgr)
        .expect("Could not build connection pool")
}

pub fn connect_db() -> PgConnection {
    dotenv().ok();

    let url = env::var("DATABASE_URL").expect("DATABASE_URL must be set!");

    PgConnection::establish(&url).unwrap_or_else(|e| panic!("Error connecting to {} \n{e}", url))
}

fn create_hotel(conn: &mut PgConnection, hotel: &Hotel) -> Result<usize, diesel::result::Error> {
    use self::schema::checkouthotel;

    // check if cart is in db
    let cart = get_cart(conn, &hotel.fk_checkout_cart.unwrap());

    if cart.is_none() {
        return Err(diesel::result::Error::NotFound);
    }

    let new_hotel = NewHotel::create(
        hotel.hotelname.as_deref(),
        hotel.land.as_deref(),
        hotel.vendor_name.as_deref(),
        hotel.hotel_description.as_deref(),
        hotel.hotel_image.as_deref(),
        hotel.fk_checkout_cart,
    );

    let inserted_rows = diesel::insert_into(checkouthotel::table)
        .values(&new_hotel)
        .execute(conn)
        .expect("Error can not create a new hotel");

    Ok(inserted_rows)
}

fn get_all_hotels(conn: &mut PgConnection, cart_id: &i32) -> Option<Vec<Hotel>> {
    use self::schema::checkouthotel::dsl::*;

    let res = checkouthotel
        .filter(fk_checkout_cart.eq(cart_id))
        .select(Hotel::as_select())
        .load::<Hotel>(conn)
        .expect(&format!("Error receiving hotels with cart id {}", cart_id));

    return Some(res);
}

fn get_single_hotel(conn: &mut PgConnection, hote_id: &i32) -> Option<Hotel> {
    use self::schema::checkouthotel::dsl::*;

    let res = checkouthotel
        .filter(id.eq(hote_id))
        .select(Hotel::as_select())
        .first(conn)
        .expect(&format!("Error receiving hotel with id {}", hote_id));

    return Some(res);
}

fn delete_hotel(conn: &mut PgConnection, hotel_id: &i32) -> Result<usize, diesel::result::Error> {
    use self::schema::checkouthotel::dsl::*;

    let deleted_rows = diesel::delete(checkouthotel.filter(id.eq(hotel_id)))
        .execute(conn)
        .expect(&format!("Error deleting hotel with id {}", hotel_id));

    Ok(deleted_rows)
}

fn create_travel_slice(
    conn: &mut PgConnection,
    travel_slice: &TravelSlice,
) -> Result<usize, diesel::result::Error> {
    use self::schema::checkouttravelslice;

    // check if hotel is in db
    let hotel = get_single_hotel(conn, &travel_slice.fk_checkout_hotel.unwrap());

    if hotel.is_none() {
        return Err(diesel::result::Error::NotFound);
    }

    // if present, then create a new travel slice
    let new_travel_slice = NewTravelSlice::create(
        travel_slice.vendor_name.as_deref(),
        travel_slice.price,
        travel_slice.from_date.as_deref(),
        travel_slice.to_date.as_deref(),
        travel_slice.fk_checkout_hotel,
    );

    let inserted_rows = diesel::insert_into(checkouttravelslice::table)
        .values(&new_travel_slice)
        .execute(conn)
        .expect("Error can not create a new travel slice");

    Ok(inserted_rows)
}

fn get_all_travel_slices(
    conn: &mut PgConnection,
    hotel_id: &i32,
) -> Option<Vec<TravelSlice>> {
    use self::schema::checkouttravelslice::dsl::*;

    let res = checkouttravelslice
        .filter(fk_checkout_hotel.eq(hotel_id))
        .select(TravelSlice::as_select())
        .load::<TravelSlice>(conn)
        .expect(&format!(
            "Error receiving travel slices with hotel id {}",
            hotel_id
        ));

    return Some(res);
}


#[allow(dead_code)]
fn get_single_travel_slice(
    conn: &mut PgConnection,
    travel_slice_id: &i32,
) -> Option<TravelSlice> {
    use self::schema::checkouttravelslice::dsl::*;

    let res = checkouttravelslice
        .filter(id.eq(travel_slice_id))
        .select(TravelSlice::as_select())
        .first(conn)
        .expect(&format!(
            "Error receiving travel slice with id {}",
            travel_slice_id
        ));

    return Some(res);
}

fn delete_travel_slice(
    conn: &mut PgConnection,
    travel_slice_id: &i32,
) -> Result<usize, diesel::result::Error> {
    use self::schema::checkouttravelslice::dsl::*;

    let deleted_rows = diesel::delete(checkouttravelslice.filter(id.eq(travel_slice_id)))
        .execute(conn)
        .expect(&format!(
            "Error deleting travel slice with id {}",
            travel_slice_id
        ));

    Ok(deleted_rows)
}

pub fn add_to_cart(
    conn: &mut PgConnection,
    cart_id: &i32,
    hotel: &Hotel,
    travel_sclice: &TravelSlice,
) -> Result<(), diesel::result::Error>{
    let cart = get_cart(conn, cart_id);

    if cart.is_none() {
        let new_cart = NewCart::create(None, None, None);
        let _ = create_cart(conn, new_cart)?;
    }

    let _ = create_hotel(conn, hotel)?;
    
    let hotel = get_single_hotel(conn, &hotel.id);
    let hotel_id = hotel.unwrap().id;

    let mut travel_sclice = travel_sclice.clone();
    travel_sclice.fk_checkout_hotel = Some(hotel_id);

    let _ = create_travel_slice(conn, &travel_sclice)?;

    // now everything is addeded since the foreign keys are set
    // and contain a reference to the cart as well as the hotel
    Ok(())
}

pub fn get_cart_content(conn: &mut PgConnection, cart_id: &i32) -> Option<CombinedCart> {
    let hotels = get_all_hotels(conn, cart_id);

    if hotels.is_none() {
        return None;
    }

    let mut hotels = hotels.unwrap();

    let mut travels: Vec<TravelSlice> = Vec::new();

    for hotel in hotels.iter_mut() {
        let travel_slices = get_all_travel_slices(conn, &hotel.id);

        if travel_slices.is_none() {
            return None;
        }

        travels.append(&mut travel_slices.unwrap());
    }

    // create and return combined cart
    let combined_cart = CombinedCart {
        cart: get_cart(conn, cart_id).unwrap(),
        hotel: Some(hotels),
        travel_slice: Some(travels),
    };

    return Some(combined_cart);

}

pub fn create_cart(
    conn: &mut PgConnection,
    new_cart: NewCart,
) -> Result<usize, diesel::result::Error> {
    use self::schema::checkoutcart;

    let inserted_rows = diesel::insert_into(checkoutcart::table)
        .values(&new_cart)
        .execute(conn)
        .expect("Error can not create a new cart");

    Ok(inserted_rows)
}

pub fn remove_cart(conn: &mut PgConnection, cart_id: &i32) -> Result<usize, diesel::result::Error> {
    use self::schema::checkoutcart::dsl::*;

    let res = diesel::delete(checkoutcart.filter(id.eq(cart_id))).execute(conn);

    res
}

pub fn get_cart(conn: &mut PgConnection, cart_id: &i32) -> Option<Cart> {
    use self::schema::checkoutcart::dsl::*;

    let res = checkoutcart
        .filter(id.eq(cart_id))
        .select(Cart::as_select())
        .first(conn)
        .expect(&format!("Error receiving cart with id {}", cart_id));

    return Some(res);
}

pub fn remove_hotel_and_travel_slice(
    conn: &mut PgConnection,
    hotel_id: &i32,
    travel_slice_id: &i32,
) -> Result<(), diesel::result::Error> {
    let _ = delete_travel_slice(conn, travel_slice_id)?;
    let _ = delete_hotel(conn, hotel_id)?;

    Ok(())
}