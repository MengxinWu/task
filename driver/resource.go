package driver

import (
	"context"
	"fmt"
	"log"
	"time"

	"task/models"
)

func AddResource(ctx context.Context, resourceId int64, dagId int, name string) error {
	var err error
	resource := &models.Resource{
		ResourceId: resourceId,
		DagId:      dagId,
		Name:       name,
		Status:     0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if _, err = engine.Insert(resource); err != nil {
		return err
	}
	return nil
}

func GetResource(ctx context.Context, resourceId int64) (*models.Resource, error) {
	var (
		ok  bool
		err error
	)
	resource := new(models.Resource)
	if ok, err = engine.Id(resourceId).Get(resource); err != nil {
		log.Fatalf("GetResource mysql error: %v", err)
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("not exist resource %d", resourceId)
	}
	return resource, nil
}
