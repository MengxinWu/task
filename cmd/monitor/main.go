package main

import (
	"context"
	"encoding/json"

	"task/driver"
	"task/models"
	"task/pb/dispatch"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

func main() {
	var err error
	// 监听调度事件消息 - kafka
	r := driver.CreateDispatchConsumer()
	for {
		var (
			ctx = context.Background()
			m   kafka.Message
		)
		// 接受消息
		if m, err = r.ReadMessage(ctx); err != nil {
			log.Errorf("receive dispatch message error: %v", err)
			break
		}
		log.Printf("receive dispatch message: %s", string(m.Value))
		// 解析消息
		event := new(models.DispatchEvent)
		if err = json.Unmarshal(m.Value, &event); err != nil {
			log.Errorf("message unmarshal error: %v", err)
			break
		}
		// 调用调度接口
		if _, err = dispatch.Dispatch(ctx, event.Event, event.ResourceId, event.DagId, event.ProcessorId); err != nil {
			return
		}
	}
	if err = r.Close(); err != nil {
		log.Panic(err)
	}
}
