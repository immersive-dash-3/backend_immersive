package handler

import (
	"immersive_project/klp3/features/classes"
	user "immersive_project/klp3/features/users/handler"
)


type ClassesResponse struct {
	Id   			uint    			`json:"id" form:"id"`
	Name 			string  			`json:"name" form:"name"`
	UserID 			uint 				`json:"user_id" form:"user_id"`
	StartDate 		string 				`json:"start_date" form:"start_date"`
	GraduateDate 	string 				`json:"graduate_date" form:"graduate_date"`
	Users 			user.UserResponse 	`json:"users,omitempty"`
}

func ResponseToEntity(class classes.ClassessEntity) ClassesResponse{
	return ClassesResponse{
		Id: 			class.Id,
		Name: 			class.Name,
		UserID: 		class.UserID,
		StartDate: 		class.StartDate,
		GraduateDate: 	class.GraduateDate,
		Users: 			user.UserEntityToResponse(class.Users),

	}
}

func ResponseToEntityAll(class []classes.ClassessEntity)[]ClassesResponse{
	var classes []ClassesResponse
	for _,value:=range class{
		classes=append(classes, ResponseToEntity(value))
	}
	return classes
}