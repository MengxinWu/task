package driver

import (
	"context"
	"fmt"
	"time"

	"task/models"

	log "github.com/sirupsen/logrus"
)

// GetProcessState get process state.
func GetProcessState(_ context.Context, resourceId int64, processorId int) (*models.ProcessState, error) {
	var (
		ok  bool
		err error
	)
	state := new(models.ProcessState)
	if ok, err = engine.Where("resource_id = ? AND processor_id = ?", resourceId, processorId).Get(state); err != nil {
		log.Errorf("GetProcessState engine error: %v", err)
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("not exist resourceId %d, processorId %d", resourceId, processorId)
	}
	return state, nil
}

// AddProcessState add process state.
func AddProcessState(_ context.Context, resourceId int64, processorId, processState int) error {
	var err error
	state := &models.ProcessState{
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
func UpdateProcessState(_ context.Context, resourceId int64, processorId, processState int) error {
	var err error
	state := new(models.ProcessState)
	state.ResourceId = resourceId
	state.ProcessorId = processorId
	state.ProcessState = processState
	if _, err = engine.Where("resource_id = ? AND processor_id = ?", resourceId, processorId).Cols("first_on_shelf_time").Update(state); err != nil {
		log.Errorf("UpdateProcessState engine error: %v", err)
		return err
	}
	return nil
}
