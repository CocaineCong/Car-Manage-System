package api

import (
	logging "github.com/sirupsen/logrus"
	"CarDemo1/service"
	"github.com/gin-gonic/gin"
)

func UpLoad(c *gin.Context) {
	file , fileHeader ,_ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	var service service.UpLoadFile
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpLoadFile(file,fileSize)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//func UploadToken(c *gin.Context) {
//	service := service.UploadAvatarService{}
//	if err := c.ShouldBind(&service); err == nil {
//		res := service.Post()
//		c.JSON(200, res)
//	} else {
//		c.JSON(200, ErrorResponse(err))
//		logging.Info(err)
//	}
//}
