package service

import "immersive_project/klp3/features/mentee"

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
func (service *MenteeService) GetAll(page, item uint) ([]mentee.MenteeEntity, error) {
	limit := item
	offset := (page - 1) * item
	result, err := service.menteeData.SelectAll(limit, offset)
	return result, err
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
