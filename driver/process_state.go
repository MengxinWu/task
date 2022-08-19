package driver

import (
	"context"
	"time"

	"task/ecode"
	"task/models"

	log "github.com/sirupsen/logrus"
)

// GetResourceProcessState get resource process state.
func GetResourceProcessState(_ context.Context, resourceId int64, processorId int) (*models.ResourceProcessState, error) {
	var (
		ok  bool
		err error
	)
	state := new(models.ResourceProcessState)
	if ok, err = engine.Where("resource_id = ? AND processor_id = ?", resourceId, processorId).Get(state); err != nil {
		log.Errorf("GetResourceProcessState engine error: %v", err)
		return nil, ecode.EngineError
	}
	if !ok {
		return nil, ecode.ProcessStateNotFound
	}
	return state, nil
}

// AddResourceProcessState add resource process state.
func AddResourceProcessState(_ context.Context, resourceId int64, processorId, processState, processCnt int, processMsg string) error {
	var err error
	state := &models.ResourceProcessState{
		ResourceId:   resourceId,
		ProcessorId:  processorId,
		ProcessCnt:   processCnt,
		ProcessState: processState,
		ProcessMsg:   processMsg,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	if _, err = engine.Insert(state); err != nil {
		log.Errorf("AddResourceProcessState engine error: %v", err)
		return ecode.EngineError
	}
	return nil
}

// UpdateResourceProcessState update resource process state.
func UpdateResourceProcessState(_ context.Context, id, processState, processCnt int, processMsg string) error {
	var err error
	state := new(models.ResourceProcessState)
	state.ProcessCnt = processCnt
	state.ProcessState = processState
	state.ProcessMsg = processMsg
	if _, err = engine.Id(id).Cols("process_cnt", "process_state", "process_msg").Update(state); err != nil {
		log.Errorf("UpdateResourceProcessState engine error: %v", err)
		return ecode.EngineError
	}
	return nil
}
