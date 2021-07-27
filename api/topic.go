package api

import (
	"CarDemo1/pkg/logging"
	"CarDemo1/service"
	"github.com/gin-gonic/gin"
)


func GetTopic(c *gin.Context) {
	service := service.TopicInfoShow{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// CreateCategory 创建分类
func CreateCategory(c *gin.Context) {
	service := service.CreateCategoryService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
