package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/D4ykoo/travelplatform-case-m2/usermanagement/adapter/kafka/dto"
	"github.com/IBM/sarama"
	"log"
	"os"
	"time"
)

// initProducer Instantiates a SyncProducer
//
// brokerUrls: e. g. ["localhost:9092"]
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

func Publish(message dto.UserEventMessage) {
	brokerUrls := []string{os.Getenv("BROKERS")}
	topic := os.Getenv("TOPIC")
	producer, err := initProducer(brokerUrls)

	if err != nil {
		log.Print("Can not create SyncProducer: " + err.Error())
		return
	}

	defer func(producer sarama.SyncProducer) {
		err := producer.Close()
		if err != nil {
			log.Print(err.Error())
		}
	}(producer)

	marshalMsg, _ := json.Marshal(message)

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.ByteEncoder(marshalMsg),
		Partition: -1,
		Timestamp: time.Time{},
	}

	_, _, errSend := producer.SendMessage(msg)

	if err != nil {
		fmt.Printf("Can not push message: %s\n", errSend.Error())
		return
	}
}
