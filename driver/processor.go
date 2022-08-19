package driver

import (
	"context"
	"fmt"

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
	if ok, err = engine.Id(processorId).Get(processor); err != nil {
		log.Errorf("GetProcessState engine error: %v", err)
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("not exist processorId %d", processorId)
	}
	return processor, nil
}
