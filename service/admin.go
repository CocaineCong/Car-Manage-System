package service

//func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
import (
	"CarDemo1/model"
	"CarDemo1/pkg/e"
	logging "github.com/sirupsen/logrus"
	"CarDemo1/pkg/util"
	"CarDemo1/serializer"
	"github.com/jinzhu/gorm"
)

//管理员注册服务
type AdminRegisterService struct {
	UserName         string `form:"user_name" json:"user_name"`
	Password         string `form:"password" json:"password"`
	PasswordCongfirm string `form:"password_confirm" json:"password_congfirm"`
}

//登陆服务
type AdminLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password"`
}

func (service *AdminRegisterService) Valid() *serializer.Response {
	var code int
	if service.PasswordCongfirm != service.Password {
		code = e.ErrorNotComparePassword
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	count := 0
	model.DB.Model(&model.Admin{}).Where("user_name=？", service.UserName).Count(&count)
	if count > 0 {
		code = e.ErrorExistUser
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return nil
}

//注册
func (service *AdminRegisterService) Register() *serializer.Response {
	admin := model.Admin{
		UserName: service.UserName,
	}
	code := e.Success
	//表单验证
	if res := service.Valid(); res != nil {
		return res
	}
	if err := admin.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ErrorFailEncryption
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建用户
	if err := model.DB.Create(&admin).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

//管理员登陆函数
func (service *AdminLoginService) Login() serializer.Response {
	var admin model.Admin
	code := e.Success
	if err := model.DB.Where("user_name=?", service.UserName).First(&admin).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = e.ErrorNotExistUser
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		code = e.ErrorNotExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if admin.CheckPassword(service.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	token, err := util.GenerateToken(admin.ID,service.UserName, service.Password, 1)
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
		Data:   serializer.TokenData{User: serializer.BuildAdmin(admin), Token: token},
		Msg:    e.GetMsg(code),
	}
}
