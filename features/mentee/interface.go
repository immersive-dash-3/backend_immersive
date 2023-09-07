package mentee

import (
	class "immersive_project/klp3/features/classes"
	"time"
)

type MenteeEntity struct {
	Id              uint                 `json:"id" form:"id"`
	CreatedAt       time.Time            `json:"created_at,omitempty"`
	UpdatedAt       time.Time            `json:"updated_at,omitempty"`
	DeletedAt       time.Time            `json:"deleted_at,omitempty"`
	FirstName       string               `json:"first_name" form:"first_name"`
	LastName        string               `json:"last_name" form:"last_name"`
	Gender          string               `json:"gender" form:"gender"`
	Email           string               `json:"email" form:"email"`
	PhoneNumber     string               `json:"phone_number" form:"phone_number"`
	Telegram        string               `json:"telegram" form:"telegram"`
	Discord         string               `json:"discord" form:"discord"`
	ClassID         uint                 `json:"class_id" form:"class_id"`
	StatusID        uint                 `json:"status_id" from:"status_id"`
	EducationType   string               `json:"education_type" form:"education_type"`
	CurrentAddress  string               `json:"current_address" form:"current_address"`
	HomeAddress     string               `json:"home_address" form:"home_address"`
	EmergencyName   string               `json:"emergency_name" form:"emergency_name"`
	EmergencyPhone  string               `json:"emergency_phone" form:"emergency_phone"`
	EmergencyStatus string               `json:"emergency_status" form:"emergency_status"`
	Major           string               `json:"major" form:"major"`
	Graduate        string               `json:"graduate" form:"graduate"`
	Institution     string               `json:"institution" form:"institution"`
	Class           class.ClassessEntity `json:"class,omitempty"`
	Status          StatusEntity         `json:"status,omitempty"`
}

type StatusEntity struct {
	Id        uint      `json:"id" form:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Name      string    `json:"name" form:"name"`
}

type MenteeDataInterface interface {
	SelectAll(limit, offset, status_id, class_id uint, education_type, search_name string) ([]MenteeEntity, int64, error)
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
