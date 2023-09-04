package handler

import (
	"immersive_project/klp3/features/classes"
)

type ClassesRequest struct {
	Name 			string 					`json:"name" form:"name" validate:"required"`
}

func EntityToRequest(class classes.ClassessEntity)ClassesRequest{
	return ClassesRequest{
		Name: class.Name,
	}
}

func RequestToEntity(class ClassesRequest) classes.ClassessEntity{
	return classes.ClassessEntity{
		Name: 			class.Name,
	}
}