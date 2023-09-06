package data

import (
	"errors"
	"immersive_project/klp3/exception"
	"immersive_project/klp3/features/users"
	"immersive_project/klp3/helper"

	"gorm.io/gorm"
)

type UserDataImplementation struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserDataImplementation{db: db}
}

// Create implements users.UserDataInterface
func (data *UserDataImplementation) Create(user users.UserEntity) error {
	userGorm := EntityToModel(user)
	tx := data.db.Create(&userGorm)

	if tx.Error != nil {
		return exception.ErrInternalServer
	}
	return nil
}

// Delete implements users.UserDataInterface
func (data *UserDataImplementation) Delete(id uint) error {
	tx := data.db.Delete(&User{}, id)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return exception.ErrIdIsNotFound
		}
	} else {
		return exception.ErrInternalServer
	}

	return nil
}

// FindAll implements users.UserDataInterface
func (data *UserDataImplementation) FindAll(page, itemsPerPage int) ([]users.UserEntity, int64, error) {
	var userEntities []users.UserEntity
	var userDatas []User

	tx := data.db.Find(&userDatas)
	if tx.Error != nil {
		return []users.UserEntity{}, 0, exception.ErrInternalServer
	}

	for _, data := range userDatas {
		entity := ModelToEntity(data)
		userEntities = append(userEntities, entity)
	}

	count_users := tx.RowsAffected

	return userEntities, count_users, nil

}

// FindById implements users.UserDataInterface
func (data *UserDataImplementation) FindById(id uint) (users.UserEntity, error) {
	var user User
	var userEntity users.UserEntity

	tx := data.db.Find(&user, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return users.UserEntity{}, exception.ErrIdIsNotFound
		}
	} else {
		return users.UserEntity{}, exception.ErrInternalServer
	}

	userEntity = ModelToEntity(user)
	return userEntity, nil
}

// Login implements users.UserDataInterface
func (data *UserDataImplementation) Login(email string, password string) (users.UserEntity, error) {
	var user User
	hashPassword, err := helper.HassPassword(password)
	if err != nil {
		return users.UserEntity{}, exception.ErrInternalServer
	}

	tx := data.db.Where("email = ? and password = ?", email, hashPassword).Find(&user)

	if tx.RowsAffected == 0 {
		return users.UserEntity{}, gorm.ErrRecordNotFound
	}

	if tx.Error != nil {
		return users.UserEntity{}, exception.ErrInternalServer
	}

	dataLogin := ModelToEntity(user)

	return dataLogin, nil

}

// Update implements users.UserDataInterface
func (data *UserDataImplementation) Update(user users.UserEntity) error {
	var userGorm User

	tx := data.db.Model(&userGorm).Updates(user)
	if tx.Error != nil {
		return exception.ErrInternalServer
	}

	if tx.RowsAffected == 0 {
		return exception.ErrIdIsNotFound
	}

	return nil

}
