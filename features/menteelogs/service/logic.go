package service

import "immersive_project/klp3/features/menteelogs"

type MenteeLogService struct {
	menteeLogService menteelogs.MenteeLogDataInterface
}

// GetAll implements menteelogs.MenteeLogServiceInterface.
func (service *MenteeLogService) GetAll() ([]menteelogs.MenteeLogEntity, error) {
	data,err:=service.menteeLogService.SelectAll()
	return data,err
}

// Delete implements menteelogs.MenteeLogServiceInterface.
func (service *MenteeLogService) Delete(idLog uint) error {
	err := service.menteeLogService.Delete(idLog)
	if err != nil {
		return err
	}
	return nil
}

// Edit implements menteelogs.MenteeLogServiceInterface.
func (service *MenteeLogService) Edit(idLog uint, input menteelogs.MenteeLogEntity) error {
	err := service.menteeLogService.Update(idLog, input)
	if err != nil {
		return err
	}
	return nil
}

// Get implements menteelogs.MenteeLogServiceInterface.
func (service *MenteeLogService) Get(idMentee uint) (menteelogs.MenteeEntity, error) {
	data, err := service.menteeLogService.Select(idMentee)
	if err != nil {
		return menteelogs.MenteeEntity{}, err
	}
	return data, nil
}

// Add implements menteelogs.MenteeLogServiceInterface.
func (service *MenteeLogService) Add(input menteelogs.MenteeLogEntity) error {
	id, errInsert := service.menteeLogService.Insert(input)
	if errInsert != nil {
		return errInsert
	}
	err := service.menteeLogService.UpdateMentee(id, input.MenteeID)
	if err != nil {
		return err
	}
	return nil
}

func New(mentee menteelogs.MenteeLogDataInterface) menteelogs.MenteeLogServiceInterface {
	return &MenteeLogService{
		menteeLogService: mentee,
	}
}
