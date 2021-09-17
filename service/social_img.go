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

type CreateImgService struct {
	SocialID uint   `form:"social_id" json:"social_id"`
	ImgPath   string `form:"img_path" json:"img_path"`
}

type ShowImgsService struct {
}



func (service *CreateImgService) Create (file multipart.File ,fileSize int64) serializer.Response {
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
	if err != nil {
		code = e.ErrorUploadFile
		return serializer.Response{
			Status: code,
			Data:   err.Error(),
			Msg:    e.GetMsg(code),
		}
	}
	url := ImgUrl + ret.Key
	img := model.SocialImg{
		SocialID: service.SocialID,
		ImgPath:  url,
	}
	err = model.DB.Create(&img).Error
	if err != nil {
		logging.Info(err)
		code := e.ErrorDatabase
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
// ShowImgsService 商品图片详情的服务


// Show 商品图片
func (service *ShowImgsService) Show(id string) serializer.Response {
	var imgs []model.SocialImg
	code := e.Success

	err := model.DB.Where("social_id=?", id).Find(&imgs).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response {
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildImgs(imgs),
	}
}



