package kafka

import (
	"encoding/json"
	"time"

	"github.com/IBM/sarama"
)

type MessageService struct {
	topic       string
	KafkaClient sarama.SyncProducer
}

func NewMsgService(url, topic string) (MessageService, error) {

	config := sarama.NewConfig()
	config.Version = sarama.DefaultVersion
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	config.Producer.Retry.Max = 5
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	conn, err := sarama.NewSyncProducer([]string{url}, config)

	return MessageService{
		topic:       topic,
		KafkaClient: conn,
	}, err
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

	service.KafkaClient.SendMessage(msg)

	return err

}
