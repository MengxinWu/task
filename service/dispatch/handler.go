package dispatch

import (
	"context"
	"task/driver"
	"task/models"
)

type EventHandler interface {
	Prepare(context.Context, *models.DispatchEvent) error
	Compute(context.Context, *models.DispatchEvent) error
	After(context.Context, *models.DispatchEvent) error
}

var EventHandlerMap = make(map[string]EventHandler)

func InitEventHandler() {
	EventHandlerMap["resource_add"] = ResourceAddHandler{}
	EventHandlerMap["processor_done"] = ProcessorDoneHandler{}
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
	if event.Graph, err = models.GenerateGraph(event.Dag.Config); err != nil {
		return err
	}
	for _, node := range event.Graph {
		if node.Parents == nil {
			event.ExecutorList = append(event.ExecutorList, int64(node.ProcessorId))
			if err = driver.AddProcessor(ctx, event.ResourceId, node.ProcessorId); err != nil {
				return err
			}
		}
	}
	return nil
}

func (h ResourceAddHandler) After(ctx context.Context, event *models.DispatchEvent) error {
	// todo 插入执行的kafka
	return nil
}

type ProcessorDoneHandler struct {
}

func (h ProcessorDoneHandler) Prepare(ctx context.Context, event *models.DispatchEvent) error {
	return nil
}

func (h ProcessorDoneHandler) Compute(ctx context.Context, event *models.DispatchEvent) error {
	return nil
}

func (h ProcessorDoneHandler) After(ctx context.Context, event *models.DispatchEvent) error {
	return nil
}
