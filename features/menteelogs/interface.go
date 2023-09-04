package menteelogs

import (
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
	StatusID      string      `json:"status_id" form:"status_id"`
	Log 		string 		`json:"log" form:"log"`
	Users       user.UserEntity `json:"users,omitempty"`

}