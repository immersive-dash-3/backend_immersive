package menteelogs

import (
	mentee "immersive_project/klp3/features/mentee"
	user "immersive_project/klp3/features/users"
	"time"
)

type MenteeLogEntity struct {
	Id        	uint      	`json:"id" form:"id"`
	CreatedAt 	time.Time 	`json:"created_at,omitempty"`
	UpdatedAt 	time.Time 	`json:"updated_at,omitempty"`
	DeletedAt 	time.Time 	`json:"deleted_at,omitempty"`
	MenteeID 	uint 		`json:"mentee_id" form:"mentee_id"`
	UserID 		uint 		`json:"user_id" form:"user_id"`
	Status      string      `json:"status" form:"status"`
	Log 		string 		`json:"log" form:"log"`
	Users       user.UserEntity `json:"users,omitempty"`
	Mentee      mentee.MenteeEntity `json:"mentee,omitempty"`
}