package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"task/models"
	"task/pb/dispatch"
)

func main() {

	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka:9092"},
		GroupID:  "consumer-group-dispatch",
		Topic:    "dispatch-topic",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		dispatchEvent := new(models.DispatchEvent)
		if err = json.Unmarshal(m.Value, &dispatchEvent); err != nil {
			log.Fatalf("message error: %v", err)
			return
		}
		if _, err = dispatch.Dispatch(context.Background(), dispatchEvent.Event, dispatchEvent.ResourceId,
			dispatchEvent.DagId, dispatchEvent.ProcessorId); err != nil {
			return
		}
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}

}
