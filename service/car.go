package service

import (
	"CarDemo1/conf"
	"CarDemo1/model"
	"CarDemo1/pkg/e"
	"CarDemo1/pkg/logging"
	"CarDemo1/serializer"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

type ShowCarService struct {
}

type CreateCarService struct {
	CarNum    string `form:"car_num" json:"car_num"`
	CarImages string `form:"car_images" json:"car_images"`
	CarBoss    string `form:"car_boss" json:"car_boss"`
	CarName		string `form:"car_name" json:"car_name"`
	CarBossId	uint 	`form:"car_boss_id" json:"car_boss_id"`
}

type DeleteCarService struct {
	CarNum    uint `form:"car_num" json:"car_num"`
	CarBossId uint `form:"car_boss_id" json:"car_boss_id"`
}

type ListCarsService struct {
	Limit int `form:"limit" json:"limit"`
	Start int `form:"start" json:"start"`
	Type  int `form:"type" json:"type"`
}

//车辆图片
func (service *ShowCarService) Show(id string) serializer.Response {
	var Car []model.Car
	total := 0
	code := e.Success
	if err := model.DB.Model(&Car).Where("car_boss_id=?", id).Count(&total).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err := model.DB.Where("car_boss_id=?", id).Find(&Car).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCars(Car), uint(total))
}

//创建车辆
func (service *CreateCarService) Create(file multipart.File ,fileSize int64,carNum, carName  string, carBossID uint) serializer.Response {
	code := e.Success
	var AccessKey = conf.AccessKey
	var SerectKey = conf.SerectKey
	var Bucket = conf.Bucket
	var ImgUrl = conf.QiniuServer
	putPlicy := storage.PutPolicy{
		Scope:Bucket,
	}
	mac := qbox.NewMac(AccessKey,SerectKey)
	upToken := putPlicy.UploadToken(mac)
	cfg := storage.Config{
		Zone : &storage.ZoneHuanan,
		UseCdnDomains : false,
		UseHTTPS : false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(),&ret,upToken,file,fileSize,&putExtra)
	//fmt.Println(err)
	if err != nil {
		code = e.ErrorUploadFile
		return serializer.Response{
			Status: code,
			Data:   err.Error(),
			Msg:    e.GetMsg(code),
		}
	}
	url := ImgUrl + ret.Key
	var Car model.Car
	Car = model.Car{
		CarName:  carNum,
		CarNum:    carName ,
		CarImages: url,
		CarBossId:  carBossID,
	}
	err = model.DB.Create(&Car).Error
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
		Data:   serializer.BuildCar(Car),
		Msg:    e.GetMsg(code),
	}
}

func (service *DeleteCarService) Delete(car_num string) serializer.Response {
	var Car model.Car
	code := e.Success
	err := model.DB.Where("id=?",car_num).Find(&Car).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&Car).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   e.GetMsg(code),
	}
}

func (service *ListCarsService) List() serializer.Response {
	var Car []model.Car
	total := 0
	code := e.Success
	if service.Limit == 0 {
		service.Limit = 15
	}
	if err := model.DB.Model(model.Car{}).Count(&total).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if err := model.DB.Model(model.Car{}).Limit(service.Limit).Offset(service.Start).Find(&Car).
		Error; err != nil {
			logging.Info(err)
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	return serializer.BuildListResponse(serializer.BuildCars(Car), uint(total))
}

