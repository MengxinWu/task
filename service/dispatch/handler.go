package service

import (
	"context"
	"fmt"

	"task/driver"
	"task/models"

	log "github.com/sirupsen/logrus"
)

type DispatchHandler interface {
	Prepare(context.Context, *models.DispatchEvent) error
	Compute(context.Context, *models.DispatchEvent) error
	After(context.Context, *models.DispatchEvent) error
}

var DispatchHandlerMap = make(map[string]DispatchHandler)

func InitDispatchHandler() {
	DispatchHandlerMap["resource_add"] = ResourceAddHandler{}
	DispatchHandlerMap["processor_done"] = ProcessorDoneHandler{}
}

type ResourceAddHandler struct {
}

func (h ResourceAddHandler) Prepare(ctx context.Context, event *models.DispatchEvent) error {
	var err error
	if event.Resource, err = driver.GetResource(ctx, event.ResourceId); err != nil {
		return err
	}
	if event.Dag, err = driver.GetDag(ctx, event.DagId); err != nil {
		return err
	}
	return nil
}

func (h ResourceAddHandler) Compute(ctx context.Context, event *models.DispatchEvent) error {
	var err error
	// 生成DAG图
	if event.Graph, err = models.GenerateGraph(event.Dag.Config); err != nil {
		return err
	}
	// 计算根处理单元
	for _, node := range event.Graph {
		if node.Parents == nil {
			event.ExecutorList = append(event.ExecutorList, int64(node.ProcessorId))
			// 设置处理状态为等待执行
			if err = driver.AddResourceProcessState(ctx, event.ResourceId, node.ProcessorId, models.ProcessStateReady, 0, ""); err != nil {
				return err
			}
		}
	}
	return nil
}

func (h ResourceAddHandler) After(_ context.Context, event *models.DispatchEvent) error {
	var (
		executeEvents []*models.ExecuteEvent
		err           error
	)
	if len(event.ExecutorList) == 0 {
		return nil
	}
	// 进入执行消息队列 等待执行
	for _, processorId := range event.ExecutorList {
		executeEvents = append(executeEvents, &models.ExecuteEvent{
			ResourceId:  event.ResourceId,
			ProcessorId: int(processorId),
		})
	}
	if err = driver.SendExecuteEventMsg(executeEvents); err != nil {
		return err
	}
	return nil
}

type ProcessorDoneHandler struct {
}

func (h ProcessorDoneHandler) Prepare(ctx context.Context, event *models.DispatchEvent) error {
	var err error
	if event.Resource, err = driver.GetResource(ctx, event.ResourceId); err != nil {
		return err
	}
	if event.ProcessState, err = driver.GetProcessState(ctx, event.ResourceId, event.ProcessorId); err != nil {
		return err
	}
	return nil
}

func (h ProcessorDoneHandler) Compute(ctx context.Context, event *models.DispatchEvent) error {
	var (
		err      error
		children []*models.Node
	)
	if event.ProcessState.ProcessorCnt >= 4 {
		return fmt.Errorf("process state cnt > 4 (%d)", event.ProcessState.ProcessorCnt)
	}
	// 失败任务重试
	if event.ProcessState.ProcessState != 400 {
		log.Printf("process execute unsuccess %d, %d, retrying...", event.ResourceId, event.ProcessorId)
		event.ExecutorList = append(event.ExecutorList, int64(event.ProcessorId))
		// 设置处理状态为等待执行
		if err = driver.AddProcessState(ctx, event.ResourceId, event.ProcessorId, models.ProcessStateReady); err != nil {
			return err
		}
	}
	// 寻找下游节点的上游节点
	// 生成DAG图
	if event.Graph, err = models.GenerateGraph(event.Dag.Config); err != nil {
		return err
	}
	// 寻找children节点
	for _, node := range event.Graph {
		if node.ProcessorId == event.ProcessorId {
			children = node.Children
		}
	}
	if children == nil {

	}
	return nil
}

func (h ProcessorDoneHandler) After(_ context.Context, event *models.DispatchEvent) error {
	var (
		executeEvents []*models.ExecuteEvent
		err           error
	)
	if len(event.ExecutorList) == 0 {
		return nil
	}
	// 进入执行消息队列 等待执行
	for _, processorId := range event.ExecutorList {
		executeEvents = append(executeEvents, &models.ExecuteEvent{
			ResourceId:  event.ResourceId,
			ProcessorId: int(processorId),
		})
	}
	if err = driver.SendExecuteEventMsg(executeEvents); err != nil {
		return err
	}
	return nil
}
