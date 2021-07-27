package service

import (
	"CarDemo1/model"
	"CarDemo1/pkg/e"
	"CarDemo1/pkg/logging"
	"CarDemo1/serializer"
)

//ListCarouselsService 视频列表服务
type ListCarouselsService struct {
}

type CreateCarouselService struct {
	ImgPath string `form:"img_path" json:"img_path"`
}

//创建轮播图
func (service *CreateCarouselService) Create() serializer.Response {
	carousel := model.Carousel{
		ImgPath: service.ImgPath,
	}
	code := e.Success
	err := model.DB.Create(&carousel).Error
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
		Data:   serializer.BuildCarousel(carousel),
		Msg:    e.GetMsg(code),
	}
}

//视频列表
func (service *ListCarouselsService) List() serializer.Response {
	var carousels []model.Carousel
	code := e.Success
	if err := model.DB.Find(&carousels).Error; err != nil {
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
		Data:   serializer.BuildCarousels(carousels),
		Msg:    e.GetMsg(code),
	}
}
