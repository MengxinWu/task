package driver

import (
	"context"
	"fmt"
	"log"

	"task/models"
)

// ListDag list dag.
func ListDag(ctx context.Context) ([]*models.Dag, error) {
	var err error
	dags := make([]*models.Dag, 0)
	if err = engine.Where("status != ?", models.DagStatusDelete).Find(&dags); err != nil {
		return nil, err
	}
	return dags, nil
}

func GetDag(ctx context.Context, dagId int) (*models.Dag, error) {
	var (
		ok  bool
		err error
	)
	dag := new(models.Dag)
	if ok, err = engine.Id(int64(dagId)).Get(dag); err != nil {
		log.Fatalf("GetDag mysql error: %v", err)
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("not exist dag %d", dagId)
	}
	return dag, nil
}
