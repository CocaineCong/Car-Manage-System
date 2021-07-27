package model

import "github.com/jinzhu/gorm"

type Information struct {
	gorm.Model
	UserName string
	UserId string
	IsRead bool
	TypeID int
}
