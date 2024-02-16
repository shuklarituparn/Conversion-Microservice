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
	deliveryChan := make(chan kafka.Event)
	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value),
	}, deliveryChan)

	if err != nil {
		fmt.Println("error producing message", err)
		return err
	} else {
		go func() {
			e := <-deliveryChan
			m := e.(*kafka.Message)
			if m.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
			} else {
				fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
					*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
			}
		}()
	}

	return nil
}
