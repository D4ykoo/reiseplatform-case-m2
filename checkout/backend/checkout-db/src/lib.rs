use diesel::r2d2::ConnectionManager;
use diesel::{prelude::*, PgConnection};
use diesel_migrations::{embed_migrations, EmbeddedMigrations, MigrationHarness};

use dotenvy::dotenv;
use models::{Hotel, NewHotel, NewTravelSlice, TravelSlice, CombinedCart};
use r2d2::Pool;
use std::env;
use std::error::Error;
use crate::models::{Cart, NewCart};

pub mod models;
mod schema;

pub type PostgresPool = Pool<ConnectionManager<PgConnection>>;
const MIGRATIONS: EmbeddedMigrations = embed_migrations!("migrations/");

pub fn init_migrations(conn: &mut PgConnection) ->  Result<(), Box<dyn Error + Send + Sync + 'static>> {
    conn.run_pending_migrations(MIGRATIONS)?;
    Ok(())
}

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

fn create_hotel(conn: &mut PgConnection, new_hotel: &NewHotel) -> Result<i32, diesel::result::Error> {
    use self::schema::checkouthotel;

    // check if cart is in db
    let cart = get_cart(conn, &new_hotel.fk_checkout_cart.unwrap());

    if cart.is_none() {
        return Err(diesel::result::Error::NotFound);
    }

    // let new_hotel = NewHotel::create(
    //     hotel.hotelname.as_deref(),
    //     hotel.land.as_deref(),
    //     hotel.vendor_name.as_deref(),
    //     hotel.hotel_description.as_deref(),
    //     hotel.hotel_image.as_deref(),
    //     hotel.fk_checkout_cart,
    // );

    let inserted_rows = diesel::insert_into(checkouthotel::table)
        .values(new_hotel)
        .returning(checkouthotel::id)
        // .execute(conn)
        // .expect("Error can not create a new hotel");
        .get_result::<i32>(conn);

    Ok(inserted_rows.unwrap())
}

fn get_all_hotels(conn: &mut PgConnection, cart_id: &i32) -> Option<Vec<Hotel>> {
    use self::schema::checkouthotel::dsl::*;

    let res = checkouthotel
        .filter(fk_checkout_cart.eq(cart_id))
        .select(Hotel::as_select())
        .load::<Hotel>(conn)
        .unwrap_or_else(|_| panic!("Error receiving hotels with cart id {}", cart_id));

    Some(res)
}

pub fn get_single_hotel(conn: &mut PgConnection, hote_id: &i32) -> Option<Hotel> {
    use self::schema::checkouthotel::dsl::*;

    let res = checkouthotel
        .filter(id.eq(hote_id))
        .select(Hotel::as_select())
        .first(conn)
        .unwrap_or_else(|_| panic!("Error receiving hotel with id {}", hote_id));

    Some(res)
}

fn delete_hotel(conn: &mut PgConnection, hotel_id: &i32) -> Result<usize, diesel::result::Error> {
    use self::schema::checkouthotel::dsl::*;

    let deleted_rows = diesel::delete(checkouthotel.filter(id.eq(hotel_id)))
        .execute(conn)
        .unwrap_or_else(|_| panic!("Error deleting hotel with id {}", hotel_id));

    Ok(deleted_rows)
}

fn create_travel_slice(
    conn: &mut PgConnection,
    travel_slice: &NewTravelSlice,
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
        .unwrap_or_else(|_| panic!("Error receiving travel slices with hotel id {}",
            hotel_id));

    Some(res)
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
        .unwrap_or_else(|_| panic!("Error receiving travel slice with id {}",
            travel_slice_id));

    Some(res)
}

fn delete_travel_slice(
    conn: &mut PgConnection,
    travel_slice_id: &i32,
) -> Result<usize, diesel::result::Error> {
    use self::schema::checkouttravelslice::dsl::*;

    let deleted_rows = diesel::delete(checkouttravelslice.filter(id.eq(travel_slice_id)))
        .execute(conn)
        .unwrap_or_else(|_| panic!("Error deleting travel slice with id {}",
            travel_slice_id));

    Ok(deleted_rows)
}

