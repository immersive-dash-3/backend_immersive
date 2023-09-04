package handler

import (
	"immersive_project/klp3/features/classes"
)


type ClassesResponse struct {
	Id   			uint    			`json:"id" form:"id"`
	Name 			string  			`json:"name" form:"name"`
}

func ResponseToEntity(class classes.ClassessEntity) ClassesResponse{
	return ClassesResponse{
		Id: 			class.Id,
		Name: 			class.Name,

	}
}

func ResponseToEntityAll(class []classes.ClassessEntity)[]ClassesResponse{
	var classes []ClassesResponse
	for _,value:=range class{
		classes=append(classes, ResponseToEntity(value))
	}
	return classes
}