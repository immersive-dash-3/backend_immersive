package handler

import "immersive_project/klp3/features/users"

type UserReguest struct {
	FullName    string `json:"full_name" form:"full_name"`
	Email    	string `json:"email" form:"email"`
	Team     	string `json:"team" form:"team"`
	Password 	string `json:"password" form:"password"`
	Role     	string `json:"role" form:"role"`
	Address  	string `json:"address" form:"address"`
	Status   	string `json:"status" form:"status"`
}

func UserRequestToEntity(user UserReguest) users.UserEntity{
	return users.UserEntity{
		FullName:  user.FullName,
		Email:     user.Email,
		Team:      user.Team,
		Password:  user.Password,
		Role:      user.Role,
		Address:   user.Address,
		Status:    user.Status,
	}
}