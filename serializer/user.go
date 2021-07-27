package serializer

import "CarDemo1/model"

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

//BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		Phone : user.Phone,
		Email:    user.Email,
		Avatar:   user.AvatarURL(),
		CreateAt: user.CreatedAt.Unix(),
	}
}

func BuildUsers(items []model.User) (users []User) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}
