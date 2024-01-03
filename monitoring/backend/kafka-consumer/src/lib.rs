use std::sync::mpsc::Sender;
use std::{thread, time};
use monitoring_db::{add_user_event, establish_connection, model::NewUserEvent};
use rdkafka::{
    config::RDKafkaLogLevel,
    consumer::{CommitMode, Consumer, StreamConsumer},
    error::KafkaError,
    message::Headers,
    ClientConfig, Message,
};

use tracing::warn;

pub async fn sub(s: Sender<i32>) {
    for i in 0..10 {
        if let Err(_) = s.send(i) {
            println!("receiver dropped");
            return;
        }
        println!("BIN SA");
        let ten_millis = time::Duration::from_millis(1000);
        thread::sleep(ten_millis);
    }
}

pub async fn subscribe(consumer: StreamConsumer, server: &str, topic: &str) {
    let conn = &mut establish_connection();
    loop {
        match consumer.recv().await {
            Err(e) => warn!("Kafka error: {}", e),
            Ok(m) => {
                let payload = match m.payload_view::<str>() {
                    None => "",
                    Some(Ok(s)) => s,
                    Some(Err(e)) => {
                        println!("Error while deserializing message payload: {:?}", e);
                        ""
                    }
                };

                println!("key: '{:?}', payload: '{}', topic: {}, partition: {}, offset: {}, timestamp: {:?}",
                      m.key(), payload, m.topic(), m.partition(), m.offset(), m.timestamp());
                if let Some(headers) = m.headers() {
                    for header in headers.iter() {
                        println!("  Header {:#?}: {:?}", header.key, header.value);
                    }
                }
                add_user_event(
                    conn,
                    NewUserEvent::new(
                        "type_".into(),
                        Some("log".into()),
                        chrono::offset::Utc::now(),
                    ),
                )
                .unwrap();
                consumer.commit_message(&m, CommitMode::Async).unwrap();
            }
        };
    }
}

pub fn add_to_db(payload: &str) {
    match payload {
        "travelmanagenet" => {}
        "usermanagement" => {}
        "checkout" => {}
        &_ => {}
    }
}

pub async fn start_and_subscribe(server: &str, topic: &str) {
    let consumer_res: Result<StreamConsumer, KafkaError> = create_consumer(server, topic);
    if consumer_res.is_err() {
        warn!(
            "Unable to subscribe Kafka Server. No message will be received. Error: {}",
            consumer_res.err().unwrap().to_string()
        )
    } else {
        let conn = &mut establish_connection();
        let consumer = consumer_res.unwrap();
        loop {
            match consumer.recv().await {
                Err(e) => println!("Kafka error: {}", e),
                Ok(m) => {
                    let payload = match m.payload_view::<str>() {
                        None => "",
                        Some(Ok(s)) => s,
                        Some(Err(e)) => {
                            println!("Error while deserializing message payload: {:?}", e);
                            ""
                        }
                    };
                    match m.topic() {
                        "dadf" => {}
                        "dadf" => {}
                        "dadf" => {}
                        &_ => {}
                    }
                    println!("key: '{:?}', payload: '{}', topic: {}, partition: {}, offset: {}, timestamp: {:?}",
                      m.key(), payload, m.topic(), m.partition(), m.offset(), m.timestamp());
                    if let Some(headers) = m.headers() {
                        for header in headers.iter() {
                            println!("  Header {:#?}: {:?}", header.key, header.value);
                        }
                    }
                    add_user_event(
                        conn,
                        NewUserEvent::new(
                            "type_".into(),
                            Some("log".into()),
                            chrono::offset::Utc::now(),
                        ),
                    )
                    .unwrap();
                    consumer.commit_message(&m, CommitMode::Async).unwrap();
                }
            };
        }
    }
}

pub fn create_consumer(server: &str, topic: &str) -> Result<StreamConsumer, KafkaError> {
    let consumer: StreamConsumer = ClientConfig::new()
        .set("group.id", "123")
        .set("bootstrap.servers", server)
        .set("enable.partition.eof", "false")
        .set("session.timeout.ms", "6000")
        .set("enable.auto.commit", "true")
        .set("auto.commit.interval.ms", "1000")
        .set("enable.auto.offset.store", "false")
        .set_log_level(RDKafkaLogLevel::Debug)
        .create()?;

    consumer.subscribe(&[topic])?;
    Ok(consumer)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let a = create_consumer("", "");
        assert_eq!(1, 4);
    }
}
