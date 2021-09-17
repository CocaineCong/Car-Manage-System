package model

import (
	"github.com/jinzhu/gorm"
)

// 帖子图片模型 多图片
type SocialImg struct {
	gorm.Model
	SocialID uint
	ImgPath   string
}
