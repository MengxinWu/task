package driver

import (
	"context"

	"task/ecode"
	"task/models"

	log "github.com/sirupsen/logrus"
)

// GetProcessor get processor.
func GetProcessor(_ context.Context, processorId int) (*models.Processor, error) {
	var (
		ok  bool
		err error
	)
	processor := new(models.Processor)
	if ok, err = engine.Id(processorId).Where("status != ?", models.ProcessorStatusDelete).Get(processor); err != nil {
		log.Errorf("GetProcessState engine error: %v", err)
		return nil, ecode.EngineError
	}
	if !ok {
		return nil, ecode.ProcessorNotFound
	}
	return processor, nil
}
