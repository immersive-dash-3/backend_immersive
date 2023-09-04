package data

import (
	class "immersive_project/klp3/features/classes/data"
	mentees "immersive_project/klp3/features/mentee"
	menteeLog "immersive_project/klp3/features/menteelogs/data"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	FirstName        string `gorm:"column:first_name;not nul"`
	LastName        string `gorm:"column:last_name"`
	Gender          string `gorm:"column:gender;not nul"`
	Email           string `gorm:"column:email;unique;not nul"`
	PhoneNumber     string `gorm:"column:phone_number;not nul"`
	Telegram        string `gorm:"column:telegram"`
	ClassID         uint   `gorm:"column:class_id"`
	StatusID        uint 	`gorm:"column:status"`
	Discord 		string `gorm:"column:discord"`
	EducationType   string `gorm:"type:enum('Informatics','Non-Informatics');default:'Informatics';column:education_type"`
	CurrentAddress  string `gorm:"column:current_address"`
	HomeAddress     string `gorm:"column:home_address"`
	EmergencyName   string `gorm:"column:emergency_name;not nul"`
	EmergencyPhone  string `gorm:"column:emergency_phone"`
	EmergencyStatus string `gorm:"column:emergency_status"`
	Major           string `gorm:"column:major"`
	Graduate        string `gorm:"column:graduate"`
	Institution     string `gorm:"column:institution"`
	Class           class.Classes `gorm:"foreignKey:ClassID"`
	Status          Status `gorm:"foreignKey:StatusID"`
	MenteeLog       []menteeLog.MenteeLog `gorm:"foreignKey:MenteeID"`
}

type Status struct{
	gorm.Model
	Name            string	  `json:"name" form:"name"`
}
func StatusEntityToModel(status mentees.StatusEntity)Status{
	return Status{
		Name:  status.Name,
	}
}

func StatusModelToEntity(status Status)mentees.StatusEntity{
	return mentees.StatusEntity{
		Name:  status.Name,
	}
}

func EntityToModel(mentee mentees.MenteeEntity)Mentee{
	return Mentee{
		FirstName:       mentee.FirstName,
		LastName:        mentee.LastName,
		Gender:          mentee.Gender,
		Email:           mentee.Email,
		PhoneNumber:     mentee.PhoneNumber,
		Telegram:        mentee.Telegram,
		ClassID:         mentee.ClassID,
		StatusID:        mentee.StatusID,
		Discord:         mentee.Discord,
		EducationType:   mentee.EducationType,
		CurrentAddress:  mentee.CurrentAddress,
		HomeAddress:     mentee.HomeAddress,
		EmergencyName:   mentee.EmergencyName,
		EmergencyPhone:  mentee.EmergencyPhone,
		EmergencyStatus: mentee.EmergencyStatus,
		Major:           mentee.Major,
		Graduate:        mentee.Graduate,
		Institution:     mentee.Institution,
		Class:           class.EntityToModel(mentee.Class),
		Status:          StatusEntityToModel(mentee.Status),
	}
}

func ModelToEntity(mentee Mentee)mentees.MenteeEntity{
	return mentees.MenteeEntity{
		FirstName:       mentee.FirstName,
		LastName:        mentee.LastName,
		Gender:          mentee.Gender,
		Email:           mentee.Email,
		PhoneNumber:     mentee.PhoneNumber,
		Telegram:        mentee.Telegram,
		ClassID:         mentee.ClassID,
		StatusID:        mentee.StatusID,
		Discord:         mentee.Discord,
		EducationType:   mentee.EducationType,
		CurrentAddress:  mentee.CurrentAddress,
		HomeAddress:     mentee.HomeAddress,
		EmergencyName:   mentee.EmergencyName,
		EmergencyPhone:  mentee.EmergencyPhone,
		EmergencyStatus: mentee.EmergencyStatus,
		Major:           mentee.Major,
		Graduate:        mentee.Graduate,
		Institution:     mentee.Institution,
		Class:           class.ModelToEntity(mentee.Class),
		Status:          StatusModelToEntity(mentee.Status),
	}
}

func ModelToEntityAll(mente []Mentee)[]mentees.MenteeEntity{
	var menteeAll []mentees.MenteeEntity
	for _,value:=range mente{
		menteeAll=append(menteeAll, ModelToEntity(value))
	}
	return menteeAll
}