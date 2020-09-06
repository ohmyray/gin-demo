package dto

import "github.com/ohmyray/gin-demo01/model"

type UserDto struct {
	Username string `json:"username"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Username: user.Username,
		Telephone: user.Telephone,
	}

}