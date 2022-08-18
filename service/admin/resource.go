package service

import (
	"task/driver"
	"task/models"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// AddResource add resource.
func AddResource(c *gin.Context) (interface{}, error) {
	var (
		resourceId int64
		err        error
	)
	// 解析参数
	request := new(models.AddResourceRequest)
	if err = c.BindJSON(request); err != nil {
		log.Errorf("AddResource bind json request error(%v)", err)
		return nil, err
	}
	// resource 入库
	resourceId = driver.GetSnowFlakeId()
	if err = driver.AddResource(resourceId, request.DagId, request.Name); err != nil {
		return nil, err
	}
	// 进入调度消息队列 等待调度
	if err = driver.SendDispatchEventMsg(&models.DispatchEvent{
		Event:      models.DispatchEventResourceAdd,
		ResourceId: resourceId,
		DagId:      request.DagId,
	}); err != nil {
		return nil, err
	}
	return resourceId, nil
}
