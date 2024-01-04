
use std::env;

use dotenvy::dotenv;

use crate::models::HotelTravels;

pub async fn get_hotel_travel_request(hotel_id: i32, travel_id: i32) -> Result<HotelTravels, Box<dyn std::error::Error>> {
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

pub async fn get_hotel_travel_request_with_cookie(hotel_id: i32, travel_id: i32, cookie: &str) -> Result<HotelTravels, Box<dyn std::error::Error>> {
    dotenv().ok();
    // TODO: MAKE ENV URL WORK
    let travel_api_url = env::var("TRAVEL_API_URL").expect("TRAVEL_API_URL must be set");
    let composed_api_url = format!("{}/hotels/{}/travels/{}", travel_api_url, hotel_id, travel_id);
    let res = reqwest::Client::new()
        // .get("http://localhost:8086/api/v1/hotels/1/travels/1")
        .get(composed_api_url.as_str())
        // .get(
        //     &format!("/hotels/{}/travels/{}", travel_api_url, hotel_id, travel_id)
        // )
        .header("Cookie", cookie)
        .send()
        .await?
        .json::<HotelTravels>()
        .await?;

    Ok(res)
} 