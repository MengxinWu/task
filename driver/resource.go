package driver

import (
	"time"

	"task/models"

	log "github.com/sirupsen/logrus"
)

// AddResource add resource.
func AddResource(resourceId int64, dagId int, name string) error {
	var err error
	task := &models.Resource{
		ResourceId: resourceId,
		DagId:      dagId,
		Name:       name,
		Status:     models.ResourceStatusNormal,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if _, err = engine.Insert(task); err != nil {
		log.Errorf("AddResource engine error: %v", err)
		return err
	}
	return nil
}
