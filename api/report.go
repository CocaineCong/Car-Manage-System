package api

import (
	"CarDemo1/pkg/logging"
	"CarDemo1/service"
	"github.com/gin-gonic/gin"
)

func ShowReport(c *gin.Context) {
	service := service.ReportInfoShow{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// CreateCategory 创建分类
func CreateReport(c *gin.Context) {
	service := service.CreateReportService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func DeleteReport(c *gin.Context) {
	service := service.DeleteReportService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func UpdateReport(c *gin.Context) {
	service := service.UpdateReportService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

