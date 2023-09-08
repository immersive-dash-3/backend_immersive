package data

import (
	"immersive_project/klp3/features/dashboard"
	menteeModel "immersive_project/klp3/features/mentee/data"
	menteeLogModel "immersive_project/klp3/features/menteelogs/data"

	"gorm.io/gorm"
)

type DashboardQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) dashboard.DashboardDataInterface {
	return &DashboardQuery{
		db: db,
	}
}

// GetData implements dashboard.DashboardDataInterface.
func (repo *DashboardQuery) GetData() (dashboard.DashboardEntity, error) {

	var active, placement, log int64

	tx := repo.db.Where("status_id in (?,?,?,?,?,?)", 3, 4, 5, 6, 7, 8).Find(&[]menteeModel.Mentee{})
	if tx.Error != nil {
		return dashboard.DashboardEntity{}, tx.Error
	}
	active = tx.RowsAffected
	tx2 := repo.db.Where("status_id = ?", 9).Find(&[]menteeModel.Mentee{})
	if tx2.Error != nil {
		return dashboard.DashboardEntity{}, tx2.Error
	}
	placement = tx2.RowsAffected
	tx3 := repo.db.Preload("Users").Find(&[]menteeLogModel.MenteeLog{})
	if tx3.Error != nil {
		return dashboard.DashboardEntity{}, tx3.Error
	}
	log = tx3.RowsAffected

	var dashEntity = dashboard.DashboardEntity{
		CountActive:    active,
		CountPlacement: placement,
		CountLog:       log,
	}
	return dashEntity, nil
}
