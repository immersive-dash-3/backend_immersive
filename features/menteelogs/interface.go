package menteelogs

import (
	"immersive_project/klp3/features/classes"
	user "immersive_project/klp3/features/users"
	"time"
)

type MenteeLogEntity struct {
	Id        	uint      	
	CreatedAt 	time.Time 	
	UpdatedAt 	time.Time 	
	DeletedAt 	time.Time 	
	MenteeID 	uint 		
	UserID 		uint 		
	StatusID      uint      
	Log 		string 		
	Users       user.UserEntity 
	Status      StatusEntity
}

type MenteeEntity struct {
	Id        		uint      
	CreatedAt 		time.Time 
	UpdatedAt 		time.Time 
	DeletedAt 		time.Time 
	FirstName 		string 
	LastName 		string 
	Gender 			string 
	Email 			string 
	PhoneNumber 	string 
	Telegram 		string 
	Discord 		string  
	ClassID 		uint 
	StatusID 		uint 
	EducationType 	string 
	CurrentAddress 	string 
	HomeAddress 	string 
	EmergencyName 	string 
	EmergencyPhone 	string 
	EmergencyStatus string 
	Major 			string 
	Graduate 		string 
	Institution     string 
	Class           classes.ClassessEntity 
	Status          StatusEntity 
	MenteeLog 		[]MenteeLogEntity 
}

type StatusEntity struct{
	Id        		uint      
	CreatedAt 		time.Time 
	UpdatedAt 		time.Time 
	DeletedAt 		time.Time 
	Name            string	  
}


type MenteeLogDataInterface interface{
	Insert(input MenteeLogEntity)(uint,error)
	InsertStatus(status string)(uint,error)
	UpdateMentee(idStatus uint,idMentee uint)(error)
	Select(idMentee uint)(MenteeEntity,error)
	Update(idLog uint, input MenteeLogEntity)(error)
	Delete(idLog uint)(error)
	SelectAll()([]MenteeLogEntity,error)
}
type MenteeLogServiceInterface interface{
	Add(input MenteeLogEntity)error
	Get(idMentee uint)(MenteeEntity,error)
	Edit(idLog uint,input MenteeLogEntity)error
	Delete(idLog uint)(error)
	GetAll()([]MenteeLogEntity,error)
}