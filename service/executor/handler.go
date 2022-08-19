package service

import (
	"context"

	"task/driver"
	"task/ecode"
	"task/models"

	log "github.com/sirupsen/logrus"
)

var ExecutorHandlerMap = make(map[string]ExecutorHandler)

// ExecutorHandler executor handler.
// 执行处理器 接口
type ExecutorHandler interface {
	Prepare(context.Context, *models.ExecuteEvent) error
	Execute(context.Context, *models.ExecuteEvent) error
	After(context.Context, *models.ExecuteEvent) error
}

// InitExecutorHandler init executor handler.
// 初始化执行处理器
func InitExecutorHandler() {
	ExecutorHandlerMap["processor1"] = CommonProcessor{}
	ExecutorHandlerMap["processor2"] = CommonProcessor{}
	ExecutorHandlerMap["processor3"] = CommonProcessor{}
	ExecutorHandlerMap["processor4"] = CommonProcessor{}
}

// getExecutorHandler get executor handler.
func getExecutorHandler(name string) (ExecutorHandler, error) {
	var (
		hdl ExecutorHandler
		ok  bool
	)
	if hdl, ok = ExecutorHandlerMap[name]; !ok {
		return nil, ecode.ExecutorHandlerNotFound
	}
	return hdl, nil
}

// processExecuteEvent process execute event.
func processExecuteEvent(ctx context.Context, event *models.ExecuteEvent) error {
	log.Infof("resource:%d processor:%d start execute...", event.ResourceId, event.ProcessorId)
	var (
		hdl ExecutorHandler
		err error
	)
	// 检查处理单元
	if event.Processor, err = driver.GetProcessor(ctx, event.ProcessorId); err != nil {
		return err
	}
	// 获取处理器
	if hdl, err = getExecutorHandler(event.Processor.Handler); err != nil {
		return err
	}
	// 处理准备
	if err = hdl.Prepare(ctx, event); err != nil {
		return err
	}
	// 处理执行
	if err = hdl.Execute(ctx, event); err != nil {
		return err
	}
	// 处理结束
	if err = hdl.After(ctx, event); err != nil {
		return err
	}
	return nil
}

// handleProcessError handle process error.
func handleProcessError(ctx context.Context, event *models.ExecuteEvent, err error) {
	// 更新处理状态 处理错误
	event.ProcessState = models.ProcessStateError
	event.ProcessMsg = err.Error()
	if event.ResourceProcessState != nil {
		_ = driver.UpdateResourceProcessState(ctx, event.ResourceProcessState.Id, event.ProcessState, event.ProcessCnt, event.ProcessMsg)
	} else {
		_ = driver.AddResourceProcessState(ctx, event.ResourceId, event.ProcessorId, event.ProcessState, event.ProcessCnt, event.ProcessMsg)
	}
	return
}
