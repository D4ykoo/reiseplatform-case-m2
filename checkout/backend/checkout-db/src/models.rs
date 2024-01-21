use diesel::prelude::*;
use serde::Deserialize;
use serde::Serialize;

use crate::schema::checkoutcart;
use crate::schema::checkouthotel;
use crate::schema::checkouttravelslice;

#[derive(Debug, Deserialize, Serialize, Clone)]
pub struct CombinedCart {
    pub cart: Cart,
    pub hotel: Option<Vec<Hotel>>,
    pub travel_slice: Option<Vec<TravelSlice>>,

}

#[derive(Queryable, Selectable, Debug, Deserialize, Serialize, Clone, PartialEq, Eq)]
#[diesel(table_name = checkoutcart)]
#[diesel(check_for_backend(diesel::pg::Pg))]
pub struct Cart {
    pub id: i32,
    pub user_id: Option<i32>,
    pub paid: Option<bool>,
    pub payment_method: Option<String>,
}

#[derive(Queryable, Selectable, Debug, Deserialize, Serialize, Clone)]
#[diesel(table_name = checkouthotel)]
#[diesel(check_for_backend(diesel::pg::Pg))]
pub struct Hotel {
    pub id: i32,
    pub hotelname: Option<String>,
    pub land: Option<String>,
    pub vendor_name: Option<String>,
    pub hotel_description: Option<String>,
    pub hotel_image: Option<String>,
    pub fk_checkout_cart: Option<i32>,
}

#[derive(Queryable, Selectable, Debug, Deserialize, Serialize, Clone)]
#[diesel(table_name = checkouttravelslice)]
#[diesel(check_for_backend(diesel::pg::Pg))]
pub struct TravelSlice {
    pub id: i32,
    pub vendor_name: Option<String>,
    pub price: Option<i32>,
    pub from_date: Option<String>,
    pub to_date: Option<String>,
    pub fk_checkout_hotel: Option<i32>,
}


// Now all all insertable structs are defined
// this is needed due to lifetime stuff in rust

#[derive(Insertable, Clone, PartialEq, Eq)]
#[diesel(table_name = checkoutcart)]
pub struct NewCart<'a> {
    pub user_id: Option<i32>,
    pub paid: Option<bool>,
    pub payment_method: Option<&'a str>,
}

impl<'a> NewCart<'a> {
    pub fn create(
        user_id: Option<i32>,
        paid: Option<bool>,
        payment_method: Option<&'a str>,
    ) -> Self {
        NewCart {
            user_id,
            paid,
            payment_method,
        }
    }
}

#[derive(Insertable, Clone, PartialEq, Eq)]
#[diesel(table_name = checkouthotel)]
pub struct NewHotel<'a> {
    pub hotelname: Option<&'a str>,
    pub land: Option<&'a str>,
    pub vendor_name: Option<&'a str>,
    pub hotel_description: Option<&'a str>,
    pub hotel_image: Option<&'a str>,
    pub fk_checkout_cart: Option<i32>,
}

impl<'a> NewHotel<'a> {
    pub fn create(
        hotelname: Option<&'a str>,
        land: Option<&'a str>,
        vendor_name: Option<&'a str>,
        hotel_description: Option<&'a str>,
        hotel_image: Option<&'a str>,
        fk_checkout_cart: Option<i32>,
    ) -> Self {
        NewHotel {
            hotelname,
            land,
            vendor_name,
            hotel_description,
            hotel_image,
            fk_checkout_cart,
        }
    }
}

#[derive(Insertable, Clone, PartialEq, Eq)]
#[diesel(table_name = checkouttravelslice)]
pub struct NewTravelSlice<'a> {
    pub vendor_name: Option<&'a str>,
    pub price: Option<i32>,
    pub from_date: Option<&'a str>,
    pub to_date: Option<&'a str>,
    pub fk_checkout_hotel: Option<i32>,
}

impl<'a> NewTravelSlice<'a> {
    pub fn create(
        vendor_name: Option<&'a str>,
        price: Option<i32>,
        from_date: Option<&'a str>,
        to_date: Option<&'a str>,
        fk_checkout_hotel: Option<i32>,
    ) -> Self {
        NewTravelSlice {
            vendor_name,
            price,
            from_date,
            to_date,
            fk_checkout_hotel,
        }
    }
}