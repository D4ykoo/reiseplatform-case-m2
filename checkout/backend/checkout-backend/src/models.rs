use serde::{Serialize, Deserialize};

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct HotelTravels {
    pub id: i64,
    pub hotelname: String,
    pub street: String,
    pub state: String,
    pub land: String,
    pub vendorid: String,
    pub vendorname: String,
    pub description: String,
    pub pictures: Vec<Picture>,
    pub tags: Vec<String>,
    pub travels: Vec<Travel>,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Picture {
    pub id: i64,
    pub description: String,
    pub payload: String,
}

#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Travel {
    pub id: i64,
    pub vendorid: i64,
    pub vendorname: String,
    pub from: String,
    pub to: String,
    pub price: i64,
    pub description: String,
    pub createdat: String,
    pub updatedat: String,
}

// impl to convert to db model
impl HotelTravels {
    pub fn to_db_hotel(&self, cart_id: i32) -> checkout_db::models::NewHotel {
        checkout_db::models::NewHotel {
            hotelname: Some(&self.hotelname),
            land: Some(&self.land),
            vendor_name: Some(&self.vendorname),
            hotel_description: Some(&self.description),
            hotel_image: Some(&self.pictures[0].payload),
            fk_checkout_cart: Some(cart_id),
        }
    }

    pub fn to_db_travel_slice(&self, hotel_id: i32) -> checkout_db::models::NewTravelSlice {
        checkout_db::models::NewTravelSlice {
            vendor_name: Some(&self.vendorname),
            price: Some(self.travels[0].price as i32),
            from_date: Some(&self.travels[0].from),
            to_date: Some(&self.travels[0].to),
            fk_checkout_hotel: Some(hotel_id),
        }
    }

}