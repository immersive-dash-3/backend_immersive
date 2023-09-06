package users

import "time"

type UserEntity struct {
	Id        uint      `json:"id" form:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email" validate:"required, email"`
	Team      string    `json:"team" form:"team"`
	Password  string    `json:"password" form:"password" validate:"required"`
	Role      string    `json:"role" form:"role"`
	Address   string    `json:"address" form:"address"`
	Status    string    `json:"status" form:"status"`
}

type DataUserInterface interface {
}
