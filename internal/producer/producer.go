package producer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	producer *kafka.Producer
)

func NewProducer(bootstrapServers string) (*kafka.Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
	})
	if err != nil {
		fmt.Println("Error creating Kafka producer", err)
		return nil, err

	}
	fmt.Printf("Created producer %v\n", p)

	producer = p
	return producer, nil
}

func ProduceNewMessage(p *kafka.Producer, topic string, value string) error {
	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value),
	}, nil)

	if err != nil {
		fmt.Println("error producing message", err)
		return err
	}

	return nil
}
