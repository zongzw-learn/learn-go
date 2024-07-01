package main

// https://developer.confluent.io/get-started/go/#introduction

// $ wget https://dlcdn.apache.org/kafka/3.7.0/kafka_2.13-3.7.0.tgz
// $ tar zxf kafka_2.13-3.7.0.tgz
// $ cd kafka_2.13-3.7.0/bin
// $ ./kafka-topics.sh --create --topic mytopic --bootstrap-server localhost:9092
import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var topic string = "topic5"

func main() {

	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	stop := make(chan struct{})
	createTopic()
	go startConsumer(stop)
	go startProducer()

	sig := <-sigchan
	log.Printf("Caught signal %v: terminating\n", sig)
	close(stop)

}

func createTopic() {
	c, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		log.Fatalf("cannot create admin client: %s", err.Error())
		os.Exit(1)
	}
	r, err := c.CreateTopics(context.Background(), []kafka.TopicSpecification{{
		Topic:         topic,
		NumPartitions: 1,
	}})
	if err != nil {
		log.Fatalf("failed to create topic: %s", err.Error())
		os.Exit(1)
	}
	log.Printf("%v", r)
}

func startProducer() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		// User-specific properties that you must set
		"bootstrap.servers": "localhost:9092",

		// Fixed properties
		// "acks": "all",
	})

	if err != nil {
		log.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	users := [...]string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
	items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}

	for n := 0; n < 10; n++ {
		key := users[rand.Intn(len(users))]
		data := items[rand.Intn(len(items))]
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          []byte(data),
		}, nil)
	}

	// Wait for all messages to be delivered
	p.Flush(15 * 1000)
	p.Close()
}

func startConsumer(stop chan struct{}) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		// User-specific properties that you must set
		"bootstrap.servers": "localhost:9092",

		// Fixed properties
		"group.id": "kafka-go-getting-started",
		// "auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}
	defer c.Close()

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Printf("failed to subscribe topic: %s\n", err.Error())
		os.Exit(1)
	}

	// Process messages
	for {
		select {
		case <-stop:
			return
		default:
			ev, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				if err.(kafka.Error).Code() == kafka.ErrTimedOut {
					continue
				}
				log.Printf("failed to read message from topic: %s", err.Error())
			}
			log.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
		}
	}

}
