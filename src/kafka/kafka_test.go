package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"testing"
)

func TestKafka01(t *testing.T) {
	topic := "test-topic"
	groupID := "orleans-consumer"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  groupID, // 配置消费者组
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	defer r.Close()

	fmt.Println("Start consuming messages...")

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Error reading message:", err)
		}
		fmt.Printf("Message at topic/partition/offset %v/%v/%v: %s = %s\n",
			m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
