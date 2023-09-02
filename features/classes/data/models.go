package data

import (
	"immersive_project/klp3/features/classes"

	"gorm.io/gorm"
)

type ClassesModel struct {
	gorm.Model
	Name string `gorm:"column:name;not nul"`
}

func EntityToModel(class classes.ClassessEntity)ClassesModel{
	return ClassesModel{
		Name: class.Name,
	}
}

func ModelToEntity(class ClassesModel)classes.ClassessEntity{
	return classes.ClassessEntity{
		Id:        class.ID,
		CreatedAt: class.CreatedAt,
		UpdatedAt: class.UpdatedAt,
		DeletedAt: class.DeletedAt.Time,
		Name:      class.Name,
	}
}

func ModelToEntityAll(class []ClassesModel)[]classes.ClassessEntity{
	var classes []classes.ClassessEntity
	for _,value:=range class{
		classes=append(classes, ModelToEntity(value))
	}
	return classes
}