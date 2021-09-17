package serializer

import "CarDemo1/model"

type Social struct {
	ID          uint   	`json:"id"`
	UserID 		uint 	`json:"user_id"`
	Title 		string 	`json:"title"`
	Content  	string 	`json:"content"`
	CategoryID 	uint 	`json:"category_id"`
	CategoryName string `json:"category_name"`
	EnglishName string  `json:"english_name"`
	Picture 	string 	`json:"picture"`
	UserName 	string 	`json:"user_name"`
	UserAvatar 	string 	`json:"user_avatar"`
}

//序列化话题
func BuildSocial(item model.Society) Social {
	return Social{
		ID:           item.ID,
		UserID : 	  item.UserID,
		Title: 		  item.Title,
		Content  :item.Content,
		CategoryID :	item.CategoryID,
		CategoryName :	item.CategoryName,
		EnglishName:item.EnglishName,
		Picture :item.Picture,
		UserName :item.UserName,
		UserAvatar :item.UserAvatar,
	}
}

//序列化轮播图列表
func BuildSocials(items []model.Society) (socials []Social) {
	for _, item := range items {
		social := BuildSocial(item)
		socials = append(socials, social)
	}
	return socials
}

