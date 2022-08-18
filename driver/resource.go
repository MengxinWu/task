package driver

import (
	"context"
	"fmt"
	"time"

	"task/models"

	log "github.com/sirupsen/logrus"
)

// AddResource add resource.
func AddResource(resourceId int64, dagId int, name string) error {
	var err error
	resource := &models.Resource{
		ResourceId: resourceId,
		DagId:      dagId,
		Name:       name,
		Status:     models.ResourceStatusNormal,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if _, err = engine.Insert(resource); err != nil {
		log.Errorf("AddResource engine error: %v", err)
		return err
	}
	return nil
}

// GetResource get resource.
func GetResource(_ context.Context, resourceId int64) (*models.Resource, error) {
	var (
		ok  bool
		err error
	)
	resource := new(models.Resource)
	if ok, err = engine.Id(resourceId).Get(resource); err != nil {
		log.Errorf("GetResource engine error: %v", err)
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("not exist resource %d", resourceId)
	}
	return resource, nil
}
