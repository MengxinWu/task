package driver

import (
	"context"
	"time"

	"task/models"
)

func AddResource(ctx context.Context, resourceId int64, dagId int, name string) error {
	var err error
	task := &models.Resource{
		ResourceId: resourceId,
		DagId:      dagId,
		Name:       name,
		Status:     0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if _, err = engine.Insert(task); err != nil {
		return err
	}
	return nil
}
