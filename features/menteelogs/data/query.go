package data

import (
	"errors"
	"immersive_project/klp3/features/menteelogs"

	"gorm.io/gorm"
)

type MenteeLogData struct {
	db *gorm.DB
}

// SelectAll implements menteelogs.MenteeLogDataInterface.
func (repo *MenteeLogData) SelectAll() ([]menteelogs.MenteeLogEntity, error) {
	var inputModel []MenteeLog
	tx:=repo.db.Preload("Users").Find(&inputModel)
	if tx.Error != nil {
		return nil,errors.New("failed get all log")
	}	
	output:=ListModelToEntity(inputModel)
	return output,nil
}

// Delete implements menteelogs.MenteeLogDataInterface.
func (repo *MenteeLogData) Delete(idLog uint) error {
	var input MenteeLog
	tx := repo.db.Delete(&input, idLog)
	if tx.Error != nil {
		return errors.New("failed delete log")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected")
	}
	return nil
}

// Update implements menteelogs.MenteeLogDataInterface.
func (repo *MenteeLogData) Update(idLog uint, input menteelogs.MenteeLogEntity) error {
	inputLog := EntityToModel(input)
	tx := repo.db.Model(&MenteeLog{}).Where("id=?", idLog).Updates(inputLog)
	if tx.Error != nil {
		return errors.New("failed update log")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected")
	}
	return nil
}

// Select implements menteelogs.MenteeLogDataInterface.
func (repo *MenteeLogData) Select(idMentee uint) (menteelogs.MenteeEntity, error) {
	var input Mentee
	tx := repo.db.Preload("Status").Preload("Class").Preload("MenteeLogs").Preload("MenteeLogs.Users").First(&input, idMentee)
	if tx.Error != nil {
		return menteelogs.MenteeEntity{}, errors.New("failed read feedback mentee")
	}
	output := ModelMenteeToEntity(input)
	return output, nil
}

// Insert implements menteelogs.MenteeLogDataInterface.
func (repo *MenteeLogData) Insert(input menteelogs.MenteeLogEntity) (string, error) {
	inputModel := EntityToModel(input)
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return "", errors.New("failed insert log")
	}
	if tx.RowsAffected == 0 {
		return "", errors.New("row not affected table log")
	}
	return inputModel.Status, nil
}

// InsertStatus implements menteelogs.MenteeLogDataInterface.
func (repo *MenteeLogData) InsertStatus(status string) (uint, error) {
	var inputModel Status
	inputModel.Name = status
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return 0, errors.New("failed insert status")
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("row not affected table status")
	}
	return inputModel.ID, nil
}

// UpdateMentee implements menteelogs.MenteeLogDataInterface.
func (repo *MenteeLogData) UpdateMentee(idStatus uint, idMentee uint) error {
	var menteeInput Mentee
	menteeInput.StatusID = idStatus
	tx := repo.db.Model(&Mentee{}).Where("id=?", idMentee).Updates(menteeInput)
	if tx.Error != nil {
		return errors.New("failed update status_id di tabel mentee")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected table mentee")
	}
	return nil
}

func New(db *gorm.DB) menteelogs.MenteeLogDataInterface {
	return &MenteeLogData{
		db: db,
	}
}
