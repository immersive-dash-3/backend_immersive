package data

import (
	"immersive_project/klp3/features/classes"

	"gorm.io/gorm"
)

type Classes struct {
	gorm.Model
	Name   string `gorm:"column:name;not nul"`
	UserID uint   `gorm:"column:user_id"`
}

func EntityToModel(class classes.ClassessEntity) Classes {
	return Classes{
		Name: class.Name,
	}
}

func ModelToEntity(class Classes) classes.ClassessEntity {
	return classes.ClassessEntity{
		Id:        class.ID,
		CreatedAt: class.CreatedAt,
		UpdatedAt: class.UpdatedAt,
		DeletedAt: class.DeletedAt.Time,
		Name:      class.Name,
	}
}

func ModelToEntityAll(class []Classes) []classes.ClassessEntity {
	var classes []classes.ClassessEntity
	for _, value := range class {
		classes = append(classes, ModelToEntity(value))
	}
	return classes
}
