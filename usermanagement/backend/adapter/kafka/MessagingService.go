package kafka

import (
	"encoding/json"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/kafka/dto"
	"github.com/IBM/sarama"
	"log"
	"os"
	"time"
)

type MessageService struct {
	Publisher sarama.AsyncProducer
}

func NewMessageService(brokerUrls []string) (MessageService, error) {

	config := sarama.NewConfig()
	config.Version = sarama.DefaultVersion
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	config.Producer.Retry.Max = 5
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = false

	conn, err := tryToGetConnection(brokerUrls, config)

	return MessageService{
		Publisher: conn,
	}, err
}

func tryToGetConnection(url []string, config *sarama.Config) (sarama.AsyncProducer, error) {
	var timeToSleep = 2 * time.Second
	var connection sarama.AsyncProducer
	var err error
	for {
		connection, err = sarama.NewAsyncProducer(url, config)

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

func (service MessageService) Publish(message dto.UserEventMessage) error {
	packedJson, err := json.Marshal(message)

	if err != nil {
		return err
	}
	topic := os.Getenv("TOPIC")

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.ByteEncoder(packedJson),
		Partition: -1,
		Timestamp: time.Time{},
	}

	service.Publisher.Input() <- msg
	errBegin := service.Publisher.BeginTxn()
	if errBegin != nil {
		return errBegin
	}
	errCommit := service.Publisher.CommitTxn()
	if errCommit != nil {
		return errCommit
	}

	return err
}
