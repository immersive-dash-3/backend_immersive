package handler

import (
	class "immersive_project/klp3/features/classes/handler"
	"immersive_project/klp3/features/menteelogs"
	user "immersive_project/klp3/features/users/handler"
)

type MenteeLogResponse struct {
	Id        uint              `json:"id" form:"id"`
	CreatedAt string       		`json:"created_at,omitempty"`
	MenteeID  uint              `json:"mentee_id" form:"mentee_id"`
	UserID    uint              `json:"user_id" form:"user_id"`
	Status    string            `json:"status" form:"status"`
	Log       string            `json:"log" form:"log"`
	Users     user.UserResponse `json:"users,omitempty"`
}

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
	HomeAddress     string                `json:"home_address" form:"home_address"`
	EmergencyName   string                `json:"emergency_name" form:"emergency_name"`
	EmergencyPhone  string                `json:"emergency_phone" form:"emergency_phone"`
	EmergencyStatus string                `json:"emergency_status" form:"emergency_status"`
	Major           string                `json:"major" form:"major"`
	Graduate        string                `json:"graduate" form:"graduate"`
	Institution     string                `json:"institution" form:"institution"`
	Class           class.ClassesResponse `json:"class,omitempty"`
	Status          StatusResponse        `json:"status,omitempty"`
	MenteeLog       []MenteeLogResponse   `json:"logs,omitempty"`
}

type StatusResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func StatusEntityToResponse(status menteelogs.StatusEntity) StatusResponse {
	return StatusResponse{
		Id:   status.Id,
		Name: status.Name,
	}
}

func EntityToResponse(menteeLog menteelogs.MenteeLogEntity) MenteeLogResponse {

	formattedTime := menteeLog.CreatedAt.Format("Jan, 2 2006")

	return MenteeLogResponse{
		Id:        menteeLog.Id,
		CreatedAt: formattedTime,
		MenteeID:  menteeLog.MenteeID,
		UserID:    menteeLog.UserID,
		Status:    menteeLog.Status,
		Log:       menteeLog.Log,
		Users:     user.UserEntityToResponse(menteeLog.Users),
	}
}

func EntityToResponseAll(menteeLog []menteelogs.MenteeLogEntity) []MenteeLogResponse {
	var menteeAll []MenteeLogResponse
	for _, value := range menteeLog {
		menteeAll = append(menteeAll, EntityToResponse(value))
	}
	return menteeAll
}

func EntityResponseById(mente menteelogs.MenteeEntity) MenteeResponse {
	var logs []MenteeLogResponse
	for _, value := range mente.MenteeLog {
		logs = append(logs, EntityToResponse(value))
	}
	return MenteeResponse{
		Id:          mente.Id,
		FirstName:   mente.FirstName,
		LastName:    mente.LastName,
		Email:       mente.Email,
		PhoneNumber: mente.PhoneNumber,
		Telegram:    mente.Telegram,
		Discord:     mente.Discord,
		ClassID:     mente.ClassID,
		StatusID:    mente.StatusID,
		Major:       mente.Major,
		Graduate:    mente.Graduate,
		Institution: mente.Institution,
		Class:       class.EntityToResponse(mente.Class),
		Status:      StatusEntityToResponse(mente.Status),
		MenteeLog:   logs,
	}
}
