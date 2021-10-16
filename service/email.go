package service

import (
	"CarDemo1/model"
	"CarDemo1/pkg/e"
	logging "github.com/sirupsen/logrus"
	"CarDemo1/pkg/util"
	"CarDemo1/serializer"
)

//VaildEmailService 绑定、解绑邮箱和修改密码的服务
type VaildEmailService struct {
	OperationType int `form:"operation_type" json:"operation_type"`
	Email string `form:"email" json:"email"`
}



//Vaild 绑定邮箱
func (service *VaildEmailService) Vaild(authorization string) serializer.Response {
	var email string
	var openid string
	code := e.Success
	claims , _ := util.ParseToken(authorization)
	openid  = claims.OpenID
	email = service.Email
	if service.OperationType == 1 {
		//1.绑定邮箱
		if err := model.DB.Table("user").Where("open_id=?", openid).Update("email", email).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else if service.OperationType == 2 {
		//2.解绑邮箱
		if err := model.DB.Table("user").Where("open_id=?", openid).Update("email", "" ).Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	//获取该用户信息
	var user model.User
	if err := model.DB.First(&user).Where("open_id = ?",openid).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//返回用户信息
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}
