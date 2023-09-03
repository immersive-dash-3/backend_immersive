package data

import (
	"immersive_project/klp3/features/classes"
	user "immersive_project/klp3/features/users/data"

	"gorm.io/gorm"
)

type ClassesModel struct {
	gorm.Model
	Name 			string 			`gorm:"column:name;not nul"`
	UserID 			uint   			`gorm:"column:user_id;not nul"`
	StartDate 		string 			`gorm:"column:start_date;not nul"`
	GraduateDate 	string 			`gorm:"column:graduate_date"`
	Users           user.UserModel 	`gorm:"foreignKey:UserID"`
}

func EntityToModel(class classes.ClassessEntity)ClassesModel{
	return ClassesModel{
		Name: 			class.Name,
		UserID: 		class.UserID,
		StartDate: 		class.StartDate,
		GraduateDate: 	class.GraduateDate,
		Users: 			user.EntityToModel(class.Users),
	}
}

func ModelToEntity(class ClassesModel)classes.ClassessEntity{
	return classes.ClassessEntity{
		Id:        		class.ID,
		CreatedAt: 		class.CreatedAt,
		UpdatedAt: 		class.UpdatedAt,
		DeletedAt: 		class.DeletedAt.Time,
		Name:      		class.Name,
		UserID: 		class.UserID,
		StartDate: 		class.StartDate,
		GraduateDate: 	class.GraduateDate,
		Users: 			user.ModelToEntity(class.Users),
		
	}
}

func ModelToEntityAll(class []ClassesModel)[]classes.ClassessEntity{
	var classes []classes.ClassessEntity
	for _,value:=range class{
		classes=append(classes, ModelToEntity(value))
	}
	return classes
}