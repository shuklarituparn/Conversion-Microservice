package producer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	kafkaProducer *KafkaProducer
)

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer() (*KafkaProducer, error) {
	if kafkaProducer != nil {
		return kafkaProducer, nil
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		return nil, err
	}

	kafkaProducer = &KafkaProducer{
		producer: p,
	}

	return kafkaProducer, nil
}

func (p *KafkaProducer) Close() {
	p.producer.Close()
}

func (p *KafkaProducer) ProduceMessage(topic string, message string) error {
	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, deliveryChan)

	if err != nil {
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	return nil
}
