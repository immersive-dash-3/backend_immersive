package data

import (
	"errors"
	"immersive_project/klp3/features/menteelogs"

	"gorm.io/gorm"
)

type MenteeLogData struct {
	db *gorm.DB
}

// Insert implements menteelogs.MenteeLogDataInterface.
func (repo *MenteeLogData) Insert(input menteelogs.MenteeLogEntity) (string, error) {
	inputModel:=EntityToModel(input)
	tx:=repo.db.Create(&inputModel)
	if tx.Error != nil{
		return "",errors.New("failed insert log")
	}
	if tx.RowsAffected ==0{
		return "",errors.New("row not affected table log")
	}
	return inputModel.Status,nil
}

// InsertStatus implements menteelogs.MenteeLogDataInterface.
func (repo *MenteeLogData) InsertStatus(status string) (uint, error) {
	var inputModel Status
	inputModel.Name=status
	tx:=repo.db.Create(&inputModel)
	if tx.Error != nil{
		return 0,errors.New("failed insert status")
	}
	if tx.RowsAffected ==0{
		return 0,errors.New("row not affected table status")
	}
	return inputModel.ID,nil
}

// UpdateMentee implements menteelogs.MenteeLogDataInterface.
func (repo *MenteeLogData) UpdateMentee(idStatus uint,idMentee uint) error {
	var menteeInput Mentee
	menteeInput.StatusID=idStatus
	tx:=repo.db.Model(&Mentee{}).Where("id=?",idMentee).Updates(menteeInput)
	if tx.Error != nil{
		return errors.New("failed update status_id di tabel mentee")
	}
	if tx.RowsAffected ==0{
		return errors.New("row not affected table mentee")
	}
	return nil	
}

func New(db *gorm.DB) menteelogs.MenteeLogDataInterface {
	return &MenteeLogData{
		db: db,
	}
}
