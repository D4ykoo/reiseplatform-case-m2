use diesel::{prelude::*};
use crate::schema::users;
use serde::Deserialize;
use serde::Serialize;

#[derive(Queryable, Selectable, Debug, Deserialize, Serialize, Clone)]
#[diesel(table_name = cart)]
#[diesel(check_for_backend(diesel::pg:Pg))]
pub struct cart{
    pub id: i8,
}