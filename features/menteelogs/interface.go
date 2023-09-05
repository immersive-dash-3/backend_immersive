package menteelogs

import (
	"immersive_project/klp3/features/classes"
	user "immersive_project/klp3/features/users"
	"time"
)

type MenteeLogEntity struct {
	Id        	uint      	`json:"id" form:"id"`
	CreatedAt 	time.Time 	`json:"created_at,omitempty"`
	UpdatedAt 	time.Time 	`json:"updated_at,omitempty"`
	DeletedAt 	time.Time 	`json:"deleted_at,omitempty"`
	MenteeID 	uint 		`json:"mentee_id" form:"mentee_id"`
	UserID 		uint 		`json:"user_id" form:"user_id"`
	Status      string      `json:"status" form:"status"`
	Log 		string 		`json:"log" form:"log"`
	Users       user.UserEntity `json:"users,omitempty"`
}

type MenteeEntity struct {
	Id        		uint      `json:"id" form:"id"`
	CreatedAt 		time.Time `json:"created_at,omitempty"`
	UpdatedAt 		time.Time `json:"updated_at,omitempty"`
	DeletedAt 		time.Time `json:"deleted_at,omitempty"`
	FirstName 		string `json:"first_name" form:"first_name"`
	LastName 		string `json:"last_name" form:"last_name"`
	Gender 			string `json:"gender" form:"gender"`
	Email 			string `json:"email" form:"email"`
	PhoneNumber 	string `json:"phone_number" form:"phone_number"`
	Telegram 		string `json:"telegram" form:"telegram"`
	Discord 		string  `json:"discord" form:"discord"`
	ClassID 		uint `json:"class_id" form:"class_id"`
	StatusID 		uint `json:"status_id" from:"status_id"`
	EducationType 	string `json:"education_type" form:"education_type"`
	CurrentAddress 	string `json:"current_address" form:"current_address"`
	HomeAddress 	string `json:"home_address" form:"home_address"`
	EmergencyName 	string `json:"emergency_name" form:"emergency_name"`
	EmergencyPhone 	string `json:"emergency_phone" form:"emergency_phone"`
	EmergencyStatus string `json:"emergency_status" form:"emergency_status"`
	Major 			string `json:"major" form:"major"`
	Graduate 		string `json:"graduate" form:"graduate"`
	Institution     string `json:"institution" form:"institution"`
	Class           classes.ClassessEntity `json:"class,omitempty"`
	Status          StatusEntity `json:"status,omitempty"`
	MenteeLog 		[]MenteeLogEntity `json:"logs,omitempty"`
}

type StatusEntity struct{
	Id        		uint      `json:"id" form:"id"`
	CreatedAt 		time.Time `json:"created_at,omitempty"`
	UpdatedAt 		time.Time `json:"updated_at,omitempty"`
	DeletedAt 		time.Time `json:"deleted_at,omitempty"`
	Name            string	  `json:"name" form:"name"`
}


type MenteeLogDataInterface interface{
	Insert(input MenteeLogEntity)(string,error)
	InsertStatus(status string)(uint,error)
	UpdateMentee(idStatus uint,idMentee uint)(error)
	Select(idMentee uint)(MenteeEntity,error)
	Update(idLog uint, input MenteeLogEntity)(error)
	Delete(idLog uint)(error)
}
type MenteeLogServiceInterface interface{
	Add(input MenteeLogEntity)error
	Get(idMentee uint)(MenteeEntity,error)
	Edit(idLog uint,input MenteeLogEntity)error
	Delete(idLog uint)(error)
}