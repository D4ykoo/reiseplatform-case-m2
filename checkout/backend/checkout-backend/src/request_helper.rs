
use std::env;

use dotenvy::dotenv;

use crate::models::HotelTravels;

async fn get_hotel_travel_request(hotel_id: i32, travel_id: i32) -> Result<HotelTravels, Box<dyn std::error::Error>> {
    dotenv().ok();
    let travel_api_url = env::var("TRAVEL_API_URL").expect("TRAVEL_API_URL must be set");

    let res = reqwest::get(
        &format!("{}/hotels/{}/travels/{}", travel_api_url, hotel_id, travel_id)
        )
        .await?
        .json::<HotelTravels>()
        .await?;

    Ok(res)
}
