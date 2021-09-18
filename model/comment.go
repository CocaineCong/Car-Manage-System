package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	Content    string
	ParentId   uint   	// 父评论ID
	UserID     uint   	// 用户ID
	ReplyName  string 	// 回复者名字
	UserName   string
	UserAvatar string
	SocialId   uint  `gorm:"ForeignKey:Social;AssociationForeignKey:social_id"` //社区帖子ID
	Children   []Comment `gorm:"ForeignKey:ParentId"`
}
