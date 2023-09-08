package handler

import (
	"immersive_project/klp3/features/classes"
)

type ClassesResponse struct {
	Id   uint   `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	UserId uint `json:"user_id"`	
	User UserResponse `json:"user"`
}

type UserResponse struct {
	Id     uint   `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
}

func EntityToResponse(class classes.ClassessEntity) ClassesResponse {
	return ClassesResponse{
		Id:   class.Id,
		Name: class.Name,
		UserId: class.UserID,
		User: UserResponses(class.User),
	}
}
func UserResponses(user classes.UserEntity)UserResponse{
	return UserResponse{
		Id: user.Id,
		Name: user.Name,
	}
}

func EntityToResponseAll(class []classes.ClassessEntity) []ClassesResponse {
	var classes []ClassesResponse
	for _, value := range class {
		classes = append(classes, EntityToResponse(value))
	}
	return classes
}
