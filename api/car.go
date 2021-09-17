package api

import (
	"CarDemo1/pkg/logging"
	"CarDemo1/pkg/util"
	"CarDemo1/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

//创建车
func CreateCar(c *gin.Context) {
	file , fileHeader ,_ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	service := service.CreateCarService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(file,fileSize)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//展示车
func ShowCar(c *gin.Context) {
	service := service.ShowCarService{}
	chain ,_ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(chain.UserID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//解绑车
func DeleteCar(c *gin.Context) {
	service := service.DeleteCarService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//展示车
func ListCars(c *gin.Context) {
	service := service.ListCarsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//搜索车主信息
func SearchCar(c *gin.Context) {
	service := service.SearchCarBossService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Search()
		fmt.Println("res",res)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}