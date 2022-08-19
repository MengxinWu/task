package driver

import (
	"context"
	"time"

	"task/models"

	log "github.com/sirupsen/logrus"
)

// GetProcessState get process state.
func GetProcessState(ctx context.Context, resourceId int64, processorId int) (*models.ResourceState, error) {
	// todo
	return nil, nil
}

// AddProcessState add process state.
func AddProcessState(_ context.Context, resourceId int64, processorId, processState int) error {
	var err error
	state := &models.ResourceState{
		ResourceId:   resourceId,
		ProcessorId:  processorId,
		ProcessState: processState,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	if _, err = engine.Insert(state); err != nil {
		log.Errorf("AddProcessState engine error: %v", err)
		return err
	}
	return nil
}

// UpdateProcessState update process state.
func UpdateProcessState(ctx context.Context, resourceId int64, processorId, processState int) error {
	// todo
	return nil
}
