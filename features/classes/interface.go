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

type ClassDataInterface interface{
	Insert(input ClassessEntity)(uint,error)
	SelectById(id uint)(ClassessEntity,error)
	SelectAll(page, pageSize int)([]ClassessEntity,error)
	Update(id uint, input ClassessEntity)(uint,error)
	Delete(id uint)(error)
}
type ClassServiceInterface interface{
	Add(input ClassessEntity)(error)
	GetAll(page, pageSize int)([]ClassessEntity,error)
	Edit(id uint,input ClassessEntity)(error)
	GetById(id uint)(ClassessEntity,error)
	Delete(id uint)error
}