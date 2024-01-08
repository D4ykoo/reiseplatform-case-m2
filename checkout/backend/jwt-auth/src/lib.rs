//! Simple JWT validator that uses HS512 signing algorithm.
//!
//! Claims: username, iat
//! Examples not due to simplicity not provided.

use dotenvy::dotenv;
use jsonwebtoken::{decode, Algorithm, DecodingKey, Validation};
use models::*;
use std::env;

pub mod models;

pub fn validate_jwt(token: &str) -> Result<bool, jsonwebtoken::errors::Error> {
    dotenv().ok();
    let jwt_secret = env::var("JWT_SECRET").expect("Secret must be set");

    let mut val = Validation::new(Algorithm::HS512);
    val.set_required_spec_claims(&["username", "user_id", "iat"]);

    let decoded = decode::<Claims>(
        &token,
        &DecodingKey::from_secret(jwt_secret.as_bytes()),
        &val,
    );

    match decoded {
        Ok(_) => return Ok(true),
        Err(e) => return Err(e),
    };
}

pub fn decode_jwt(token: &str) -> Result<Claims, jsonwebtoken::errors::Error> {
    dotenv().ok();

    let mut val = Validation::new(Algorithm::HS512);
    val.set_required_spec_claims(&["username", "user_id", "iat"]);

    let jwt_secret = env::var("JWT_SECRET").expect("Secret must be set");
    let decoded = decode::<Claims>(
        &token,
        &DecodingKey::from_secret(jwt_secret.as_bytes()),
        &val,
    )?;

    return Ok(decoded.claims);
}

#[cfg(test)]
mod tests {
    use crate::validate_jwt;

    #[test]
    fn token_valid() {
        let token = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjY4NjYyMDgwMCwidXNlcm5hbWUiOiJ0ZXN0In0.CbmybuROnf_3ClsxXiYiTbK26Dc0e2zSwMeCZZz4guszI-q8LL6HO42HJTeAjQ0gDFRmL4PikQoP8QzdPC03yw";
        let result = validate_jwt(token);
        assert_eq!(result.is_ok(), true);
    }

    #[test]
    fn token_invalid() {
        let token = "pyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjY4NjYyMDgwMCwidXNlcm5hbWUiOiJ0ZXN0In0.CbmybuROnf_3ClsxXiYiTbK26Dc0e2zSwMeCZZz4guszI-q8LL6HO42HJTeAjQ0gDFRmL4PikQoP8QzdPC03yw";
        let result = validate_jwt(token);
        assert_eq!(result.is_ok(), false);
    }
}
