package main

import (
	"context"
	"encoding/json"

	"task/models"
	"task/pb/dispatch"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

func main() {
	// 创建kafka reader
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{models.KafkaAddress},
		GroupID:  models.KafkaConsumerDispatch,
		Topic:    models.KafkaTopicDispatch,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	// 监听kafka消息
	for {
		ctx := context.Background()
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}
		log.Printf("message at topic/partition/offset %v/%v/%v: %s = %s", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		event := new(models.DispatchEvent)
		if err = json.Unmarshal(m.Value, &event); err != nil {
			log.Errorf("message error: %v", err)
			return
		}
		// 调用调度rpc接口
		if _, err = dispatch.Dispatch(ctx, event.Event, event.ResourceId, event.DagId, event.ProcessorId); err != nil {
			return
		}
	}
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
