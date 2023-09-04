package service

import "immersive_project/klp3/features/menteelogs"

type MenteeLogService struct {
	menteeLogService menteelogs.MenteeLogDataInterface
}


func New(mentee menteelogs.MenteeLogDataInterface) menteelogs.MenteeLogServiceInterface {
	return &MenteeLogService{
		menteeLogService: mentee,
	}
}
