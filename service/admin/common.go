package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpResponse(c *gin.Context, result interface{}, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", err))
	} else {
		c.JSON(http.StatusOK, result)
	}
}
