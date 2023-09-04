package handler

import "immersive_project/klp3/features/users"

type UserReguest struct {
	Name    	string `json:"name" form:"name"`
	Email    	string `json:"email" form:"email"`
	Team     	string `json:"team" form:"team"`
	Password 	string `json:"password" form:"password"`
	Role     	string `json:"role" form:"role"`
	Address  	string `json:"address" form:"address"`
	Status   	string `json:"status" form:"status"`
}

func UserRequestToEntity(user UserReguest) users.UserEntity{
	return users.UserEntity{
		Name:  	   user.Name,
		Email:     user.Email,
		Team:      user.Team,
		Password:  user.Password,
		Role:      user.Role,
		Address:   user.Address,
		Status:    user.Status,
	}
}