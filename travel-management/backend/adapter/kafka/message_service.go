package kafka

import (
	"encoding/json"
	"log"
	"time"

	"github.com/IBM/sarama"
)

type MessageService struct {
	topic       string
	KafkaClient sarama.AsyncProducer
}

func NewMsgService(url, topic string) (MessageService, error) {

	config := sarama.NewConfig()
	config.Version = sarama.DefaultVersion
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	config.Producer.Retry.Max = 5
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = false

	conn, err := tryToGetConnection(url, config)

	return MessageService{
		topic:       topic,
		KafkaClient: conn,
	}, err
}

func tryToGetConnection(url string, config *sarama.Config) (sarama.AsyncProducer, error) {
	var timeToSleep = 2 * time.Second
	var connection sarama.AsyncProducer
	var err error
	for {
		connection, err = sarama.NewAsyncProducer([]string{url}, config)

		// Exit - Connection established
		if err == nil {
			break
		}

		// Exit - Connection establishment is not possible
		if timeToSleep.Seconds() >= 10 {
			break
		}

		// Failure - Retry
		if err != nil && timeToSleep.Seconds() < 10 {
			log.Println("Failed to connect to Kafa Broker. Retry in " + string(timeToSleep.String()))
			time.Sleep(timeToSleep)
			timeToSleep = timeToSleep * 2
		}

	}
	return connection, err
}

func (service MessageService) PublishAsJSON(obj interface{}) error {
	json, err := json.Marshal(obj)

	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     service.topic,
		Value:     sarama.ByteEncoder(json),
		Partition: -1,
		Timestamp: time.Time{},
	}

	service.KafkaClient.Input() <- msg
	service.KafkaClient.BeginTxn()
	service.KafkaClient.CommitTxn()
	return err

}
