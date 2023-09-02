package data

import (
	mentee "immersive_project/klp3/features/mentee/data"
	"immersive_project/klp3/features/menteelogs"
	user "immersive_project/klp3/features/users/data"

	"gorm.io/gorm"
)

type MenteeLogModel struct {
	gorm.Model
	MenteeID uint   `gorm:"column:mentee_id;default:1;not null"`
	UserID   uint   `gorm:"column:user_id;not null"`
	Log      string `gorm:"column:log;not null"`
	Status   string `gorm:"column:status"`
	Users    user.UserModel `gorm:"foreignKey:UserID"`
	Mentee   mentee.MenteeModel `gorm:"foreignKey:MenteeID"` 
}

func EntityToModel(menteeLog menteelogs.MenteeLogEntity)MenteeLogModel{
	return MenteeLogModel{
		MenteeID: menteeLog.MenteeID,
		UserID:   menteeLog.UserID,
		Log:      menteeLog.Log,
		Status:   menteeLog.Status,
		Users: 	  user.EntityToModel(menteeLog.Users),
		Mentee:   mentee.EntityToModel(menteeLog.Mentee),
	}
}

func ModelToEntity(menteeLog MenteeLogModel)menteelogs.MenteeLogEntity{
	return menteelogs.MenteeLogEntity{
		Id:        menteeLog.ID,
		CreatedAt: menteeLog.CreatedAt,
		UpdatedAt: menteeLog.UpdatedAt,
		DeletedAt: menteeLog.DeletedAt.Time,
		MenteeID:  menteeLog.MenteeID,
		UserID:    menteeLog.UserID,
		Status:    menteeLog.Status,
		Log:       menteeLog.Log,
		Users: 	   user.ModelToEntity(menteeLog.Users),
		Mentee:    mentee.ModelToEntity(menteeLog.Mentee),
	}
}
func ModelToEntityAll(menteeLog []MenteeLogModel)[]menteelogs.MenteeLogEntity{
	var menteeLogs []menteelogs.MenteeLogEntity
	for _,value:=range menteeLog{
		menteeLogs=append(menteeLogs, ModelToEntity(value))
	}
	return menteeLogs
}