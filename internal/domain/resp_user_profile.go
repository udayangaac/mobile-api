package domain

type UserProfileResponse struct {
	UserId     					  int    `json:"userId"`
	Name            			  string `json:"name"`
	Email           			  string `json:"email"`
	// Mobile     					  int    `json:"mobile"`
	Address    					  string `json:"address"`
	DOB        					  string `json:"dob"`
	Gender     					  string `json:"gender"`
	EmployeeStatus       		  int16  `json:"employee_status"`
	CivilStatus 				  int16  `json:"civil_status"`
	JobCompanyName                string `json:"job_company_name"`
	JobCompanyLocation            string `json:"job_company_location"`
	Configuration struct {
		LoginStatus               int `json:"login_status"`
		PushNotificationStatus    int `json:"push_notification_status"`
		SoundStatus				  int `json:"sound_status"`
		LocationServiceStatus     int `json:"location_status"`
		AnyStatus				  int `json:"any_status"`
	} `json:"configuration"`
	UserAdvertisementsCategories  []UserAdvertisementsCategories
}

type UserAdvertisementsCategories struct {
	Id                         int `json:"id"`
	AdvertisementCategoryId    int `json:"advertisement_catId"`
	CategoryName               string `json:"category_name"`
}

type Family  struct {
	Kids int `json:"kids"`
}