package handler

import (
	"immersive_project/klp3/features/classes"
)

type ClassesRequest struct {
	Name 			string 					`json:"name" form:"name" validate:"required"`
	UserId 			uint 					`json:"user_id" form:"user_id"`
}


func EntityToRequest(class classes.ClassessEntity)ClassesRequest{
	return ClassesRequest{
		Name: 	class.Name,
		UserId: class.UserID,
	}
}

func RequestToEntity(class ClassesRequest) classes.ClassessEntity{
	return classes.ClassessEntity{
		Name: 			class.Name,
		UserID: 		class.UserId,
	}
}