use serde::Deserialize;

#[derive(Deserialize, Debug)]
pub struct Period {
    pub from: Option<String>,
    pub to: Option<String>,
}
