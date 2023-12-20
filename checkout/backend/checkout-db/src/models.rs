use diesel::prelude::*;
use crate::schema::cart;
use serde::Deserialize;
use serde::Serialize;


#[derive(Queryable, Selectable, Debug, Deserialize, Serialize, Clone)]
#[diesel(table_name = cart)]
#[diesel(check_for_backend(diesel::pg::Pg))]
pub struct Cart {
    pub id: i32,
    pub paid: Option<bool>,
    pub payment_method: Option<String>,
    pub offers: Option<Vec<Option<i32>>>
}

#[derive(Insertable)]
#[diesel(table_name = cart)]
pub struct NewCart<'a> {
    pub paid: Option<bool>,
    pub payment_method: Option<&'a str>,
    pub offers: Option<Vec<Option<i32>>>,
}

impl<'a> NewCart<'a> {
    pub fn create(paid: Option<bool>, payment_method: Option<&'a str>, offers: Option<Vec<Option<i32>>>) -> Self{
        NewCart {
            paid,
            payment_method,
            offers,
        }
    }
}