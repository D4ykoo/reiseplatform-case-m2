//! Message Service for producing messages and sending them
//!
//! Uses async kafka, so it is non blocking.
//! Uses only one broker defined in the .env file.

use dotenvy::dotenv;
use rdkafka::config::{ClientConfig, RDKafkaLogLevel};
use rdkafka::producer::{FutureProducer, FutureRecord};
use serde::Serialize;
use std::env;
use std::time::Duration;

#[derive(Serialize, Debug)]
pub struct EventMessage {
    pub event: String,
    pub data: String,
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
        if let Some(producer) = &self.producer {
            let _ = producer
                .send(
                    FutureRecord::to(env::var("TOPIC").unwrap().as_str())
                        .payload(&format!("Checkout {:?}", payload))
                        .key("checkout"),
                    Duration::from_secs(0),
                )
                .await;
        }
    }
}
