package main

import (
	"context"
	"encoding/json"

	"task/driver"
	"task/models"
	"task/service/executor"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

func main() {
	var err error
	// 初始化资源
	driver.InitEngine()
	// 初始化执行处理器
	service.InitExecutorHandler()
	// 监听执行事件
	go service.ListenExecuteEvent()
	// 监听执行事件消息 - kafka
	r := driver.CreateExecuteConsumer()
	for {
		var (
			ctx = context.Background()
			m   kafka.Message
		)
		// 接受消息
		if m, err = r.ReadMessage(ctx); err != nil {
			log.Errorf("receive executor message error: %v", err)
			break
		}
		log.Printf("receive executor message: %s", string(m.Value))
		// 解析消息
		event := new(models.ExecuteEvent)
		if err = json.Unmarshal(m.Value, &event); err != nil {
			log.Errorf("message unmarshal error: %v", err)
			break
		}
		// 接受执行事件消息 - channel
		service.ReceiveExecuteEvent(event)
	}
	if err = r.Close(); err != nil {
		log.Panic(err)
	}
}
