package handler

import (
	"immersive_project/klp3/features/classes"
)

type ClassesRequest struct {
	Name 			string 					`json:"name" form:"name"`
	UserID 			uint 	  				`json:"user_id" form:"user_id"`
	StartDate 		string 	  				`json:"start_date" form:"start_date"`
	GraduateDate 	string 	  				`json:"graduate_date" form:"graduate_date"`
}

func RequestToEntity(class ClassesRequest) classes.ClassessEntity{
	return classes.ClassessEntity{
		Name: 			class.Name,
		UserID: 		class.UserID,
		StartDate: 		class.StartDate,
		GraduateDate: 	class.GraduateDate,
	}
}