pub fn get_cart_id(conn: &mut PgConnection, userid: &i32) -> Result<i32, diesel::result::Error>{
    use self::schema::checkoutcart::dsl::*;
    let res = checkoutcart
        .filter(user_id.eq(userid))
        .select(id)
        .first(conn);

    res
}

// find hotel by cart id and hotel name
pub fn find_hotel_by_cart_id_and_hotel_name(
    conn: &mut PgConnection,
    cart_id: &i32,
    hotel_name: &str,
) -> Option<Hotel> {
    use self::schema::checkouthotel::dsl::*;

    let res = checkouthotel
        .filter(fk_checkout_cart.eq(cart_id))
        .filter(hotelname.eq(hotel_name))
        .select(Hotel::as_select())
        .first(conn)
        .unwrap_or_else(|_| panic!("Error receiving hotel with cart id {} and hotel name {}",
            cart_id, hotel_name));

    Some(res)
}

pub fn add_to_cart(
    conn: &mut PgConnection,
    cart_id: &i32,
    new_hotel: &NewHotel,
    travel_sclice: &NewTravelSlice,
) -> Result<(), diesel::result::Error>{
    let cart = get_cart(conn, cart_id);

    if cart.is_none() {
        let new_cart = NewCart::create(None, None, None);
        let _ = create_cart(conn, new_cart)?;
    }

    let hotel_id: i32;

    // check if hotel is in db otherwise create it
    let hotel = find_hotel_by_cart_id_and_hotel_name(conn, cart_id, new_hotel.hotelname.unwrap());

    if hotel.is_none() {
        hotel_id = create_hotel(conn, new_hotel)?;
    } else {
        hotel_id = hotel.unwrap().id;
    }

    let mut travel_sclice = travel_sclice.clone();
    travel_sclice.fk_checkout_hotel = Some(hotel_id);

    let _ = create_travel_slice(conn, &travel_sclice)?;

    // now everything is addeded since the foreign keys are set
    // and contain a reference to the cart as well as the hotel
    Ok(())
}

pub fn get_cart_content(conn: &mut PgConnection, cart_id: &i32) -> Option<CombinedCart> {
    let hotels = get_all_hotels(conn, cart_id);

    hotels.as_ref()?;

    let mut hotels = hotels.unwrap();

    let mut travels: Vec<TravelSlice> = Vec::new();

    for hotel in hotels.iter_mut() {
        let travel_slices = get_all_travel_slices(conn, &hotel.id);

        travel_slices.as_ref()?;

        travels.append(&mut travel_slices.unwrap());
    }

    // create and return combined cart
    let combined_cart = CombinedCart {
        cart: get_cart(conn, cart_id).unwrap(),
        hotel: Some(hotels),
        travel_slice: Some(travels),
    };

    Some(combined_cart)

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

    diesel::delete(checkoutcart.filter(id.eq(cart_id))).execute(conn)
}

pub fn get_cart(conn: &mut PgConnection, cart_id: &i32) -> Option<Cart> {
    use self::schema::checkoutcart::dsl::*;

    let res = checkoutcart
        .filter(id.eq(cart_id))
        .select(Cart::as_select())
        .first(conn)
        .unwrap_or_else(|_| panic!("Error receiving cart with id {}", cart_id));

    Some(res)
}


// since the ids of the hotels and travel slices are not constant and increment 
// automatically get the ids from the db and then delete them,
// since the array like structure can be assumed when extracting the data from the db
pub fn remove_hotel_and_travel_slice(
    conn: &mut PgConnection,
    cart_id: &i32,
    hotel_id: &i32,
    travel_slice_id: &i32,
) -> Result<(), diesel::result::Error> {
    let hotels = get_all_hotels(conn, cart_id);
     

    let hotel = &hotels
        .as_ref()
        .unwrap()[*hotel_id as usize];
        

    let travel_slices = get_all_travel_slices(conn, &hotel.id);

    let travel_slice = &travel_slices
        .as_ref()
        .unwrap()[*travel_slice_id as usize];


    let _ = delete_travel_slice(conn, &travel_slice.id)?;
    let _ = delete_hotel(conn, &hotel.id);

    Ok(())
}