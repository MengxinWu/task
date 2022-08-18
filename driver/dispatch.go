package driver

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"task/models"
	"time"
)

func InsertDispatchEvent(event *models.DispatchEvent) error {
	// to produce messages
	topic := "dispatch-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	dataJson, _ := json.Marshal(event)

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: dataJson},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	return nil
}
