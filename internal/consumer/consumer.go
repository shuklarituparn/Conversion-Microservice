// consumer/consumer.go
package consumer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

var (
	kafkaConsumer *kafka.Consumer
)

func NewConsumer(bootstrapServers string, groupID string, topics []string) (*kafka.Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  bootstrapServers,
		"group.id":           groupID,
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	})
	if err != nil {
		return nil, err
	}

	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		err := c.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	kafkaConsumer = c
	return kafkaConsumer, nil
}

func Close(c *kafka.Consumer) {
	err := c.Close()
	if err != nil {
		log.Println(err)
	}
}
