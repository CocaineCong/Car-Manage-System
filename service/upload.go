package service

import (
	"CarDemo1/conf"
	"CarDemo1/pkg/e"
	"CarDemo1/serializer"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

type UpLoadFile struct {

}

func (service *UpLoadFile) UpLoadFile(file multipart.File ,fileSize int64) serializer.Response {
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
			Data:   err,
			Msg:    e.GetMsg(code),
		}
	}
	url := ImgUrl + ret.Key
	return serializer.Response{
		Status: code,
		Data:   url,
		Msg:    e.GetMsg(code),
	}
}