package service

import (
	"CarDemo1/model"
	"CarDemo1/pkg/e"
	logging "github.com/sirupsen/logrus"
	"CarDemo1/serializer"
)

type TopicInfoShow struct {

}

type CreateCategoryService struct {
	CategoryID   uint   `form:"category_id" json:"category_id"`
	CategoryName string `form:"category_name" json:"category_name"`
	EnglishName string `form:"english_name" json:"english_name"`
}

// 分类
func (service *TopicInfoShow) List() serializer.Response {
	var Topics []model.Category
	code := e.Success
	if err := model.DB.Find(&Topics).Error; err != nil {
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
		Data:   serializer.BuildTopics(Topics),
		Msg:    e.GetMsg(code),
	}
}

//创建分类
func (service *CreateCategoryService) Create() serializer.Response {
	category := model.Category{
		CategoryName: service.CategoryName,
		EnglishName:service.EnglishName,
	}
	code := e.Success
	err := model.DB.Create(&category).Error
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
		Data:   serializer.BuildTopic(category),
		Msg:    e.GetMsg(code),
	}
}