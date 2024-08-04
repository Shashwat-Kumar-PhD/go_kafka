package cmd

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

// StartKafkaProducer start Kafka producer
func StartKafkaProducer(brokers []string, topic string) {
	log.Print("Kafka prducer is started...")
	// Kafka configuration
	config := sarama.NewConfig()
	config.ClientID = "KafkaProducer1"
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start Sarama producer: %v", err)
	}
	defer producer.Close()

	// Create message
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}
	// Send message
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", "test_topic", partition, offset)
}
