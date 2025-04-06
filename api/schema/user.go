package schema

import "Go_Practice/domain/entity"

type UserRes struct {
	UserId      string `json:"userId"`
	Email       string `json:"email"`
}

func UserResFromEntity(user entity.User) UserRes {
	return UserRes{
		UserId:      user.UserID,
		Email:       user.Email,
	}
}