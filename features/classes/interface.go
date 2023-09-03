package classes

import (
	"immersive_project/klp3/features/users"
	"time"
)

type ClassessEntity struct {
	Id        		uint      				`json:"id" form:"id"`
	CreatedAt 		time.Time 				`json:"created_at,omitempty"`
	UpdatedAt 		time.Time 				`json:"updated_at,omitempty"`
	DeletedAt 		time.Time 				`json:"deleted_at,omitempty"`
	Name 	  		string	  				`json:"name" form:"name"`
	UserID 			uint 	  				`json:"user_id" form:"user_id"`
	StartDate 		string 	  				`json:"start_date" form:"start_date"`
	GraduateDate 	string 	  				`json:"graduate_date" form:"graduate_date"`
	Users			users.UserEntity 	  	`json:"users,omitempty"`
}