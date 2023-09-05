package service

import (
	"errors"
	"immersive_project/klp3/features/classes"
	handlers "immersive_project/klp3/features/classes/handler"

	"github.com/go-playground/validator/v10"
)

type ClassService struct {
	classService classes.ClassDataInterface
	validate     *validator.Validate
}

// Delete implements classes.ClassServiceInterface.
func (service *ClassService) Delete(id uint) error {
	err:=service.classService.Delete(id)
	if err != nil{
		return err
	}
	return nil
}

// GetById implements classes.ClassServiceInterface.
func (service *ClassService) GetById(id uint) (classes.ClassessEntity, error) {
	data, err := service.classService.SelectById(id)
	if err != nil {
		return classes.ClassessEntity{}, err
	}
	return data, nil
}

// Edit implements classes.ClassServiceInterface.
func (service *ClassService) Edit(id uint, input classes.ClassessEntity) (error) {
	_,err := service.classService.Update(id, input)
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements classes.ClassServiceInterface.
func (service *ClassService) GetAll() ([]classes.ClassessEntity, error) {
	data, err := service.classService.SelectAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Add implements classes.ClassServiceInterface.
func (service *ClassService) Add(input classes.ClassessEntity) (error) {
	inputValidate := handlers.EntityToRequest(input)
	errValidate := service.validate.Struct(&inputValidate)
	if errValidate != nil {
		return errors.New("nama class harus diisi")
	}
	inputEntity := handlers.RequestToEntity(inputValidate)
	_, errInsert := service.classService.Insert(inputEntity)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func New(class classes.ClassDataInterface) classes.ClassServiceInterface {
	return &ClassService{
		classService: class,
		validate:     validator.New(),
	}
}
