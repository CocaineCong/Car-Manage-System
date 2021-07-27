package service

import (
	"CarDemo1/conf"
	"CarDemo1/model"
	"CarDemo1/pkg/e"
	"CarDemo1/pkg/logging"
	"CarDemo1/pkg/util"
	"CarDemo1/serializer"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type WxSession struct {
	SessionKey string `json:"session_key"`
	ExpireIn   int    `json:"expires_in"`
	OpenID     string `json:"openid"`
	Errcode	   int	  `json:"errcode"`
	Errmsg	   string `json:"errmsg"`
}

type UserInfoService struct {
	Token string `form:"token" json:"token"`
}

type UserLoginService struct {
	UserName string `form:"username" json:"username"`
	Code string `form:"code",json:"code"`
	//OpenID string `form:"openid" json:"openid"`
	Avatar string `form:"avatar" json:"avatar"`
	Phone string `form:"Phone" json:"Phone"`
	Email string `form:"email" json:"email"`
	Validate  string `form:"validate" json:"validate"`
	Seccode   string `form:"seccode" json:"seccode"`
}

type MessageInfoService struct {

}


//用户登陆
func (service *UserLoginService)Login(userID interface{}, wx_code string, status int) serializer.Response {
	var user model.User
	code := e.Success
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", conf.AppID, conf.Secret, wx_code)
	var session WxSession
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return serializer.Response{
			Status : e.ErrorCodeReq,
			Msg:  e.GetMsg(e.ErrorCodeReq),
		}
	}
	response ,_ := client.Do(request)
	body , err := ioutil.ReadAll(response.Body)
	if err := json.Unmarshal(body, &session); err != nil {
		return serializer.Response{
			Status: e.ErrorCodeResp,
			Msg:   e.GetMsg(e.ErrorCodeResp),
		}
	}

	if session.Errcode != 0{
		return serializer.Response{
			Status:e.ErrorCodeOrder ,
			Data:   nil,
			Msg:    session.Errmsg,
		}
	}
	fmt.Println(session.OpenID)
	err = model.DB.Where("open_id=?", session.OpenID).Find(&user).Error
	//如果查询不到，返回相应的错误
	if err != nil {
		user.Avatar = service.Avatar
		user.UserName = service.UserName
		if err := model.DB.Create(&user).Error; err != nil {
			logging.Info(err)
		}
		err = model.DB.Save(&user).Error
		if err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	user.Avatar = service.Avatar
	user.UserName = service.UserName
	err = model.DB.Save(&user).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	token, err := util.GenerateToken(user.ID,	service.UserName, session.OpenID, 0)
	if err != nil {
		logging.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}

func (service *UserInfoService) UserInfo() serializer.Response {
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
	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
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

func  (service *MessageInfoService)UserInfo(id string) serializer.Response {
	var user model.User
	code := e.Success
	err := model.DB.First(&user, id).Error
	//fmt.Println(user)
	//err := model.DB.Model(&user).Association("Relations").Find(&friendsList).Error
	//fmt.Println(friendsList)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status:    code,
			Data:      err,
			Msg:       e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status:    code,
		Data:      serializer.BuildUser(user),
		Msg:       e.GetMsg(code),
	}

}