package classes

import (
	"time"
)

type ClassessEntity struct {
	Id        uint      
	CreatedAt time.Time 
	UpdatedAt time.Time 
	DeletedAt time.Time 
	Name      string    
	UserID    uint
	User UserEntity
}

type QueryParams struct {
	Page            int
	ItemsPerPage    int
	SearchName      string
	IsClassDashboard bool
}

type UserEntity struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Name      string
}

type ClassDataInterface interface {
	Insert(input ClassessEntity) (uint, error)
	SelectById(id uint) (ClassessEntity, error)
	SelectAll(input QueryParams) (int64, []ClassessEntity, error)
	Update(id uint, input ClassessEntity) (uint, error)
	Delete(id uint) error
}
type ClassServiceInterface interface {
	Add(input ClassessEntity) error
	GetAll(input QueryParams) (bool, []ClassessEntity, error)
	Edit(id uint, input ClassessEntity) error
	GetById(id uint) (ClassessEntity, error)
	Delete(id uint) error
}
