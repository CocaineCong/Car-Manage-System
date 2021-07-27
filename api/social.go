package api

import (
	"CarDemo1/pkg/logging"
	"CarDemo1/service"
	"fmt"
	"github.com/gin-gonic/gin"
)


//获取所有的帖子
func GetSocial(c *gin.Context) {
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
func ShowSocial(c *gin.Context) {
	service := service.ShowSocialService{}
	fmt.Println(c.Param("id"))
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

func GetMySocial(c *gin.Context) {
	service := service.ShowMySocialService{}
	res:=service.Show(c.Param("id"))
	c.JSON(200,res)
}


//创造帖子
func CreateSocial(c *gin.Context) {
	file , fileHeader ,_ := c.Request.FormFile("file")
	service := service.CreateSocialService{}
	userId := c.Request.Header.Get("user_id")
	title := c.Param("content")
	content := c.Param("content")
	categoryId := c.Request.Header.Get("category_id")
	fileSize := fileHeader.Size
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(file,fileSize,title,content,StrToUInt(userId),StrToUInt(categoryId))
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
