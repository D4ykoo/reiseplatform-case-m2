package adapter

import (
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/ports"
	"log"
	"time"
)
import "github.com/IBM/sarama"

// InitProducer Instantiates a SyncProducer
//
// brokerUrls: e. g. ["localhost:8002"]
//
// Uses a random partitioner
func initProducer(brokerUrls []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()

	config.Version = sarama.DefaultVersion
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	config.Producer.Retry.Max = 5
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	conn, err := sarama.NewSyncProducer(brokerUrls, config)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

// SendEvent Send kafka message of a user management event with some content
// Sarama package is used. Fast and efficient IBM kafka library.
//
// No return error since it does not matter, just a local panic
func SendEvent(brokerUrls []string, topic string, event ports.PostEvent, content string) {
	producer, err := initProducer(brokerUrls)

	if err != nil {
		log.Panic("Can not create SyncProducer: " + err.Error())
		return
	}

	defer func(producer sarama.SyncProducer) {
		err := producer.Close()
		if err != nil {
			log.Panic(err.Error())
		}
	}(producer)

	messageString := event.String() + ": " + content

	msg := &sarama.ProducerMessage{
		Topic: topic,
		// Key: nil,
		Value: sarama.StringEncoder(messageString),
		// Headers:   nil,
		// Metadata:  nil,
		// Offset:    0,
		Partition: -1,
		Timestamp: time.Time{},
	}

	_, _, errSend := producer.SendMessage(msg)

	if err != nil {
		log.Panic("Can not push message: " + errSend.Error())
		return
	}

	return
}
