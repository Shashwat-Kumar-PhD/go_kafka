package cmd

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

// StartKafkaConsumer start kafka consumer
func StartKafkaConsumer(brokers []string, topic string) {
	log.Print("Kafka consumer is started...")
	config := sarama.NewConfig()
	config.ClientID = "Consumer1"
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	// start consumer
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatal("error in creating Kafka consumer", err)
	}
	// Consume messages
	defer master.Close()
	consumer, err := master.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start consumer for partition: %v", err)
	}
	defer consumer.Close()
	for {
		select {
		case msg := <-consumer.Messages():
			fmt.Printf("Consumed message: %s\n", string(msg.Value))
		case err := <-consumer.Errors():
			log.Printf("Error consuming message: %v", err)
		}
	}
}
