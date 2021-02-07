package main

import (
	"sync"

	"github.com/codeedu/imersao/codepix-go/application/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

var wg sync.WaitGroup

func main() {
	deliveryChan := make(chan ckafka.Event)
	producer := KafkaProducer()

	topic := "topic-fullcycle"
	msg := "Curso Imersão FullCycle - Desafio 2"

	err := kafka.Publish(msg, topic, producer, deliveryChan)
	if err != nil {
		panic(err)
	}

	wg.Add(1)
	go kafka.DeliveryReport(deliveryChan)
	wg.Wait()
}

func KafkaProducer() *ckafka.Producer {

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}
	producer, err := ckafka.NewProducer(configMap)

	if err != nil {
		panic(err)
	}
	return producer
}
