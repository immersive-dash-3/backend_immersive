package handler

import (
	mentee "immersive_project/klp3/features/mentee/handler"
	"immersive_project/klp3/features/menteelogs"
	user "immersive_project/klp3/features/users/handler"
	"time"
)

type MenteeLogResponse struct {
	Id        uint      `json:"id" form:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	MenteeID  uint      `json:"mentee_id" form:"mentee_id"`
	UserID    uint      `json:"user_id" form:"user_id"`
	Status    string    `json:"status" form:"status"`
	Log       string    `json:"log" form:"log"`
	Users     user.UserResponse `json:"users,omitempty"`
	Mentee    mentee.MenteeResponse `json:"mentee,omitempty"`
}

func EntityToResponse(menteeLog menteelogs.MenteeLogEntity)MenteeLogResponse{
	return MenteeLogResponse{
		Id:        menteeLog.Id,
		CreatedAt: menteeLog.CreatedAt,
		MenteeID:  menteeLog.MenteeID,
		UserID:    menteeLog.UserID,
		Status:    menteeLog.Status,
		Log:       menteeLog.Log,
		Users: 	   user.UserEntityToResponse(menteeLog.Users),
		Mentee:    mentee.EntityResponseById(menteeLog.Mentee),
	}
}

func EntityToResponseAll(menteeLog []menteelogs.MenteeLogEntity)[]MenteeLogResponse{
	var menteeAll []MenteeLogResponse
	for _,value:=range menteeLog{
		menteeAll=append(menteeAll, EntityToResponse(value))
	}
	return menteeAll
}