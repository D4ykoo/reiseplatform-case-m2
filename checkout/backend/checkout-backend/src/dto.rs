use checkout_db::models::Cart;
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct CombinedCartResponse {
    pub cart: Cart,
    pub hotel: Option<Vec<HotelResponse>>,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct HotelResponse {
    pub hotelname: String,
    pub land: String,
    pub vendorname: String,
    pub description: String,
    pub pictures: String,
    pub travels: Vec<TravelSliceResponse>,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct TravelSliceResponse {
    pub vendorname: String,
    pub from: String,
    pub to: String,
    pub price: i64,
}

impl CombinedCartResponse {
    pub fn from_db_combined_cart(cart: checkout_db::models::CombinedCart) -> Self {
        CombinedCartResponse {
            cart: cart.cart,
            hotel: cart.hotel.map(|hotels| {
                hotels.into_iter().map(|hotel| {
                    HotelResponse::from_db_hotel(hotel)
                }).collect()
            }),
        }
    }
}

impl TravelSliceResponse {
    pub fn from_db_travel_slice(travel_slice: checkout_db::models::TravelSlice) -> Self {
        TravelSliceResponse {
            vendorname: travel_slice.vendor_name.unwrap(),
            from: travel_slice.from_date.unwrap(),
            to: travel_slice.to_date.unwrap(),
            price: travel_slice.price.unwrap() as i64,
        }
    }
}

impl HotelResponse {
    pub fn from_db_hotel(hotel: checkout_db::models::Hotel) -> Self {
        HotelResponse {
            hotelname: hotel.hotelname.unwrap(),
            land: hotel.land.unwrap(),
            vendorname: hotel.vendor_name.unwrap(),
            description: hotel.hotel_description.unwrap(),
            pictures: hotel.hotel_image.unwrap(),
            travels: Vec::new(),
        }
    }

    pub fn add_travel_slice(&mut self, travel_slices: Vec<checkout_db::models::TravelSlice>) {
        for travel_slice in travel_slices {
            self.travels.push(TravelSliceResponse::from_db_travel_slice(travel_slice));
        }
    }
}