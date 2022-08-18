package driver

import (
	"context"
	"fmt"

	"task/models"

	log "github.com/sirupsen/logrus"
)

// ListDag list dag.
func ListDag(_ context.Context) ([]*models.Dag, error) {
	var err error
	dags := make([]*models.Dag, 0)
	if err = engine.Where("status != ?", models.DagStatusDelete).Find(&dags); err != nil {
		log.Errorf("ListDag engine error: %v", err)
		return nil, err
	}
	return dags, nil
}

// GetDag get dag.
func GetDag(_ context.Context, dagId int) (*models.Dag, error) {
	var (
		ok  bool
		err error
	)
	dag := new(models.Dag)
	if ok, err = engine.Id(dagId).Get(dag); err != nil {
		log.Errorf("GetDag engine error: %v", err)
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("not exist dag %d", dagId)
	}
	return dag, nil
}
