package dispatch

import (
	"context"
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
}

type ResourceAddHandler struct {
}

func (h ResourceAddHandler) Prepare(ctx context.Context, event *models.DispatchEvent) error {
	// todo resource
	event.Resource, err = driver.GetReource(ctx, event.ResourceId)
	// todo dag
	event.Dag, err = driver.GetDag(ctx, event.DagId)
	return nil
}

func (h ResourceAddHandler) Compute(ctx context.Context, event *models.DispatchEvent) error {
	//
	event.Graph, err = models.GenerateGraph(event.Dag.Config)

	//
	for _, node := range event.Graph {
		if node.Parents == nil {
			event.ExecutorList = append(event.ExecutorList, int64(node.ProcessorId))
		}
	}
	return nil
}

func (h ResourceAddHandler) After(ctx context.Context, event *models.DispatchEvent) error {
	// todo 插入执行的kafka
	return nil
}
