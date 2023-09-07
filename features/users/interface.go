package users

import (
	"immersive_project/klp3/features/classes"
	"time"

	"github.com/labstack/echo/v4"
)

type UserEntity struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Name      string
	Email     string
	Team      string
	Password  string
	Role      string
	Address   string
	Status    string
	Classes   []classes.ClassessEntity
}

type QueryParams struct {
	Page            int
	ItemsPerPage    int
	SearchName      string
	IsUserDashboard bool
}

type UserDataInterface interface {
	Login(email, password string) (UserEntity, error)
	Create(user UserEntity) error
	Update(user UserEntity) error
	Delete(id uint) error
	FindById(id uint) (UserEntity, error)
	FindAll(qparams QueryParams) ([]UserEntity, int64, error)
}

type UserServiceInterface interface {
	Login(email, password string) (UserEntity, string, error)
	Create(user UserEntity) error
	Update(user UserEntity) error
	Delete(id uint) error
	FindById(id int) (UserEntity, error)
	FindAll(qparams QueryParams) ([]UserEntity, bool, error)
}

type UserHandlerInterface interface {
	UserLogin(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	FindById(c echo.Context) error
	FindAll(c echo.Context) error
}
