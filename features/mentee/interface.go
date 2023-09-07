package mentee

import (
	class "immersive_project/klp3/features/classes"
	"time"
)

type MenteeEntity struct {
	Id              uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
	FirstName       string
	LastName        string
	Gender          string
	Email           string
	PhoneNumber     string
	Telegram        string
	Discord         string
	ClassID         uint
	StatusID        uint
	EducationType   string
	CurrentAddress  string
	HomeAddress     string
	EmergencyName   string
	EmergencyPhone  string
	EmergencyStatus string
	Major           string
	Graduate        string
	Institution     string
	Class           class.ClassessEntity
	Status          StatusEntity
}

type StatusEntity struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Name      string
}

type MenteeDataInterface interface {
	SelectAll(page, item, status_id, class_id uint, education_type, search_name string) ([]MenteeEntity, int64, error)
	Insert(input MenteeEntity) error
	Select(mentee_id uint) (MenteeEntity, error)
	Update(mentee_id uint, input MenteeEntity) error
	Delete(mentee_id uint) error
}

type MenteeServiceInterface interface {
	GetAll(page, item, status_id, class_id uint, education_type, search_name string) ([]MenteeEntity, bool, error)
	Insert(input MenteeEntity) error
	Get(mentee_id uint) (MenteeEntity, error)
	Update(mentee_id uint, input MenteeEntity) error
	Delete(mentee_id uint) error
}
