package data

import (
	class "immersive_project/klp3/features/classes/data"
	"immersive_project/klp3/features/mentee"

	"gorm.io/gorm"
)

type MenteeModel struct {
	gorm.Model
	FullName        string `gorm:"column:full_name;not nul"`
	NickName        string `gorm:"column:nick_name"`
	Gender          string `gorm:"type:enum('male','female');default:'male';column:gender;not nul"`
	Email           string `gorm:"column:email;unique;not nul"`
	PhoneNumber     string `gorm:"column:phone_number;not nul"`
	Telegram        string `gorm:"column:telegram"`
	ClassID         uint   `gorm:"column:class_id"`
	Status          string `gorm:"column:status"`
	Category        string `gorm:"type:enum('Informatics','Non-Informatics');default:'Informatics';column:category"`
	CurrentAddress  string `gorm:"column:current_address"`
	HomeAddress     string `gorm:"column:home_address"`
	EmergencyName   string `gorm:"column:emergency_name;not nul"`
	EmergencyPhone  string `gorm:"column:emergency_phone"`
	EmergencyStatus string `gorm:"type:enum('Orang Tua','Kakek Nenek','Saudara dari Orang Tua');default:'Orang Tua';column:emergency_status"`
	Major           string `gorm:"column:major"`
	Graduate        string `gorm:"column:graduate"`
	Institution     string `gorm:"column:institution"`
	Class           class.ClassesModel `gorm:"foreignKey:ClassID"`
}

func EntityToModel(mentee mentee.MenteeEntity)MenteeModel{
	return MenteeModel{
		FullName:        mentee.FullName,
		NickName:        mentee.NickName,
		Gender:          mentee.Gender,
		Email:           mentee.Email,
		PhoneNumber:     mentee.PhoneNumber,
		Telegram:        mentee.Telegram,
		ClassID:         mentee.ClassID,
		Status:          mentee.Status,
		Category:        mentee.Category,
		CurrentAddress:  mentee.CurrentAddress,
		HomeAddress:     mentee.HomeAddress,
		Institution: 	 mentee.Institution,
		EmergencyName:   mentee.EmergencyName,
		EmergencyPhone:  mentee.EmergencyPhone,
		EmergencyStatus: mentee.EmergencyStatus,
		Major:           mentee.Major,
		Graduate:        mentee.Graduate,
		Class: 			 class.EntityToModel(mentee.Class),
	}
}


func ModelToEntity(mente MenteeModel)mentee.MenteeEntity{
	return mentee.MenteeEntity{
		Id:              mente.ID,
		CreatedAt:       mente.CreatedAt,
		UpdatedAt:       mente.UpdatedAt,
		DeletedAt:       mente.DeletedAt.Time,
		FullName:        mente.FullName,
		NickName:        mente.NickName,
		Gender:          mente.Gender,
		Email:           mente.Email,
		PhoneNumber:     mente.PhoneNumber,
		Telegram:        mente.Telegram,
		ClassID:         mente.ClassID,
		Status:          mente.Status,
		Category:        mente.Category,
		CurrentAddress:  mente.CurrentAddress,
		HomeAddress:     mente.HomeAddress,
		EmergencyName:   mente.EmergencyName,
		EmergencyPhone:  mente.EmergencyPhone,
		EmergencyStatus: mente.EmergencyStatus,
		Major:           mente.Major,
		Graduate:        mente.Graduate,
		Institution: 	 mente.Institution,
		Class: 			 class.ModelToEntity(mente.Class),
	}
}

func ModelToEntityAll(mente []MenteeModel)[]mentee.MenteeEntity{
	var menteeAll []mentee.MenteeEntity
	for _,value:=range mente{
		menteeAll=append(menteeAll, ModelToEntity(value))
	}
	return menteeAll
}