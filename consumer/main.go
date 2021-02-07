package main

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	c := KafkaConsumer()
	topics := []string{"topic-fullcycle"}
	c.SubscribeTopics(topics, nil)

	fmt.Println("kakfa consumer has been started")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(">>", string(msg.Value))
		}
	}
}

func KafkaConsumer() *ckafka.Consumer {

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "codepix",
		"auto.offset.reset": "earliest",
	}

	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		panic(err)
	}

	return c
}
