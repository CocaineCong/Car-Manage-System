package service

import (
	"CarDemo1/model"
	"CarDemo1/pkg/e"
	logging "github.com/sirupsen/logrus"
	"CarDemo1/serializer"
	"fmt"
)

type ReportInfoShow struct {

}

type CreateReportService struct {
	TypeName string   `form:"type_name" json:"type_name"`
	UserName string   `form:"user_name" json:"user_name"`
	Content  string   `form:"content" json:"content"`
	Picture  string   `form:"picture" json:"picture"`
	Finish   uint     `form:"finish" json:"finish"`
}

type DeleteReportService struct {
	ID    uint `form:"id" json:"id"`
}

type UpdateReportService struct {
	ID       uint     `form:"id" json:"id"`
	TypeID   uint     `form:"type_id" json:"type_id"`
	TypeName string   `form:"type_name" json:"type_name"`
	Content  string   `form:"content" json:"content"`
	Picture  string   `form:"picture" json:"picture"`
	Finish   uint     `form:"finish" json:"finish"`
}

//创建反馈
func (service *CreateReportService) Create(id uint) serializer.Response {
	var user model.User
	model.DB.Where(model.User{}).Where("id=?",id).First(&user)
	Report := model.Report{
		TypeName :service.TypeName,
		UserID :user.ID,
		UserName :user.UserName,
		Content :service.Content,
		Picture :service.Picture,
		Finish : 0,
	}
	fmt.Println(Report)
	code := e.Success
	err := model.DB.Create(&Report).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildReport(Report),
		Msg:    e.GetMsg(code),
	}
}

// Show反馈
func (service *ReportInfoShow) List(id uint) serializer.Response {
	var Reports []model.Report
	code := e.Success
	if err := model.DB.Model(model.Report{}).Where("user_id = ?",id).Find(&Reports).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildReports(Reports),
		Msg:    e.GetMsg(code),
	}
}

//获取我的反馈
func (service *ReportInfoShow) Show(id string) serializer.Response {
	var Reports []model.Report
	code := e.Success
	if err := model.DB.Find(&Reports,id).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildReports(Reports),
		Msg:    e.GetMsg(code),
	}
}

//删除反馈
func (service *DeleteReportService) Delete(id string) serializer.Response {
	var Report model.Report
	code := e.Success
	err := model.DB.Where("id=?", id).Find(&Report).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&Report).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

//更改反馈

func (service *UpdateReportService) Update(id string) serializer.Response {
	var Report model.Report
	code := e.Success
	Report = model.Report{
		Finish: 1,
	}
	err := model.DB.Save(&Report).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildReport(Report),
		Msg:    e.GetMsg(code),
	}
}