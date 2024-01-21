//! Message Service for producing messages and sending them
//!
//! Uses async kafka, so it is non blocking.
//! Uses only one broker defined in the .env file.

use dotenvy::dotenv;
use rdkafka::config::{ClientConfig, RDKafkaLogLevel};
use rdkafka::producer::{FutureProducer, FutureRecord};
use serde::{Serialize, Deserialize};
use serde_json::to_vec;
use std::env;
use std::time::Duration;
use chrono::{DateTime, Utc};

#[derive(Deserialize, Serialize, Debug)]
pub struct EventMessage {
    #[serde(rename = "type")] 
    pub type_: String,
    pub log: String,
    pub time: DateTime<Utc>
}


#[derive(Clone)]
pub struct MessageProducer {
    pub producer: Option<FutureProducer>,
}

impl MessageProducer {
    pub fn init_message_producer(&mut self, kafka_url: &str) {
        dotenv().ok();
        println!("Kafka url: {}", kafka_url);
        let future_producer: FutureProducer = ClientConfig::new()
            .set("bootstrap.servers", kafka_url)
            .set("message.timeout.ms", "5000")
            .set_log_level(RDKafkaLogLevel::Debug)
            .create()
            .expect("producer creation error");

        self.producer = Some(future_producer)
    }

    pub async fn send_message(&self, payload: &str) {

        let complete_message: EventMessage = EventMessage{ 
            type_: "Checkout".to_owned(),
            log: payload.to_owned(),
            time: Utc::now()
         };

        if let Some(producer) = &self.producer {
            let _ = producer
                .send(
                    FutureRecord::to(env::var("TOPIC").unwrap().as_str())
                        .payload(&serde_json::to_vec(&complete_message).unwrap())
                        .key("checkout"),
                    Duration::from_secs(0),
                )
                .await;
        }
    }
}
