package data

import (
	"errors"
	"immersive_project/klp3/features/statuses"

	"gorm.io/gorm"
)

type StatusQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) statuses.StatusDataInterface {
	return &StatusQuery{
		db: db,
	}
}

// SelectAll implements statuses.StatusDataInterface.
func (repo *StatusQuery) SelectAll() ([]statuses.StatusEntity, error) {
	var statusData []Status
	tx := repo.db.Where("deleted_at is null").Find(&statusData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("no row affected")
	}
	var statEntity []statuses.StatusEntity
	for _, v := range statusData {
		var status = ModelToEntity(v)
		statEntity = append(statEntity, status)
	}
	return statEntity, nil
}
