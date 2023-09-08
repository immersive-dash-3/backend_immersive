package service

import "immersive_project/klp3/features/dashboard"

type DashboardService struct {
	DashboardData dashboard.DashboardDataInterface
}

func New(repo dashboard.DashboardDataInterface) dashboard.DashboardServiceInterface {
	return &DashboardService{
		DashboardData: repo,
	}
}

// GetData implements dashboard.DashboardServiceInterface.
func (service *DashboardService) GetData() (dashboard.DashboardEntity, error) {
	result, err := service.DashboardData.GetData()
	return result, err
}
