package data

import (
	"immersive_project/klp3/features/classes"

	"gorm.io/gorm"
)

type Classes struct {
	gorm.Model
	Name   string `gorm:"column:name;not nul"`
	UserID uint   `gorm:"column:user_id;default:1"`
	User  User    `gorm:"foreignKey:UserID"`
}

type User struct {
	gorm.Model
	Name     string `gorm:"column:name;not null"`
}

func UserModelToEntity(user User)classes.UserEntity{
	return classes.UserEntity{
		Id:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt.Time,	
		Name: user.Name,	
	}
}

func UserEntityToModel(user classes.UserEntity)User{
	return User{
		Name: user.Name,	
	}
}

func EntityToModel(class classes.ClassessEntity) Classes {
	return Classes{
		Name:   class.Name,
		UserID: class.UserID,
		User: UserEntityToModel(class.User),
	}
}

func ModelToEntity(class Classes) classes.ClassessEntity {
	return classes.ClassessEntity{
		Id:        class.ID,
		CreatedAt: class.CreatedAt,
		UpdatedAt: class.UpdatedAt,
		DeletedAt: class.DeletedAt.Time,
		Name:      class.Name,
		UserID:    class.UserID,
		User: UserModelToEntity(class.User),
	}
}

func ModelToEntityAll(class []Classes) []classes.ClassessEntity {
	var classes []classes.ClassessEntity
	for _, value := range class {
		classes = append(classes, ModelToEntity(value))
	}
	return classes
}
