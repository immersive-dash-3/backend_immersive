package service

import (
	"fmt"
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
func (s *UserServiceImplementation) FindAll(page int, itemsPerPage int, searchName string) ([]users.UserEntity, bool, error) {
	nextPage := true
	fmt.Println(nextPage)
	res, total_users, err := s.data.FindAll(page, itemsPerPage, searchName)

	if err != nil {
		return []users.UserEntity{}, false, err
	}

	total_pages := total_users / int64(itemsPerPage)
	fmt.Println(total_pages, total_users)
	if total_users%int64(itemsPerPage) != 0 {
		total_pages += 1
	}

	if page == int(total_pages) {
		nextPage = false
	}

	fmt.Println(nextPage, "kedua")
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
