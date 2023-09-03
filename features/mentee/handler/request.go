package handler

import "immersive_project/klp3/features/mentee"

type MenteeRequest struct {
	FullName        string `json:"full_name" form:"full_name"`
	NickName        string `json:"nick_name" form:"nick_name"`
	Gender          string `json:"gender" form:"gender"`
	Email           string `json:"email" form:"email"`
	PhoneNumber     string `json:"phone_number" form:"phone_number"`
	Telegram        string `json:"telegram" form:"telegram"`
	ClassID         uint   `json:"class_id" form:"class_id"`
	Status          string `json:"status" from:"status"`
	Category        string `json:"category" form:"category"`
	CurrentAddress  string `json:"current_address" form:"current_address"`
	HomeAddress     string `json:"home_address" form:"home_address"`
	EmergencyName   string `json:"emergency_name" form:"emergency_name"`
	EmergencyPhone  string `json:"emergency_phone" form:"emergency_phone"`
	EmergencyStatus string `json:"emergency_status" form:"emergency_status"`
	Major           string `json:"major" form:"major"`
	Graduate        string `json:"graduate" form:"graduate"`
	Institution     string `json:"institution" form:"institution"`
}

func RequestToEntity(mente MenteeRequest) mentee.MenteeEntity{
	return mentee.MenteeEntity{
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
	}
}
