package api

import (
	logging "github.com/sirupsen/logrus"
	"CarDemo1/serializer"
	"CarDemo1/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)


//UserLogin 用户登陆接口
func UserLogin(c *gin.Context) {
	session := sessions.Default(c)
	status := 200
	userID := session.Get("userId")
	code := c.Request.Header.Get("AuthCode")
	var loginService service.UserLoginService
	if err := c.ShouldBind(&loginService); err == nil {
		res := loginService.Login(userID  , code, status)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg":"Successfully",
	})
}

func MessageUserInfo(c *gin.Context) {
	var service service.MessageInfoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserInfo(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//CheckToken 用户详情
func CheckToken(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Status: 200,
		Msg:    "ok",
	})
}

func BindEmail(c *gin.Context) {
	var service service.VaildEmailService
	authorization := c.Request.Header.Get("Authorization")
	if err := c.ShouldBind(&service); err == nil {
		res := service.Vaild(authorization)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func UserShow(c *gin.Context) {
	var service service.UserInfoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserInfo()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}

func BindPhone(c *gin.Context) {
	var service service.VaildPhoneService
	//var operationType = c.Request.Header.Get("operation_type")
	authorization := c.Request.Header.Get("Authorization")
	if err := c.ShouldBind(&service); err == nil {
		res := service.Vaild(authorization)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func UserGetCode(c *gin.Context) {
	var service service.GetCodeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SendMsg()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
