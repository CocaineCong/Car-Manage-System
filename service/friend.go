package service

import (
	"CarDemo1/model"
	"CarDemo1/pkg/e"
	"CarDemo1/serializer"
	"fmt"
)

type ShowFriendService struct {

}

type ShowMyFriendInfoService struct {

}

type CreateFriendService struct {
	UserID uint  `form:"user_id" json:"user_id"`
}

type DeleteFriendService struct {
	UserID  string  `form:"user_id" json:"user_id"`
}


//展示好友
func (service *ShowFriendService) Show(id string) serializer.Response {
	var user model.User
	var friendsList []model.User
	code := e.Success
	model.DB.Table("user").Where("id = ?",id).First(&user)
	fmt.Println(user)
	err := model.DB.Model(&user).Association("Relations").Find(&friendsList).Error
	fmt.Println(friendsList)
	if err != nil {
		code = e.ErrorFriendFound
		return serializer.Response{
			Status:    code,
			Data:      err,
			Msg:       e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status:    code,
		Data:      serializer.BuildFriends(friendsList),
		Msg:       e.GetMsg(code),
	}
}

//关注好友
func (service *CreateFriendService) Create(id,userId uint) serializer.Response {
	var user model.User
	var friend model.User
	code := e.Success
	model.DB.Model(&friend).Where(`id = ? `,id).First(&friend)  			//被关注者
	model.DB.Model(&user).Where(`id = ?`,userId).First(&user)	//关注者
	err := model.DB.Model(&user).Association(`Relations`).Append([]model.User{friend}).Error
	if err != nil {
		code = e.ErrorFriendFound
		return serializer.Response{
			Status:    code,
			Msg:       e.GetMsg(code),
			Error:		err.Error(),
		}
	}
	return serializer.Response{
		Status:    code,
		Msg:       e.GetMsg(code),
	}
}

//解除好友关系
func (service *DeleteFriendService) Delete(id , userId uint) serializer.Response {
	var user model.User
	var friend []model.User
	code := e.Success
	model.DB.Model(&friend).Where(`id = ?`,id).First(&friend)  			//被关注者
	model.DB.Model(&user).Where(`id = ?`,userId).First(&user)	//关注者
	err := model.DB.Model(&user).Association(`Relations`).Delete(friend).Error
	//model.DB.Model(&user).Association("Relations").Clear()
	// Remove the relationship between source & arguments if exists
	// only delete the reference, won’t delete those objects from DB.
	if err != nil {
		code = e.ErrorFriendFound
		return serializer.Response{
			Status:    code,
			Msg:       e.GetMsg(code),
			Error:err.Error(),
		}
	}
	return serializer.Response{
		Status:   code,
		Msg:      e.GetMsg(code),
	}
}

//展示好友的详细信息
func (service *ShowMyFriendInfoService) Show(id string) serializer.Response {
	var user model.User
	code := e.Success
	err := model.DB.Model(&model.User{}).Where(`id = ?`,id).First(&user).Error
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
