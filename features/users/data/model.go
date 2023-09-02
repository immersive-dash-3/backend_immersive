package data

import (
	"immersive_project/klp3/features/users"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	FullName     	string `gorm:"column:full_name;not null"`
	Email    		string `gorm:"column:email;unique;not null"`
	Password 		string `gorm:"column:password;not null"`
	Team 			string `gorm:"type:enum('Academic','People','Placement','Admission');default:'Academic';column:team;not null"`
	Role   			string `gorm:"type:enum('User','Admin');default:'User';column:role;not null"`
	Address  		string `gorm:"column:address"`
	Status  		string `gorm:"type:enum('Active','Not-Active','Deleted');default:'Active';column:status;not null"`
}

func EntityToModel(user users.UserEntity)UserModel{
	return UserModel{
		FullName: user.FullName,
		Email:    user.Email,
		Password: user.Password,
		Team:     user.Team,
		Role:     user.Role,
		Address:  user.Address,
		Status:   user.Status,
	}
}
func ModelToEntity(user UserModel)users.UserEntity{
	return users.UserEntity{
		Id:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt.Time,
		FullName:  user.FullName,
		Email:     user.Email,
		Team:      user.Team,
		Password:  user.Password,
		Role:      user.Role,
		Address:   user.Address,
		Status:    user.Status,
	}
}

func ModelToEntityAll(user []UserModel)[]users.UserEntity{
	var userAll []users.UserEntity
	for _,value:=range user{
		userAll=append(userAll, ModelToEntity(value))
	}
	return userAll
}