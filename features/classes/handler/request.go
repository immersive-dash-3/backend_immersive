package handler

import (
	"immersive_project/klp3/features/classes"
)

type ClassesRequest struct {
	Name 			string 					`json:"name" form:"name" validate:"required"`
}

type PageRequest struct {
	Page 			int 					`json:"page" form:"page"`
	ItemsPerPage    int						`json:"item_per_page" form:"item_per_page"`
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