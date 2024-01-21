use rdkafka::{
    config::RDKafkaLogLevel,
    consumer::{Consumer, StreamConsumer},
    error::KafkaError,
    ClientConfig, Message,
};
use std::sync::mpsc::Sender;
use tracing::warn;

#[derive(Debug)]
pub struct KafkaMsg {
    pub topic: String,
    pub payload: String,
}

pub async fn subscribe(s: Sender<KafkaMsg>, consumer: StreamConsumer) {
    loop {
        match consumer.recv().await {
            Err(e) => warn!("Kafka error: {}", e),
            Ok(m) => {
                match m.payload_view::<str>() {
                    Some(Ok(pay)) => s
                        .send(KafkaMsg {
                            topic: m.topic().to_string(),
                            payload: pay.to_string(),
                        })
                        .unwrap(),
                    Some(Err(e)) => {
                        warn!("Error while deserializing message payload: {:?}", e);
                    }
                    None => (),
                };
            }
        }
    }
}

pub fn create_consumer(server: &str, topic: &[&str]) -> Result<StreamConsumer, KafkaError> {
    let consumer: StreamConsumer = ClientConfig::new()
        .set("group.id", "123")
        .set("bootstrap.servers", server)
        .set("enable.partition.eof", "false")
        .set("session.timeout.ms", "6000")
        .set("enable.auto.commit", "true")
        .set("auto.commit.interval.ms", "1000")
        .set("enable.auto.offset.store", "false")
        //  .set("auto.offset.reset", "earliest")
        .set_log_level(RDKafkaLogLevel::Debug)
        .create()?;
    consumer.subscribe(topic)?;
    Ok(consumer)
}

