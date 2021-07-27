package service

import (
	"CarDemo1/model"
	"CarDemo1/pkg/e"
	"CarDemo1/pkg/logging"
	"CarDemo1/pkg/util"
	"CarDemo1/serializer"
	"time"
)

//VaildEmailService 绑定、解绑邮箱和修改密码的服务
type VaildEmailService struct {
	Token string `form:"token" json:"token"`
	Email string `form:"email" json:"email"`
}



//Vaild 绑定邮箱
func (service *VaildEmailService) Vaild(operationType uint) serializer.Response {
	var email string
	var openid string
	code := e.Success
	//验证token
	if service.Token == "" {
		code = e.InvalidParams
	} else {
		claims, err := util.ParseEmailToken(service.Token)
		if err != nil {
			logging.Info(err)
			code = e.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		} else {
			openid = claims.OpenID
		}
	}
	email = service.Email
	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//fmt.Println("opentionType",operationType)
	if operationType == 1 {
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
	} else if operationType == 2 {
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
	//fmt.Println(user)
	//返回用户信息
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
		Msg:    e.GetMsg(code),
	}
}
