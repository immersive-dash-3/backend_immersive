package service

import (
	"immersive_project/klp3/features/mentee"
)

type MenteeService struct {
	menteeData mentee.MenteeDataInterface
}

func New(repo mentee.MenteeDataInterface) mentee.MenteeServiceInterface {
	return &MenteeService{
		menteeData: repo,
	}
}

// Delete implements mentee.MenteeServiceInterface.
func (service *MenteeService) Delete(mentee_id uint) error {
	err := service.menteeData.Delete(mentee_id)
	return err
}

// Get implements mentee.MenteeServiceInterface.
func (service *MenteeService) Get(mentee_id uint) (mentee.MenteeEntity, error) {
	result, err := service.menteeData.Select(mentee_id)
	return result, err
}

// GetAll implements mentee.MenteeServiceInterface.
func (service *MenteeService) GetAll(page, item, status_id, class_id uint, education_type, search_name string) ([]mentee.MenteeEntity, bool, error) {
	result, count, err := service.menteeData.SelectAll(page, item, status_id, class_id, education_type, search_name)

	next := true
	var pages int64
	if item != 0 {
		pages = count / int64(item)
		if count%int64(item) != 0 {
			pages += 1
		}
		if page == uint(pages) {
			next = false
		}
	}

	return result, next, err
}

// Insert implements mentee.MenteeServiceInterface.
func (service *MenteeService) Insert(input mentee.MenteeEntity) error {
	err := service.menteeData.Insert(input)
	return err
}

// Update implements mentee.MenteeServiceInterface.
func (service *MenteeService) Update(mentee_id uint, input mentee.MenteeEntity) error {
	err := service.menteeData.Update(mentee_id, input)
	return err
}
