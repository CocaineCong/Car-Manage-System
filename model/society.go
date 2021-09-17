package model

import "github.com/jinzhu/gorm"

type Society struct {
	gorm.Model
	CategoryID  uint
	CategoryName string
	EnglishName string
	Title 	string
	Content string
	Picture []string
	UserID  uint
	UserName 	string
	UserAvatar  string
	Price string
	Status string
}
