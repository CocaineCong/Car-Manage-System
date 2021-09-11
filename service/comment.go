package service

import (
	"CarDemo1/model"
	"CarDemo1/pkg/e"
	"CarDemo1/pkg/logging"
	"CarDemo1/serializer"
	"fmt"
)

type Page struct {
	Page    	int 	`form:"page"`
	Size 		int 	`form:"size"`
	Desc		int 	`form:"desc"`
	TopicHash	string	`form:"topicHash"`	// 话题唯一标识
	Keyword		string	`form:"keyword"`
	Id			uint	`form:"id"`			// 评论id
}


type CreateNewComment struct {
	Content     string 	`form:"content" json:"content" xml:"content"`
	SocialID    uint  `form:"social_id" json:"social_id" xml:"social_id"`
	ReplyName   string 	`form:"reply_name" json:"reply_name" xml:"reply_name"`
	ParentId 	uint 	`form:"parent_id" json:"parent_id" xml:"parent_id"`
}

type DeleteCommentService struct {

}

//新增评论
func (service *CreateNewComment) Create(id uint) serializer.Response {
	code := e.Success
	var user model.User
	model.DB.Model(model.User{}).Where("id=?",id).First(&user)
	var comment model.Comment
	comment = model.Comment{
		UserID:    user.ID,
		Content:   service.Content,
		ParentId:  service.ParentId,
		UserName:  user.UserName,
		//ReplyName: service.ReplyName,
		UserAvatar:  user.Avatar,
		SocialId: service.SocialID,
	}
	if err := model.DB.Create(&comment).Error; err!=nil{
		code = e.ErrorDatabase
		return serializer.Response{
			Status:    code,
			Data:      err,
			Msg:       e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status:    code,
		Data:      serializer.BuildComment(comment),
		Msg:       e.GetMsg(code),
	}
}

func (service *DeleteCommentService)Delete(id string) serializer.Response{
	code := e.Success
	var comment model.Comment
	err := model.DB.First(&comment, id).Error
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&comment).Error
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

func (service *Page) Find(id string) serializer.Response {
	code := e.Success
	var comment model.Comment
	if err := model.DB.First(&comment, id).Error; err != nil {
		code = e.ErrorCommentNotFound
		return serializer.Response{
			Status: code,
			Data:   err,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildComment(comment),
		Msg:    e.GetMsg(code),
	}
}

func (service *Page) FindChildren(id string) serializer.Response {
	var comment model.Comment
	code := e.Success
	if err := model.DB.First(&comment, id).Error; err != nil {
		code = e.ErrorCommentError
		return serializer.Response{
			Status: code,
			Data:   err,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data : serializer.BuildComment(comment),
		DataOrder: serializer.BuildComments(comment.Children),
		Msg:   e.GetMsg(code),
	}
}

func (service *Page) List(id string,pageSize int, pageNum int) serializer.Response {
	code := e.Success
	var commentList []model.Comment
	var total int
	model.DB.Find(&commentList).Count(&total)
	err := model.DB.Where("social_id = ?",id).Find(&commentList).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	if err != nil {
		code = e.ErrorCommentError
		return serializer.Response{
			Status:    code,
			Data:      err,
			Msg:       e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status:   code,
		Data:     serializer.BuildComments(commentList),
		DataOrder: serializer.BuildTotalPageNum(total,pageNum),
		Msg:      e.GetMsg(code),
	}
}

func (service *Page) ListChild() serializer.Response {
	code := e.Success
	var comments []model.Comment
	page := service.Page
	pageSize := service.Size
	desc := service.Desc
	id := service.Id
	keyword := service.Keyword
	if page <= 0 {
		page = 1
	}
	var err error
	if desc == 1{	// 按时间反向
		err = model.DB.Where("parent_id = ?", id).Where(fmt.Sprintf(" content like %q ", "%" + keyword + "%")).Limit(pageSize).Offset((page-1)*pageSize).Order("created_at desc").Find(&comments).Error
	} else {
		err = model.DB.Where("parent_id = ?", id).Where(fmt.Sprintf(" content like %q ", "%" + keyword + "%")).Limit(pageSize).Offset((page-1)*pageSize).Find(&comments).Error
	}
	if err != nil {
		code = e.ErrorCommentError
		return serializer.Response{
			Status:    code,
			Data:      err,
			Msg:       e.GetMsg(code),
		}
	}
	var total int
	model.DB.Model(&model.Comment{}).Where("parent_id = ?", id).Where(fmt.Sprintf(" content like %q ", "%" + keyword + "%")).Count(&total)
	var pageNum = total / pageSize
	if total % pageSize != 0{
		pageNum++
	}
	return serializer.Response{
		Status:    code,
		Data:      serializer.BuildComments(comments),
		DataOrder: serializer.BuildTotalPageNum(total,pageNum),
		Msg:       e.GetMsg(code),
	}
}

