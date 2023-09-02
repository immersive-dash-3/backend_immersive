package mentee

import (
	class "immersive_project/klp3/features/classes"
	"time"
)

type MenteeEntity struct {
	Id        		uint      `json:"id" form:"id"`
	CreatedAt 		time.Time `json:"created_at,omitempty"`
	UpdatedAt 		time.Time `json:"updated_at,omitempty"`
	DeletedAt 		time.Time `json:"deleted_at,omitempty"`
	FullName 		string `json:"full_name" form:"full_name"`
	NickName 		string `json:"nick_name" form:"nick_name"`
	Gender 			string `json:"gender" form:"gender"`
	Email 			string `json:"email" form:"email"`
	PhoneNumber 	string `json:"phone_number" form:"phone_number"`
	Telegram 		string `json:"telegram" form:"telegram"`
	ClassID 		uint `json:"class_id" form:"class_id"`
	Status 			string `json:"status" from:"status"`
	Category 		string `json:"category" form:"category"`
	CurrentAddress 	string `json:"current_address" form:"current_address"`
	HomeAddress 	string `json:"home_address" form:"home_address"`
	EmergencyName 	string `json:"emergency_name" form:"emergency_name"`
	EmergencyPhone 	string `json:"emergency_phone" form:"emergency_phone"`
	EmergencyStatus string `json:"emergency_status" form:"emergency_status"`
	Major 			string `json:"major" form:"major"`
	Graduate 		string `json:"graduate" form:"graduate"`
	Class           class.ClassessEntity `json:"class,omitempty"`
}
