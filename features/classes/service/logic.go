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
	err := service.classService.Delete(id)
	if err != nil {
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
func (service *ClassService) Edit(id uint, input classes.ClassessEntity) error {
	_, err := service.classService.Update(id, input)
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements classes.ClassServiceInterface.
func (service *ClassService) GetAll(input classes.QueryParams) (bool, []classes.ClassessEntity, error) {
	var total_pages int64
	nextPage := true
	count, data, err := service.classService.SelectAll(input)
	if err != nil {
		return true, nil, err
	}
	if input.IsClassDashboard{
		total_pages = count /int64(input.ItemsPerPage)
		if count % int64(input.ItemsPerPage) != 0{
			total_pages +=1
		}

		if input.Page == int(total_pages){
			nextPage = false
		}
	}

	return nextPage, data, nil
}

// Add implements classes.ClassServiceInterface.
func (service *ClassService) Add(input classes.ClassessEntity) error {
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
