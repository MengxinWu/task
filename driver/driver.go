package driver

import (
	"context"

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
