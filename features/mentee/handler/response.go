package handler

import (
	class "immersive_project/klp3/features/classes/handler"
	"immersive_project/klp3/features/mentee"
)
type MenteeResponse struct {
	Id              uint   `json:"id" form:"id"`
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
	Class           class.ClassesResponse `json:"class,omitempty"`

}

func EntityResponseById(mente mentee.MenteeEntity)MenteeResponse{
	return MenteeResponse{
		Id:              mente.Id,
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
		Class: 			 class.ResponseToEntity(mente.Class),
	}
}

func EntityToResponseAll(mente mentee.MenteeEntity)MenteeResponse{
	return MenteeResponse{
		Id: 			 mente.Id,
		FullName:        mente.FullName,
		Gender:          mente.Gender,
		ClassID:         mente.ClassID,
		Status:          mente.Status,
		Category: 		 mente.Category,
		Class: 			 class.ResponseToEntity(mente.Class),
	}
}
func MenteeEntityToResponseAll(mente []mentee.MenteeEntity)[]MenteeResponse{
	var mentee []MenteeResponse
	for _,value:=range mente{
		mentee=append(mentee, EntityToResponseAll(value))
	}
	return mentee
}