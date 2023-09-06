package data

import (
	"errors"
	"immersive_project/klp3/features/mentee"

	"gorm.io/gorm"
)

type MenteeQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.MenteeDataInterface {
	return &MenteeQuery{
		db: db,
	}
}

// Delete implements mentee.MenteeDataInterface.
func (repo *MenteeQuery) Delete(mentee_id uint) error {
	tx := repo.db.Where("id=?", mentee_id).Delete(&Mentee{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
}

// Insert implements mentee.MenteeDataInterface.
func (repo *MenteeQuery) Insert(input mentee.MenteeEntity) error {
	var inputModel = EntityToModel(input)
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no data inserted")
	}
	return nil
}

// Select implements mentee.MenteeDataInterface.
func (repo *MenteeQuery) Select(mentee_id uint) (mentee.MenteeEntity, error) {
	var menteeModel Mentee
	tx := repo.db.Where("id=?", mentee_id).Find(&menteeModel)
	if tx.Error != nil {
		return mentee.MenteeEntity{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return mentee.MenteeEntity{}, errors.New("no row affected")
	}
	var menteeData = ModelToEntity(menteeModel)
	return menteeData, nil
}

// SelectAll implements mentee.MenteeDataInterface.
func (repo *MenteeQuery) SelectAll(limit, offset uint) ([]mentee.MenteeEntity, error) {
	var menteeModel []Mentee
	tx := repo.db.Preload("Status", "deleted_at is null").Preload("Class", "deleted_at is null").Limit(int(limit)).Offset(int(offset)).Where("deleted_at is null").Find(&menteeModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("no row affected")
	}
	var menteeData []mentee.MenteeEntity
	for _, v := range menteeModel {
		var data = ModelToEntity(v)
		menteeData = append(menteeData, data)
	}
	return menteeData, nil
}

// Update implements mentee.MenteeDataInterface.
func (repo *MenteeQuery) Update(mentee_id uint, input mentee.MenteeEntity) error {
	var inputModel = EntityToModel(input)
	tx := repo.db.Model(&Mentee{}).Where("id=?", mentee_id).Updates(&inputModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
}
