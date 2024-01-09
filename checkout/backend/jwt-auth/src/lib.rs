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
        token,
        &DecodingKey::from_secret(jwt_secret.as_bytes()),
        &val,
    );

    println!("{:?}", decoded);

    match decoded {
        Ok(_) => Ok(true),
        Err(e) => Err(e),
    }
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
        let token = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjY4NjYyMDgwMCwidXNlcl9pZCI6MSwidXNlcm5hbWUiOiJ0ZXN0In0.j1McK-wmsGqNWQ9ir_DqT90uFFvIt7zjYbABe1Lf0iJ8wbZeaqedbTr5v2UB43ZLCySEwTx3QzROTXltzTIXoA";
        let result = validate_jwt(token);
        assert!(result.is_ok());
    }

    #[test]
    fn token_invalid() {
        let token = "eyJhbGciOIJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjY4NjYyMDgwMCwidXNlcm5hbWUiOiJ0ZXN0IiwidXNlcl9pZCI6MX0.mYk1vDz6qRcV-RjnFWdNT8uYHno3A7A7VklLerl1fQImiMGmuu-GlG-EmrI2Wf4FcIL_Svgrp2azZZf7B8DBmg";
        let result = validate_jwt(token);
        assert!(!result.is_ok());
    }
}
