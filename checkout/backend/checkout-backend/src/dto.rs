//! Data transfer objects for responses and requests
//!
//! Even though there is not a lot of difference to the entity model
//! for a better extensablity this dtos will be used.

use checkout_db::models::NewCart;
use serde::{Deserialize, Serialize};

#[derive(Serialize)]
pub struct CartResponse {
    pub paid: Option<bool>,
    pub payment_method: Option<String>,
    pub offers: Option<Vec<Option<i32>>>,
}

impl CartResponse {
    pub fn from_db_cart(cart: checkout_db::models::Cart) -> Self {
        CartResponse {
            paid: cart.paid,
            payment_method: cart.payment_method,
            offers: cart.offers,
        }
    }
}

#[derive(Deserialize, Debug)]
pub struct CartRequest {
    pub paid: Option<bool>,
    pub payment_method: Option<String>,
    pub offers: Option<Vec<Option<i32>>>,
}

impl CartRequest {
    pub fn into_new_cart(&self) -> checkout_db::models::NewCart {
        NewCart {
            paid: self.paid,
            payment_method: self.payment_method.as_deref(),
            offers: self.offers.as_ref(),
        }
    }
}


#[derive(Serialize)]
pub struct HotelTravelResponse {
    pub id: i64,
    pub hotelname: String,
    pub land: String,
    pub vendorname: String,
    pub description: String,
    pub pictures: String,
    pub travels: Vec<TravelSlice>,
}

#[derive(Serialize)]
pub struct TravelSlice {
    pub vendorname: String,
    pub from: String,
    pub to: String,
    pub price: i64,
}
