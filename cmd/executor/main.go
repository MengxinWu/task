package main

import (
	"context"
	"encoding/json"

	"task/driver"
	"task/models"
	service "task/service/executor"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

func main() {
	var err error
	// 初始化执行处理器
	service.InitHandler()
	// 开始监听执行事件
	go service.InitExecuteEvent()
	// 开始监听kafka消息(执行)
	r := driver.CreateExecuteConsumer()
	for {
		var (
			ctx = context.Background()
			m   kafka.Message
		)
		if m, err = r.ReadMessage(ctx); err != nil {
			log.Errorf("read message error: %v", err)
			break
		}
		log.Printf("read message success with topic(%s) partition(%d) offset(%d) key(%s) msg(%s)", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		event := new(models.ExecuteEvent)
		if err = json.Unmarshal(m.Value, &event); err != nil {
			log.Errorf("message unmarshal error: %v", err)
			break
		}
		service.ReceiveExecuteEvent(event)
	}
	if err = r.Close(); err != nil {
		log.Panic(err)
	}
}
