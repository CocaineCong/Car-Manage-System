package serializer

import (
	"CarDemo1/model"
)

type Comment struct {
	ID       		uint   `json:"id"`
	UserName 		string `json:"user_name"`
	ReplyName 		string `json:"reply_name"`
	UserAvatar   	string `json:"user_avatar"`
	ParentId   		uint   `json:"parent_id"`
	CreateAt 		int64  `json:"create_at"`
	UserId			uint   `json:"UserId"`
	Content			string `json:"content"`
	Children		[]Comment	`json:"children"`
	SocialID 	    uint   `json:"social_id"`   //社区帖子ID
	Total			int		`json:"total"`
}

//BuildComment 序列化评论
func BuildComment(comment model.Comment) Comment {
	children := make([]Comment, 0)
	n := 0
	model.DB.Model(&comment).Preload("Children").Find(&comment)
	for _, child := range comment.Children{
		n++
		if n > 3 {
			break
		}
		children = append(children, Comment{
			ID:        child.ID,
			UserName:  child.UserName,
			ReplyName: child.ReplyName,
			UserAvatar:    child.UserAvatar,
			ParentId:  child.ParentId,
			CreateAt:  child.CreatedAt.Unix(),
			UserId:    child.UserID,
			Content:   child.Content,
			Children:  nil,
		})
	}
	var total int
	total = len(comment.Children)
	return Comment{
		ID:       comment.ID,
		UserName: comment.UserName,
		ReplyName:comment.ReplyName,
		UserAvatar:	  comment.UserAvatar,
		ParentId: comment.ParentId,
		CreateAt: comment.CreatedAt.Unix(),
		UserId:   comment.UserID,
		Content:  comment.Content,
		Children: children,
		Total:	  total,
	}
}

func BuildComments(items []model.Comment) (comments []Comment) {
	for _, item := range items {
		comment := BuildComment(item)
		comments = append(comments, comment)
	}
	return comments
}