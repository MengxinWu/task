package service

import (
	"fmt"
	"task/models"

	"task/driver"

	"github.com/gin-gonic/gin"
)

type AddResourceReq struct {
	DagId int    `json:"dag_id"`
	Name  string `json:"name"`
}

// AddResource add resource.
func AddResource(c *gin.Context) {
	var err error

	req := new(AddResourceReq)
	if err = c.BindJSON(req); err != nil {
		fmt.Println(err)
	}
	resourceId := driver.GetSnowFlakeId()
	err = driver.AddResource(c, resourceId, req.DagId, req.Name)

	// 放入消息队列 kafka
	dispatchEvent := &models.DispatchEvent{
		Event:      models.EventResourceAdd,
		ResourceId: resourceId,
		DagId:      req.DagId,
	}
	err = driver.InsertDispatchEvent(dispatchEvent)
	HttpResponse(c, resourceId, err)
}
