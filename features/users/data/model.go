package data

import (
	"immersive_project/klp3/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"column:name;not null"`
	Email    string `gorm:"column:email;unique;not null"`
	Password string `gorm:"column:password;not null"`
	Team     string `gorm:"type:enum('Academic','People','Placement','Admission');default:'Academic';column:team;not null"`
	Role     string `gorm:"type:enum('Default','Admin');default:'Default';column:role;not null"`
	Address  string `gorm:"column:address"`
	Status   string `gorm:"type:enum('Active','Non-Active','Deleted');default:'Active';column:status;not null"`
}

func EntityToModel(user users.UserEntity) User {
	return User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Team:     user.Team,
		Role:     user.Role,
		Address:  user.Address,
		Status:   user.Status,
	}
}
func ModelToEntity(user User) users.UserEntity {
	return users.UserEntity{
		Id:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt.Time,
		Name:      user.Name,
		Email:     user.Email,
		Team:      user.Team,
		Password:  user.Password,
		Role:      user.Role,
		Address:   user.Address,
		Status:    user.Status,
	}
}

func ModelToEntityAll(user []User) []users.UserEntity {
	var userAll []users.UserEntity
	for _, value := range user {
		userAll = append(userAll, ModelToEntity(value))
	}
	return userAll
}
