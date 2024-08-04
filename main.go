package main

import (
	"go_kafka/cmd"
	"log"
	"os"

	"github.com/IBM/sarama"
	"github.com/alecthomas/kingpin/v2"
)

var (
	startProducer = kingpin.Command("start-producer", "Start Kakfa producer")
	startConsumer = kingpin.Command("start-consumer", "Start Kafka consumer")
	topics        = kingpin.Flag("topics", "commad seprated list of Kafka topics").Default("go-kafka").String()
	brokers       = kingpin.Flag("broker", "list of brokers").Default("localhost:9092").Envar("KAFKA_BROKERS").Strings()
)

func main() {
	sarama.Logger = log.New(os.Stdout, "kafka", log.LstdFlags)
	switch kingpin.Parse() {
	case startProducer.FullCommand():
		log.Print("Producer command is selected")
		cmd.StartKafkaProducer(*brokers, *topics)
	case startConsumer.FullCommand():
		log.Print("consumer commad is selected")
		cmd.StartKafkaConsumer(*brokers, *topics)
	}
}
