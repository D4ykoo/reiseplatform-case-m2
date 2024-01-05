
use std::env;

use dotenvy::dotenv;

use crate::models::HotelTravels;

pub async fn get_hotel_travel_request_with_cookie(hotel_id: i32, travel_id: i32, cookie: &str) -> Result<HotelTravels, Box<dyn std::error::Error>> {
    dotenv().ok();
    let travel_api_url = env::var("TRAVEL_API_URL").expect("TRAVEL_API_URL must be set");
    let composed_api_url = format!("{}/hotels/{}/travels/{}", travel_api_url, hotel_id, travel_id);
    let res = reqwest::Client::new()
        .get(composed_api_url.as_str())
        .header("Cookie", cookie)
        .send()
        .await?
        .json::<HotelTravels>()
        .await?;

    Ok(res)
} 