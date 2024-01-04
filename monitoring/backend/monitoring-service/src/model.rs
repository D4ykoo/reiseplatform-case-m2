use serde::Deserialize;
use std::error::Error;
use std::fmt;

#[derive(Deserialize, Debug)]
pub struct Period {
    pub from: Option<String>,
    pub to: Option<String>,
}

#[derive(Debug)]
pub struct TokenError {
    details: String,
}

impl TokenError {
    pub fn new(msg: &str) -> TokenError {
        TokenError {
            details: msg.to_string(),
        }
    }
}

impl fmt::Display for TokenError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.details)
    }
}

impl Error for TokenError {
    fn description(&self) -> &str {
        &self.details
    }
}
