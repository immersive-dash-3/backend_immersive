package classes

import (
	"time"
)

type ClassessEntity struct {
	Id        		uint      				`json:"id" form:"id"`
	CreatedAt 		time.Time 				`json:"created_at,omitempty"`
	UpdatedAt 		time.Time 				`json:"updated_at,omitempty"`
	DeletedAt 		time.Time 				`json:"deleted_at,omitempty"`
	Name 	  		string	  				`json:"name" form:"name"`
}