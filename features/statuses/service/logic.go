package service

import "immersive_project/klp3/features/statuses"

type StatusService struct {
	statusData statuses.StatusDataInterface
}

func New(repo statuses.StatusDataInterface) statuses.StatusServiceInterface {
	return &StatusService{
		statusData: repo,
	}
}

// GetAll implements statuses.StatusServiceInterface.
func (service *StatusService) GetAll() ([]statuses.StatusEntity, error) {
	result, err := service.statusData.SelectAll()
	return result, err
}
