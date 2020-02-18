package entities

import (
	"github.com/jinzhu/gorm"
)

type MobileAppUser struct {
	gorm.Model
	UserId         int    `json:"userId"`
	Name           string 	// `json:"name"`
	Email          string 	// `json:"email"`
	HashPassword   string 	// `json:"hash_password"`
	DOB            string 	// `json:"dob"`
	Gender         string 	// `json:"gender"`
	EmployeeStatus int16  	// `json:"employee_status"`
	Status         int16  	// `json:"status"`
	Address        string 	// `json:"address"`
	CivilStatus    int16  	// `json:"civil_status"`
	Family         Family   // `json:"family"`
	Job            JobDetails
}

type JobDetails struct {
	Name      string
	Address   string
}

type Family struct {
	Kids   int16
	KidsGender string
}