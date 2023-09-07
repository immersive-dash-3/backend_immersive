package service

import (
	"immersive_project/klp3/app/middleware"
	"immersive_project/klp3/exception"
	"immersive_project/klp3/features/users"
)

type UserServiceImplementation struct {
	data users.UserDataInterface
}

func New(data users.UserDataInterface) users.UserServiceInterface {
	return &UserServiceImplementation{data: data}
}

// Create implements users.UserServiceInterface
func (s *UserServiceImplementation) Create(user users.UserEntity) error {
	err := s.data.Create(user)
	return err
}

// Delete implements users.UserServiceInterface
func (s *UserServiceImplementation) Delete(id uint) error {
	err := s.data.Delete(id)
	return err
}

// FindAll implements users.UserServiceInterface
func (s *UserServiceImplementation) FindAll(qparams users.QueryParams) ([]users.UserEntity, bool, error) {
	var total_pages int64
	nextPage := true

	res, total_users, err := s.data.FindAll(qparams)

	if err != nil {
		return []users.UserEntity{}, false, err
	}

	if qparams.IsUserDashboard {
		total_pages = total_users / int64(qparams.ItemsPerPage)

		if total_users%int64(qparams.ItemsPerPage) != 0 {
			total_pages += 1
		}

		if qparams.Page == int(total_pages) {
			nextPage = false
		}
	}

	return res, nextPage, nil

}

// FindById implements users.UserServiceInterface
func (s *UserServiceImplementation) FindById(id int) (users.UserEntity, error) {
	res, err := s.data.FindById(uint(id))
	return res, err
}

// Login implements users.UserServiceInterface
func (s *UserServiceImplementation) Login(email string, password string) (users.UserEntity, string, error) {
	user, err := s.data.Login(email, password)
	if err != nil {
		return users.UserEntity{}, "", err
	}

	token, err := middleware.CreateToken(int(user.Id), user.Role)
	if err != nil {
		return users.UserEntity{}, "", exception.ErrInternalServer
	}

	return user, token, nil

}

// Update implements users.UserServiceInterface
func (s *UserServiceImplementation) Update(user users.UserEntity) error {
	err := s.data.Update(user)
	return err
}
