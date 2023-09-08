package data

import (
	"immersive_project/klp3/features/statuses"

	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	Name string
}

func EntityToModel(status statuses.StatusEntity) Status {
	return Status{
		Name: status.Name,
	}
}

func ModelToEntity(status Status) statuses.StatusEntity {
	return statuses.StatusEntity{
		ID:   status.ID,
		Name: status.Name,
	}
}
