package api

import (
	"CarDemo1/service"
	"github.com/gin-gonic/gin"
)

// 创建帖子图片
//func CreateSocialImg(c *gin.Context) {
//	service := service.CreateImgService{}
//	file , fileHeader ,_ := c.Request.FormFile("file")
//	chaim ,_ := util.ParseToken(c.GetHeader("Authorization"))
//	fileSize := fileHeader.Size
//	if err := c.ShouldBind(&service); err == nil {
//		res := service.Create()
//		c.JSON(200, res)
//	} else {
//		c.JSON(200, ErrorResponse(err))
//		logging.Info(err)
//	}
//}

// 帖子图片的详情地址
func ShowSocialImgs(c *gin.Context) {
	service := service.ShowImgsService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

