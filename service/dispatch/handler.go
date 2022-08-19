package service

import (
	"context"

	"task/driver"
	"task/ecode"
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
	if event.Dag, err = driver.GetDag(ctx, event.DagId); err != nil {
		return err
	}
	if event.ResourceProcessState, err = driver.GetResourceProcessState(ctx, event.ResourceId, event.ProcessorId); err != nil {
		return err
	}
	if event.ResourceProcessState.ProcessCnt >= models.MaxRetryCnt {
		log.Warnf("process count(%d) over limit", event.ResourceProcessState.ProcessCnt)
		return nil
	}
	return nil
}

func (h ProcessorDoneHandler) Compute(ctx context.Context, event *models.DispatchEvent) error {
	var (
		processState *models.ResourceProcessState
		children     []*models.Node
		err          error
	)
	// 失败任务重试
	if event.ResourceProcessState.ProcessState == models.ProcessStateFail {
		log.Printf("process execute fail %d, %d, retrying...", event.ResourceId, event.ProcessorId)
		event.ExecutorList = append(event.ExecutorList, int64(event.ProcessorId))
		// 设置处理状态为等待执行
		if err = driver.UpdateResourceProcessState(ctx, event.ResourceProcessState.Id, models.ProcessStateReady,
			event.ResourceProcessState.ProcessCnt, ""); err != nil {
			return err
		}
		return nil
	}
	// 寻找下游节点的上游节点
	// 生成DAG图
	if event.Graph, err = models.GenerateGraph(event.Dag.Config); err != nil {
		return err
	}
	if _, ok := event.Graph[event.ProcessorId]; !ok {
		log.Errorf("process processor id: %d not in dag", event.ProcessorId)
		return ecode.ProcessorNotInDag
	}
	if children = event.Graph[event.ProcessorId].Children; children == nil {
		return nil
	}
	for _, childrenNode := range children {
		ready := true
		for _, node := range childrenNode.Parents {
			if node.ProcessorId == event.ProcessorId {
				continue
			}
			if processState, err = driver.GetResourceProcessState(ctx, event.ResourceId, node.ProcessorId); err != nil {
				return err
			}
			if processState.ProcessState != models.ProcessStateSuccess {
				ready = false
				break
			}
		}
		if ready {
			if _, err = driver.GetResourceProcessState(ctx, event.ResourceId, childrenNode.ProcessorId); err != nil {
				if err == ecode.ProcessStateNotFound {
					err = nil
					// 设置处理状态为等待执行
					event.ExecutorList = append(event.ExecutorList, int64(childrenNode.ProcessorId))
					if err = driver.AddResourceProcessState(ctx, event.ResourceId, childrenNode.ProcessorId, models.ProcessStateReady, 0, ""); err != nil {
						return err
					}
				}
				return err
			}
		}
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
