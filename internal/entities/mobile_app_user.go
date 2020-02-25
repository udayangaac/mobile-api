package entities

import (
	"github.com/jinzhu/gorm"
)

type MobileAppUser struct {
	gorm.Model
	Name                        string // `json:"name"`
	Email                       string // `json:"email"`
	HashPassword                string // `json:"hash_password"`
	DOB                         string // `json:"dob"`
	Gender                      string // `json:"gender"`
	EmployeeStatus              int16  // `json:"employee_status"`
	Status                      int16  // `json:"status"`
	Address                     string // `json:"address"`
	CivilStatus                 int16  // `json:"civil_status"`
	JobCompanyName              string //`json:"job_company_name"`
	JobCompanyLocation          string
	Kids                        int    //`json:"kids"`
	LoginStatus            		int
	PushNotificationStatus 		int
	SoundStatus            		int
	LocationServiceStatus  		int
	AnyStatus              		int
	MobileUserConfigurations    MobileUserConfiguration
	// UserAdvertisementCategories []UserAdvertisementCategories
}

type UserConfiguration struct {
	UserId                 uint
	LoginStatus            int
	PushNotificationStatus int
	SoundStatus            int
	LocationServiceStatus  int
	AnyStatus              int
	LastViewedAddId         int
}
