package api

import (
	"CarDemo1/pkg/logging"
	"CarDemo1/pkg/util"
	"CarDemo1/service"
	"github.com/gin-gonic/gin"
)

//关注好友
func CreateFriend(c *gin.Context) {
	service := service.CreateFriendService{}
	chaim ,_ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(StrToUInt(c.Param("id")),chaim.UserID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//展示好友
func ShowMyFriend(c *gin.Context) {
	service := service.ShowAllFriendService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//解绑好友
func DeleteFriend(c *gin.Context) {
	service := service.DeleteFriendService{}
	token := c.GetHeader("Authorization")
	chmain,_ := util.ParseToken(token)
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete(StrToUInt(c.Param("id")),chmain.UserID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//查看好友信息
func ShowMyFriendInfo(c *gin.Context) {
	service := service.ShowMyFriendInfoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}