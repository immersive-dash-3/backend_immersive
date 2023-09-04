package handler

import "immersive_project/klp3/features/menteelogs"

type MenteeLogRequest struct {
	MenteeID uint   `json:"mentee_id" form:"mentee_id"`
	UserID   uint   `json:"user_id" form:"user_id"`
	Status   string `json:"status" form:"status"`
	Log      string `json:"log" form:"log"`
}

func RequestToEntity(menteeLog MenteeLogRequest) menteelogs.MenteeLogEntity{
	return menteelogs.MenteeLogEntity{
		MenteeID:  menteeLog.MenteeID,
		UserID:    menteeLog.UserID,
		Status:  menteeLog.Status,
		Log:       menteeLog.Log,
	}
}