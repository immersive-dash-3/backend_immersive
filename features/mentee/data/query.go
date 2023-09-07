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
func (repo *MenteeQuery) SelectAll(limit, offset, status_id, class_id uint, education_type, search_name string) ([]mentee.MenteeEntity, int64, error) {
	var menteeModel []Mentee
	var count int64
	var tx *gorm.DB
	// fmt.Println(search_name)

	if status_id == 0 && class_id == 0 && education_type == "" {
		tx = repo.db.Preload("Status").Preload("Class").Where("(first_name like ? or last_name like ?) and deleted_at is null", "%"+search_name+"%", "%"+search_name+"%").Find(&menteeModel)
		count = tx.RowsAffected
		tx = repo.db.Preload("Status").Preload("Class").Limit(int(limit)).Offset(int(offset)).Where("(first_name like ? or last_name like ?) and deleted_at is null", "%"+search_name+"%", "%"+search_name+"%").Find(&menteeModel)
	} else if status_id != 0 && class_id == 0 && education_type == "" {
		tx = repo.db.Preload("Status").Preload("Class").Where("status_id = ? and (first_name like ? or last_name like ?) and deleted_at is null", status_id, "%"+search_name+"%", "%"+search_name+"%").Find(&menteeModel)
		count = tx.RowsAffected
		tx = repo.db.Preload("Status").Preload("Class").Limit(int(limit)).Offset(int(offset)).Where("status_id = ? and (first_name like ? or last_name like ?) and deleted_at is null", status_id, "%"+search_name+"%", "%"+search_name+"%").Find(&menteeModel)
	} else if status_id == 0 && class_id != 0 && education_type == "" {
		tx = repo.db.Preload("Status").Preload("Class").Where("class_id = ? and (first_name like ? or last_name like ?) and deleted_at is null", class_id, "%"+search_name+"%", "%"+search_name+"%").Find(&menteeModel)
		count = tx.RowsAffected
		tx = repo.db.Preload("Status").Preload("Class").Limit(int(limit)).Offset(int(offset)).Where("class_id = ? and (first_name like ? or last_name like ?) and deleted_at is null", class_id, "%"+search_name+"%", "%"+search_name+"%").Find(&menteeModel)
	} else if status_id != 0 && class_id != 0 && education_type == "" {
		tx = repo.db.Preload("Status").Preload("Class").Where("status_id = ? and class_id = ? and (first_name like ? or last_name like ?) and deleted_at is null", status_id, class_id, "%"+search_name+"%", "%"+search_name+"%").Find(&menteeModel)
		count = tx.RowsAffected
		tx = repo.db.Preload("Status").Preload("Class").Limit(int(limit)).Offset(int(offset)).Where("status_id = ? and class_id = ? and (first_name like ? or last_name like ?) and deleted_at is null", status_id, class_id, "%"+search_name+"%", "%"+search_name+"%").Find(&menteeModel)
	} else if status_id == 0 && class_id == 0 && education_type != "" {
		tx = repo.db.Preload("Status").Preload("Class").Where("education_type = ? and (first_name like ? or last_name like ?) and deleted_at is null", education_type, "%"+search_name+"%", "%"+search_name+"%").Find(&menteeModel)
		count = tx.RowsAffected
		tx = repo.db.Preload("Status").Preload("Class").Limit(int(limit)).Offset(int(offset)).Where("education_type = ? and (first_name like ? or last_name like ?) and deleted_at is null", education_type, "%"+search_name+"%", "%"+search_name+"%").Find(&menteeModel)
	} else if status_id != 0 && class_id == 0 && education_type != "" {
		tx = repo.db.Preload("Status").Preload("Class").Where("status_id = ? and education_type = ? and (first_name like ? or last_name like ?) and deleted_at is null", status_id, education_type, "%"+search_name+"%", "%"+search_name+"%").Find(&menteeModel)
		count = tx.RowsAffected
		tx = repo.db.Preload("Status").Preload("Class").Limit(int(limit)).Offset(int(offset)).Where("status_id = ? and education_type = ? and (first_name like ? or last_name like ?) and deleted_at is null", status_id, education_type, "%"+search_name+"%", "%"+search_name+"%").Find(&menteeModel)
	} else if status_id == 0 && class_id != 0 && education_type != "" {
		tx = repo.db.Preload("Status").Preload("Class").Where("class_id = ? and education_type = ? and (first_name like ? or last_name like ?) and deleted_at is null", class_id, education_type, "%"+search_name+"%", "%"+search_name+"%").Find(&menteeModel)
		count = tx.RowsAffected
		tx = repo.db.Preload("Status").Preload("Class").Limit(int(limit)).Offset(int(offset)).Where("class_id = ? and education_type = ? and (first_name like ? or last_name like ?) and deleted_at is null", class_id, education_type, "%"+search_name+"%", "%"+search_name+"%").Find(&menteeModel)
	} else if status_id != 0 && class_id != 0 && education_type != "" {
		tx = repo.db.Preload("Status").Preload("Class").Where("status_id = ? and class_id = ? and education_type = ? and (first_name like ? or last_name like ?) and deleted_at is null", status_id, class_id, education_type, "%"+search_name+"%", "%"+search_name+"%").Find(&menteeModel)
		count = tx.RowsAffected
		tx = repo.db.Preload("Status").Preload("Class").Limit(int(limit)).Offset(int(offset)).Where("status_id = ? and class_id = ? and education_type = ? and (first_name like ? or last_name like ?) and deleted_at is null", status_id, class_id, education_type, "%"+search_name+"%", "%"+search_name+"%").Find(&menteeModel)
	}

	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, 0, errors.New("no row affected")
	}

	var menteeData []mentee.MenteeEntity
	for _, v := range menteeModel {
		var data = ModelToEntity(v)
		menteeData = append(menteeData, data)
	}
	return menteeData, count, nil
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
