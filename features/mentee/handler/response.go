package handler

import (
	class "immersive_project/klp3/features/classes/handler"
	"immersive_project/klp3/features/mentee"
)

type MenteeResponse struct {
	Id              uint                  `json:"id" form:"id"`
	FirstName       string                `json:"first_name" form:"first_name"`
	LastName        string                `json:"last_name" form:"last_name"`
	Gender          string                `json:"gender" form:"gender"`
	Email           string                `json:"email" form:"email"`
	PhoneNumber     string                `json:"phone_number" form:"phone_number"`
	Telegram        string                `json:"telegram" form:"telegram"`
	Discord         string                `json:"discord" form:"discord"`
	ClassID         uint                  `json:"class_id" form:"class_id"`
	StatusID        uint                  `json:"status_id" from:"status_id"`
	EducationType   string                `json:"education_type" form:"education_type"`
	CurrentAddress  string                `json:"current_address" form:"current_address"`
	EmergencyName   string                `json:"emergency_name" form:"emergency_name"`
	EmergencyPhone  string                `json:"emergency_phone" form:"emergency_phone"`
	EmergencyStatus string                `json:"emergency_status" form:"emergency_status"`
	Major           string                `json:"major" form:"major"`
	Graduate        string                `json:"graduate" form:"graduate"`
	Class           class.ClassesResponse `json:"class,omitempty"`
	Status          StatusResponse        `json:"status,omitempty"`
}

type MenteeResponseAll struct {
	Id            uint                  `json:"id" form:"id"`
	FirstName     string                `json:"first_name" form:"first_name"`
	LastName      string                `json:"last_name" form:"last_name"`
	Gender        string                `json:"gender" form:"gender"`
	EducationType string                `json:"education_type" form:"education_type"`
	Class         class.ClassesResponse `json:"class,omitempty"`
	Status        StatusResponse        `json:"status,omitempty"`
}

type StatusResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func StatusEntityToResponse(status mentee.StatusEntity) StatusResponse {
	return StatusResponse{
		Id:   status.Id,
		Name: status.Name,
	}
}

func EntityToResponseAll(mente mentee.MenteeEntity) MenteeResponseAll {
	return MenteeResponseAll{
		Id:            mente.Id,
		FirstName:     mente.FirstName,
		LastName:      mente.LastName,
		Gender:        mente.Gender,
		EducationType: mente.EducationType,
		Class:         class.EntityToResponse(mente.Class),
		Status:        StatusEntityToResponse(mente.Status),
	}
}

func EntityToResponse(input mentee.MenteeEntity) MenteeResponse {
	var resp = MenteeResponse{
		Id:              input.Id,
		FirstName:       input.FirstName,
		LastName:        input.LastName,
		Gender:          input.Gender,
		Email:           input.Email,
		PhoneNumber:     input.PhoneNumber,
		Telegram:        input.Telegram,
		Discord:         input.Discord,
		ClassID:         input.ClassID,
		StatusID:        input.StatusID,
		EducationType:   input.EducationType,
		CurrentAddress:  input.CurrentAddress,
		EmergencyName:   input.EmergencyName,
		EmergencyPhone:  input.EmergencyPhone,
		EmergencyStatus: input.EmergencyStatus,
		Major:           input.Major,
		Graduate:        input.Graduate,
		Class:           class.EntityToResponse(input.Class),
		Status:          StatusEntityToResponse(input.Status),
	}
	return resp
}

/*
func MenteeEntityToResponseAll(mente []mentee.MenteeEntity) []MenteeResponse {
	var mentee []MenteeResponse
	for _, value := range mente {
		mentee = append(mentee, EntityToResponseAll(value))
	}
	return mentee
}
*/
