package executor

import (
	"context"
	"fmt"
	"time"

	"task/driver"
	"task/models"
	"task/pb/dispatch"
)

type CommonProcessor struct {
}

func (p CommonProcessor) Prepare(ctx context.Context, event *models.ExecuteEvent) error {
	var err error
	// 检查资源
	if event.Resource, err = driver.GetResource(ctx, event.ResourceId); err != nil {
		return err
	}
	// 检查处理状态
	if event.ResourceState, err = driver.GetProcessState(ctx, event.ResourceId, event.ProcessorId); err != nil {
		return err
	}
	event.ProcessState = event.ResourceState.ProcessState
	if event.ProcessState != models.ProcessStateReady {
		return fmt.Errorf("resource(%d) processor(%d) state(%d) is not ready", event.ResourceId, event.ProcessorId, event.ProcessState)
	}
	// 更新处理状态为运行中
	event.ProcessState = models.ProcessStateRunning
	if err = driver.UpdateProcessState(ctx, event.ResourceId, event.ProcessorId, event.ProcessState); err != nil {
		return err
	}
	return nil
}

func (p CommonProcessor) Execute(ctx context.Context, event *models.ExecuteEvent) error {
	// 空等
	time.Sleep(3 * time.Second)
	event.ProcessState = models.ProcessStateSuccess
	return nil
}

func (p CommonProcessor) After(ctx context.Context, event *models.ExecuteEvent) error {
	var err error
	// 更新处理结果
	if err = driver.UpdateProcessState(ctx, event.ResourceId, event.ProcessorId, event.ProcessState); err != nil {
		return err
	}
	// 当处理结果为finish和fail时 发起任务调度
	if event.ProcessState == models.ProcessStateSuccess || event.ProcessState == models.ProcessStateFail {
		if _, err = dispatch.Dispatch(ctx, models.DispatchEventProcessorDone, event.ResourceId, event.Resource.DagId, event.ProcessorId); err != nil {
			return err
		}
	}
	return nil
}
