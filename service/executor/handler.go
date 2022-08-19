package executor

import (
	"context"

	"task/driver"
	"task/models"

	log "github.com/sirupsen/logrus"
)

type Handler interface {
	Prepare(context.Context, *models.ExecuteEvent) error
	Execute(context.Context, *models.ExecuteEvent) error
	After(context.Context, *models.ExecuteEvent) error
}

var HandlerMap = make(map[string]Handler)

func InitHandler() {
	HandlerMap["processor1"] = CommonProcessor{}
	HandlerMap["processor2"] = CommonProcessor{}
	HandlerMap["processor3"] = CommonProcessor{}
	HandlerMap["processor4"] = CommonProcessor{}
	HandlerMap["processor5"] = CommonProcessor{}
	HandlerMap["processor6"] = CommonProcessor{}
	HandlerMap["processor7"] = CommonProcessor{}
}

// run run.
func run(event *models.ExecuteEvent) error {
	log.Infof("resource:%d processor:%d start run...", event.ResourceId, event.ProcessorId)
	var (
		ctx     = context.Background()
		handler Handler
		ok      bool
		err     error
	)
	// 检查处理单元
	// todo 后期可以使用redis缓存 减少mysql查询次数
	if event.Processor, err = driver.GetProcessor(ctx, event.ProcessorId); err != nil {
		return err
	}
	// 获取处理器
	if handler, ok = HandlerMap[event.Processor.Handler]; !ok {
		return err
	}
	if err = handler.Prepare(ctx, event); err != nil {
		return err
	}
	if err = handler.Execute(ctx, event); err != nil {
		return err
	}
	if err = handler.After(ctx, event); err != nil {
		return err
	}
	return nil
}

// handleProcessError handle process error.
func handleProcessError(event *models.ExecuteEvent, err error) {
	// todo 更新处理状态
	return
}
