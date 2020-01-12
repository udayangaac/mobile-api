package entities

import "github.com/jinzhu/gorm"

type MobileAppUser struct {
	gorm.Model
	//Id             int    //`json:"id"`
	Name           string //`json:"name"`
	Email          string //`json:"email"`
	HashPassword   string //`json:"hash_password"`
	Gender         string //`json:"gender"`
	EmployeeStatus int16  //`json:"employee_status"`
	Status         int16  //`json:"status"`
}
