package internal

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func TestKafka() {
	writer := kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "example-topic",
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte("Hello, Kafka!"),
	})
	if err != nil {
		log.Fatalf("Failed to write message: %v", err)
	}

	log.Println("Message sent!")
}
