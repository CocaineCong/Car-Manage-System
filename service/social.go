package service

import (
	"CarDemo1/model"
	"CarDemo1/pkg/e"
	logging "github.com/sirupsen/logrus"
	"CarDemo1/pkg/util"
	"CarDemo1/serializer"
	"mime/multipart"
)

//全部帖子
type SocialInfoShow struct {
	PageSize int `form:"page_size" json:"page_size"`
	PageNum int `form:"page_num" json:"page_num"`
	Type  int `form:"type" json:"type"`
}
//某个帖子细节
type ShowSocialService struct {
}
//我的帖子的细节
type ShowMySocialService struct {
}
//删除商品的服务
type DeleteSocialService struct {
}
//更新帖子的服务
type UpdateSocialService struct {
	ID            uint   `form:"id" json:"id"`
	CategoryID    uint    `form:"category_id" json:"category_id"`
	CategoryName   string    `form:"category_name" json:"category_name"`
	Title         string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content       string `form:"content" json:"info" content:"max=1000"`
	Picture       string `form:"picture" json:"picture"`
}
//创建帖子的服务
type CreateSocialService struct {
	ID            uint   `form:"id" json:"id"`
	UserID        uint `form:"user_id" json:"user_id"`
	UserName 	  string  `form:"user_name" json:"user_name"`
	UserAvatar    string  `form:"user_avatar" json:"user_avatar"`
	CategoryID    uint    `form:"category_id" json:"category_id"`
	CategoryName  string    `form:"category_name" json:"category_name"`
	Title         string `form:"title" json:"title" `
	Content       string `form:"content" json:"content" binding:"max=1000"`
	Picture       string `form:"picture" json:"picture"`
}
//搜索帖子的服务
type SearchSocialsService struct {
	Search string `form:"search" json:"search"`
	PageSize int `form:"page_size" json:"page_size"`
	PageNum int `form:"page_num" json:"page_num"`
}


func (service *SocialInfoShow) List() serializer.Response {
	var Society []model.Society
	total := 0
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	if err := model.DB.Model(model.Society{}).Count(&total).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if err := model.DB.Limit(service.PageSize).Offset((service.PageNum-1)*service.PageSize).Find(&Society).
		Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildSocials(Society), uint(total))
}

// 帖子详情
func (service *ShowSocialService) Show(id string) serializer.Response {
	var social model.Society
	code := e.Success
	err := model.DB.First(&social, id).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	social.AddView()  //增加点击
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildSocial(social),
		Msg:    e.GetMsg(code),
	}
}

//创造帖子 //TODO 加上多张图片的
func (service *CreateSocialService) Create(file multipart.File ,fileSize int64, userId uint) serializer.Response {
	code:=e.Success
	var social model.Society
	var user model.User
	var topic model.Category
	status , info := util.UploadToQiNiu(file,fileSize)
	if status != 200 {
		return serializer.Response{
			Status:  status  ,
			Data:      e.GetMsg(status),
			Error:info,
		}
	}
	err := model.DB.First(&user, userId).Error //找用户
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Data:   err.Error(),
			Msg:    e.GetMsg(code),
		}
	}
	err = model.DB.First(&topic, service.CategoryID).Error  //找分类
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Data:   err.Error(),
			Msg:    e.GetMsg(code),
		}
	}
	social = model.Society{
		CategoryID :   service.CategoryID,
		CategoryName : topic.CategoryName,
		EnglishName:   topic.EnglishName,
		Title :        service.Title,
		Content :      service.Content,
		Picture :      info,
		UserID :       userId,
		UserName :     user.UserName,
		UserAvatar :   user.Avatar,
	}
	err = model.DB.Create(&social).Error
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
		Data:   serializer.BuildSocial(social),
		Msg:    e.GetMsg(code),
	}
}

//删除帖子
func (service *DeleteSocialService) Delete(id string) serializer.Response {
	var social model.Society
	code := e.Success
	err := model.DB.First(&social, id).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&social).Error
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

//更新帖子
func (service *UpdateSocialService) Update() serializer.Response {
	social := model.Society{
		CategoryID: service.CategoryID,
		Title:      service.Title,
		Content:    service.Content,
		//Picture:       service.Picture,
		CategoryName: service.CategoryName,
	}
	code := e.Success
	err := model.DB.Save(&social).Error
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

//搜索帖子
func (service *SearchSocialsService) Show() serializer.Response {
	var socials []model.Society
	code := e.Success
	err := model.DB.Where("title LIKE ? OR content LIKE ?","%"+service.Search+"%", "%"+service.Search+"%").
		Limit(service.PageSize).Offset((service.PageNum-1)*service.PageSize).Find(&socials).Error
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
		Data:   serializer.BuildSocials(socials),
		Msg:    e.GetMsg(code),
	}
}

func (service *ShowMySocialService) Show(id uint) serializer.Response{
	var Socials []model.Society
	code := e.Success
	if err := model.DB.Where("user_id = ?",id).Find(&Socials).Error; err != nil {
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
		Data:   serializer.BuildSocials(Socials),
		Msg:    e.GetMsg(code),
	}
}