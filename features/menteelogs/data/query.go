package data

import (
	"immersive_project/klp3/features/menteelogs"

	"gorm.io/gorm"
)

type MenteeLogData struct {
	db *gorm.DB
}

func New(db *gorm.DB) menteelogs.MenteeLogDataInterface {
	return &MenteeLogData{
		db: db,
	}
}
