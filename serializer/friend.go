package serializer

import "CarDemo1/model"

type Friend struct {
	ID          uint   	`json:"id"`
	UserName 	string 	`json:"user_name"`
	UserAvatar 	string 	`json:"user_avatar"`
}

//序列化话题
func BuildFriend(item model.User) Friend {
	return Friend{
		ID:          item.ID,
		UserName :   item.UserName,
		UserAvatar : item.Avatar,
	}
}

//序列化轮播图列表
func BuildFriends(items []model.User) (Friends []Friend) {
	for _, item := range items {
		Friend := BuildFriend(item)
		Friends = append(Friends, Friend)
	}
	return Friends
}
