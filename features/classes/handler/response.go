package handler

import (
	"immersive_project/klp3/features/classes"
)

type ClassesResponse struct {
	Id   uint   `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	UserId uint `json:"user_id"`
}

func EntityToResponse(class classes.ClassessEntity) ClassesResponse {
	return ClassesResponse{
		Id:   class.Id,
		Name: class.Name,
		UserId: class.UserID,
	}
}

func EntityToResponseAll(class []classes.ClassessEntity) []ClassesResponse {
	var classes []ClassesResponse
	for _, value := range class {
		classes = append(classes, EntityToResponse(value))
	}
	return classes
}
