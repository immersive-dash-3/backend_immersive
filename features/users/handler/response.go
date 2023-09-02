package handler

import "immersive_project/klp3/features/users"

type UserResponse struct {
	Id       uint   `json:"id" form:"id"`
	FullName string `json:"full_name" form:"full_name"`
	Email    string `json:"email" form:"email"`
	Team     string `json:"team" form:"team"`
	Role     string `json:"role" form:"role"`
	Status   string `json:"status" form:"status"`
}

func UserEntityToResponse(user users.UserEntity) UserResponse{
	return UserResponse{
		Id:       user.Id,
		FullName: user.FullName,
		Email:    user.Email,
		Team:     user.Team,
		Role:     user.Role,
		Status:   user.Status,
	}
}

func UserEntityToResponseAll(user []users.UserEntity)[]UserResponse{
	var userAll []UserResponse
	for _,value:=range user{
		userAll=append(userAll, UserEntityToResponse(value))
	}
	return userAll
}