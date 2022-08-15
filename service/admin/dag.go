package service

import (
	"task/driver"
	"task/models"

	"github.com/gin-gonic/gin"
)

// ListDag list dag.
func ListDag(c *gin.Context) {
	var (
		err  error
		dags []*models.Dag
	)
	dags, err = driver.ListDag(c)
	HttpResponse(c, dags, err)
}
