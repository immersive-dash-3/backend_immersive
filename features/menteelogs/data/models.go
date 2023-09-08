package data

import (
	class "immersive_project/klp3/features/classes/data"
	"immersive_project/klp3/features/menteelogs"
	user "immersive_project/klp3/features/users/data"

	"gorm.io/gorm"
)

type MenteeLog struct {
	gorm.Model
	MenteeID uint   `gorm:"column:mentee_id;default:1;not null"`
	UserID   uint   `gorm:"column:user_id;not null"`
	Log      string `gorm:"column:log;not null"`
	StatusID uint `gorm:"column:status_id"`
	Users    user.User `gorm:"foreignKey:UserID"`
	Status   Status `gorm:"foreignKey:StatusID"`
}

type Mentee struct {
	gorm.Model
	FirstName        string `gorm:"column:first_name;not nul"`
	LastName        string `gorm:"column:last_name"`
	Gender          string `gorm:"column:gender;not nul"`
	Email           string `gorm:"column:email;unique;not nul"`
	PhoneNumber     string `gorm:"column:phone_number;not nul"`
	Telegram        string `gorm:"column:telegram"`
	ClassID         uint   `gorm:"column:class_id"`
	StatusID        uint 	`gorm:"column:status_id"`
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
	Status          Status `gorm:"foreignKey:StatusID"`
	Class           class.Classes `gorm:"foreignKey:ClassID"`
	MenteeLogs      []MenteeLog `gorm:"foreignKey:MenteeID"`
}

type Status struct{
	gorm.Model
	Name            string	  `json:"name" form:"name"`
}
func StatusEntityToModel(status menteelogs.StatusEntity)Status{
	return Status{
		Name:  status.Name,
	}
}

func StatusModelToEntity(status Status)menteelogs.StatusEntity{
	return menteelogs.StatusEntity{
		Id:              status.ID,
		CreatedAt:       status.CreatedAt,
		UpdatedAt:       status.UpdatedAt,
		DeletedAt:       status.DeletedAt.Time,
		Name:  status.Name,
	}
}

func ModelMenteeToEntity(mente Mentee)menteelogs.MenteeEntity{
	var logs []menteelogs.MenteeLogEntity
	for _,value:= range mente.MenteeLogs{
		logs = append(logs, ModelToEntity(value))
	}
	return menteelogs.MenteeEntity{
		Id:              mente.ID,
		CreatedAt:       mente.CreatedAt,
		UpdatedAt:       mente.UpdatedAt,
		DeletedAt:       mente.DeletedAt.Time,
		FirstName:       mente.FirstName,
		LastName:        mente.LastName,
		Gender:          mente.Gender,
		Email:           mente.Email,
		PhoneNumber:     mente.PhoneNumber,
		Telegram:        mente.Telegram,
		Discord:         mente.Discord,
		ClassID:         mente.ClassID,
		StatusID:        mente.StatusID,
		EducationType:   mente.EducationType,
		CurrentAddress:  mente.CurrentAddress,
		HomeAddress:     mente.HomeAddress,
		EmergencyName:   mente.EmergencyName,
		EmergencyPhone:  mente.EmergencyPhone,
		EmergencyStatus: mente.EmergencyStatus,
		Major:           mente.Major,
		Graduate:        mente.Graduate,
		Institution:     mente.Graduate,
		Class:           class.ModelToEntity(mente.Class),
		Status: StatusModelToEntity(mente.Status),
		MenteeLog: logs,	
	}
}

func EntityToModel(menteeLog menteelogs.MenteeLogEntity)MenteeLog{
	return MenteeLog{
		MenteeID: menteeLog.MenteeID,
		UserID:   menteeLog.UserID,
		Log:      menteeLog.Log,
		StatusID:   menteeLog.StatusID,
		Status: StatusEntityToModel(menteeLog.Status),
		Users: 	  user.EntityToModel(menteeLog.Users),
	}
}

func ModelToEntity(menteeLog MenteeLog)menteelogs.MenteeLogEntity{
	return menteelogs.MenteeLogEntity{
		Id:        menteeLog.ID,
		CreatedAt: menteeLog.CreatedAt,
		UpdatedAt: menteeLog.UpdatedAt,
		DeletedAt: menteeLog.DeletedAt.Time,
		MenteeID:  menteeLog.MenteeID,
		UserID:    menteeLog.UserID,
		StatusID: menteeLog.StatusID,
		Status:    StatusModelToEntity(menteeLog.Status),
		Log:       menteeLog.Log,
		Users:     user.ModelToEntity(menteeLog.Users),
	}
}

func ListModelToEntity(mentee []MenteeLog)[]menteelogs.MenteeLogEntity{
	var input []menteelogs.MenteeLogEntity
	for _,value:=range mentee{
		input = append(input, ModelToEntity(value))
	}
	return input
}