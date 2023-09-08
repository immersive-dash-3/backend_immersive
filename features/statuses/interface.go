package statuses

import "time"

type StatusEntity struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StatusDataInterface interface {
	SelectAll() ([]StatusEntity, error)
}

type StatusServiceInterface interface {
	GetAll() ([]StatusEntity, error)
}
