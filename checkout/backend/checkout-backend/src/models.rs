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
