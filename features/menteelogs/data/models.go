package data

import (
	"immersive_project/klp3/features/menteelogs"
	user "immersive_project/klp3/features/users/data"

	"gorm.io/gorm"
)

type MenteeLog struct {
	gorm.Model
	MenteeID uint   `gorm:"column:mentee_id;default:1;not null"`
	UserID   uint   `gorm:"column:user_id;not null"`
	Log      string `gorm:"column:log;not null"`
	StatusID   string `gorm:"column:status"`
	Users    user.User `gorm:"foreignKey:UserID"`
}

func EntityToModel(menteeLog menteelogs.MenteeLogEntity)MenteeLog{
	return MenteeLog{
		MenteeID: menteeLog.MenteeID,
		UserID:   menteeLog.UserID,
		Log:      menteeLog.Log,
		StatusID:   menteeLog.StatusID,
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
		StatusID:  menteeLog.StatusID,
		Log:       menteeLog.Log,
		Users:     user.ModelToEntity(menteeLog.Users),
	}
}