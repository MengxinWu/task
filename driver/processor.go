package driver

import (
	"context"
	"time"

	"task/models"
)

func AddProcessor(ctx context.Context, resourceId int64, processorId int) error {
	var err error
	state := &models.ResourceState{
		ResourceId:   resourceId,
		ProcessorId:  processorId,
		ProcessState: models.ProcessStateReady,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	if _, err = engine.Insert(state); err != nil {
		return err
	}
	return nil
}
