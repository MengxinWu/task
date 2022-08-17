package service

import (
	"fmt"
	"math/rand"

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
	// todo
	resourceId := int64(111111111) + int64(rand.Int())
	fmt.Println(resourceId, req.DagId, req.Name)

	err = driver.AddResource(c, resourceId, req.DagId, req.Name)

	// todo kafka

	HttpResponse(c, resourceId, err)
}
