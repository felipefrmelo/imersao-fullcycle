package consumer

import (
	"fmt"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Subscribe struct {
	consumer *kafka.Consumer
}

func NewKafkaConsumer() *Subscribe {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	return &Subscribe{c}
}

func (c *Subscribe) SubscribeTopics(topics []string) {
	c.consumer.SubscribeTopics(topics, nil)

	for {
		msg, err := c.consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
