package dto

import "vanwhebin/try-gin-vue/model"

type UserDto struct{
	Name string `json:"name"`
	Telephone string `json:"telephone"`
}


func ToUseDto(user model.User) UserDto {
	return UserDto{
		Name: user.Name,
		Telephone: user.Telephone,
	}

}