package users

import (
	"time"

	"github.com/labstack/echo/v4"
)

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

type UserDataInterface interface {
	Login(email, password string) (UserEntity, error)
	Create(user UserEntity) error
	Update(user UserEntity) error
	Delete(id uint) error
	FindById(id uint) (UserEntity, error)
	FindAll(page, itemsPerPage int, searchName string) ([]UserEntity, int64, error)
}

type UserServiceInterface interface {
	Login(email, password string) (UserEntity, string, error)
	Create(user UserEntity) error
	Update(user UserEntity) error
	Delete(id uint) error
	FindById(id int) (UserEntity, error)
	FindAll(page, itemsPerPage int, searchName string) ([]UserEntity, bool, bool, error)
}

type UserHandlerInterface interface {
	UserLogin(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	FindById(c echo.Context) error
	FindAll(c echo.Context) error
}
