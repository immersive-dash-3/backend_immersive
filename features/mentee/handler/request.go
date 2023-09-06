package handler

import (
	"immersive_project/klp3/features/mentee"
)

type MenteeRequest struct {
	FirstName       string `json:"first_name" form:"first_name"`
	LastName        string `json:"last_name" form:"last_name"`
	Gender          string `json:"gender" form:"gender"`
	Email           string `json:"email" form:"email"`
	PhoneNumber     string `json:"phone_number" form:"phone_number"`
	Telegram        string `json:"telegram" form:"telegram"`
	Discord         string `json:"discord" form:"discord"`
	ClassID         uint   `json:"class_id" form:"class_id"`
	StatusID        uint   `json:"status_id" from:"status_id"`
	EducationType   string `json:"education_type" form:"education_type"`
	CurrentAddress  string `json:"current_address" form:"current_address"`
	EmergencyName   string `json:"emergency_name" form:"emergency_name"`
	EmergencyPhone  string `json:"emergency_phone" form:"emergency_phone"`
	EmergencyStatus string `json:"emergency_status" form:"emergency_status"`
	Major           string `json:"major" form:"major"`
	Graduate        string `json:"graduate" form:"graduate"`
}

func RequestToEntity(mente MenteeRequest) mentee.MenteeEntity {
	return mentee.MenteeEntity{
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
		EmergencyName:   mente.EmergencyName,
		EmergencyPhone:  mente.EmergencyPhone,
		EmergencyStatus: mente.EmergencyStatus,
		Major:           mente.Major,
		Graduate:        mente.Graduate,
	}
}
