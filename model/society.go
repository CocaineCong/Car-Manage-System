package model

import "github.com/jinzhu/gorm"

type Society struct {
	gorm.Model
	CategoryID  uint
	CategoryName string
	EnglishName string
	Title 	string
	Content string
	Picture string
	UserID  	uint
	UserName 	string
	UserAvatar  string
	Price 	string
	Status 	string  // 0 表示未交易  1 表示已交易  2 表示交易完成
}
