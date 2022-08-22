package service

import (
	"context"

	"task/models"

	log "github.com/sirupsen/logrus"
)

// executeEventCh execute event channel.
// 执行事件通道
var (
	executeEventCh = make(chan *models.ExecuteEvent, 100000)
	workerNum      = 100
)

// ReceiveExecuteEvent receive execute event.
func ReceiveExecuteEvent(event *models.ExecuteEvent) {
	executeEventCh <- event
	return
}

// ListenExecuteEvent listen execute event.
func ListenExecuteEvent() {
	for i := 0; i < workerNum; i++ {
		workerIdx := i
		go func() {
			for {
				select {
				case event := <-executeEventCh:
					log.Infof("worker num: %d, channel len: %d", workerIdx, len(executeEventCh))
					var ctx = context.Background()
					// 处理执行事件
					if err := processExecuteEvent(ctx, event); err != nil {
						// 处理错误处理
						handleProcessError(ctx, event, err)
					}
				}
			}
		}()
	}
}
