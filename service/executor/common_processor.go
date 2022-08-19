package service

import (
	"context"
	"math/rand"
	"time"

	"task/driver"
	"task/ecode"
	"task/models"
	"task/pb/dispatch"
	"task/utils"

	log "github.com/sirupsen/logrus"
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
	if event.ResourceProcessState, err = driver.GetResourceProcessState(ctx, event.ResourceId, event.ProcessorId); err != nil {
		return err
	}
	event.ProcessState = event.ResourceProcessState.ProcessState
	event.ProcessCnt = event.ResourceProcessState.ProcessCnt
	if event.ProcessState != models.ProcessStateReady {
		log.Errorf("resource(%d) processor(%d) state(%d) is not ready", event.ResourceId, event.ProcessorId, event.ProcessState)
		return ecode.ProcessStateWrong
	}
	// 更新处理状态 处理中
	event.ProcessState = models.ProcessStateRunning
	event.ProcessCnt += 1
	if err = driver.UpdateResourceProcessState(ctx, event.ResourceProcessState.Id, event.ProcessState, event.ProcessCnt, ""); err != nil {
		return err
	}
	return nil
}

func (p CommonProcessor) Execute(_ context.Context, event *models.ExecuteEvent) error {
	// 测试任务
	time.Sleep(3 * time.Second)
	// 80%概率成功
	event.ProcessState = models.ProcessStateSuccess
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(100) > 80 {
		event.ProcessState = models.ProcessStateFail
	}
	return nil
}

func (p CommonProcessor) After(ctx context.Context, event *models.ExecuteEvent) error {
	var err error
	// 更新处理结果
	if err = driver.UpdateResourceProcessState(ctx, event.ResourceProcessState.Id, event.ProcessState, event.ProcessCnt, event.ProcessMsg); err != nil {
		return err
	}
	// 当处理结果为success和fail时 发起任务调度
	if utils.IntInSlice(event.ProcessState, []int{models.ProcessStateSuccess, models.ProcessStateFail}) {
		if _, err = dispatch.Dispatch(ctx, models.DispatchEventProcessorDone, event.ResourceId, event.Resource.DagId, event.ProcessorId); err != nil {
			return err
		}
	}
	return nil
}
