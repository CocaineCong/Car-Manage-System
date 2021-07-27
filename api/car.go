package api

import (
	"CarDemo1/pkg/logging"
	"CarDemo1/service"
	"github.com/gin-gonic/gin"
)

//创建车
func CreateCar(c *gin.Context) {
	file , fileHeader ,_ := c.Request.FormFile("file")
	carNum := c.Request.Header.Get("car_num")
	carName := c.Request.Header.Get("car_name")
	carBossID := c.Request.Header.Get("car_boss_id")
	fileSize := fileHeader.Size
	service := service.CreateCarService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(file,fileSize,carNum,carName,StrToUInt(carBossID))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//展示车
func ShowCar(c *gin.Context) {
	service := service.ShowCarService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Param("user_id"))
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
		res := service.Delete(c.Param("car_num"))
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