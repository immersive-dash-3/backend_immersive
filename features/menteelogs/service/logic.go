package service

import "immersive_project/klp3/features/menteelogs"

type MenteeLogService struct {
	menteeLogService menteelogs.MenteeLogDataInterface
}

// Get implements menteelogs.MenteeLogServiceInterface.
func (service *MenteeLogService) Get(idMentee uint) (menteelogs.MenteeEntity, error) {
	data,err:=service.menteeLogService.Select(idMentee)
	if err!= nil{
		return menteelogs.MenteeEntity{},err
	}
	return data,nil
}

// Add implements menteelogs.MenteeLogServiceInterface.
func (service *MenteeLogService) Add(input menteelogs.MenteeLogEntity) error {
	status, errInsert := service.menteeLogService.Insert(input)
	if errInsert != nil {
		return errInsert
	}
	id, errInstStat := service.menteeLogService.InsertStatus(status)
	if errInstStat != nil {
		return errInstStat
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
