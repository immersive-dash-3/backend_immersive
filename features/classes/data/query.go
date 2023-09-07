package data

import (
	"errors"
	"immersive_project/klp3/features/classes"

	"gorm.io/gorm"
)

type ClassData struct {
	db *gorm.DB
}

// Delete implements classes.ClassDataInterface.
func (repo *ClassData) Delete(id uint) error {
	var input Classes
	tx := repo.db.Delete(&input, id)
	if tx.Error != nil {
		return errors.New("failed deleted")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected")
	}
	return nil
}

// Update implements classes.ClassDataInterface.
func (repo *ClassData) Update(id uint, input classes.ClassessEntity) (uint, error) {
	inputModel := EntityToModel(input)
	tx := repo.db.Model(&Classes{}).Where("id=?", id).Updates(inputModel)
	if tx.Error != nil {
		return 0, errors.New("update failed")
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("row not affected")
	}
	return id, nil
}

// SelectAll implements classes.ClassDataInterface.
func (repo *ClassData) SelectAll(page, pageSize int) (int, []classes.ClassessEntity, error) {
	offset := (page - 1) * pageSize
	var input []Classes
	txx := repo.db.Find(&input)
	if txx.Error != nil {
		return 0, nil, errors.New("error get all data")
	}
	var count int = 0
	var dataInput []Classes
	for _, value := range input {
		count++
		dataInput = append(dataInput, value)
	}

	tx := repo.db.Offset(offset).Limit(pageSize).Find(&dataInput)
	if tx.Error != nil {
		return 0, nil, errors.New("error get all data per page")
	}
	output := ModelToEntityAll(dataInput)
	return count, output, nil
}

// Insert implements classes.ClassDataInterface.
func (repo *ClassData) Insert(input classes.ClassessEntity) (uint, error) {
	inputModel := EntityToModel(input)
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return 0, errors.New("failed insert data")
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("row not affected")
	}
	return inputModel.ID, nil
}

// SelectById implements classes.ClassDataInterface.
func (repo *ClassData) SelectById(id uint) (classes.ClassessEntity, error) {
	var inputModel Classes
	tx := repo.db.First(&inputModel, id)
	if tx.Error != nil {
		return classes.ClassessEntity{}, errors.New("failed get data by id")
	}
	output := ModelToEntity(inputModel)
	return output, nil
}

func New(db *gorm.DB) classes.ClassDataInterface {
	return &ClassData{
		db: db,
	}
}
