package api

import (
	"CarDemo1/pkg/logging"
	"CarDemo1/pkg/util"
	"CarDemo1/service"
	"github.com/gin-gonic/gin"
)


//获取所有的帖子
func GetAllSocial(c *gin.Context) {
	service := service.SocialInfoShow{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//详细的帖子
func SocialDetail(c *gin.Context) {
	service := service.ShowSocialService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

func GetMySocial(c *gin.Context) {
	service := service.ShowMySocialService{}
	chain ,_ := util.ParseToken(c.GetHeader("Authorization"))
	res:=service.Show(chain.UserID)
	c.JSON(200,res)
}

//创造帖子
func CreateSocial(c *gin.Context) {
	file , fileHeader ,_ := c.Request.FormFile("file")
	service := service.CreateSocialService{}
	chaim ,_ := util.ParseToken(c.GetHeader("Authorization"))
	fileSize := fileHeader.Size
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(file,fileSize,chaim.UserID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//删除帖子
func DeleteSocial(c *gin.Context) {
	service := service.DeleteSocialService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}

//更新帖子
func UpdateSocial(c *gin.Context) {
	service := service.UpdateSocialService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//搜索帖子
func SearchSocial(c *gin.Context) {
	service := service.SearchSocialsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
