use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct Claims {
    pub username: String,
    pub user_id: usize,
    pub iat: usize,
}
