package domain

import "github.com/udayangaac/mobile-api/internal/entities"

type UserProfileResponse struct {
	UserId uint   `json:"userId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	// Mobile     					  int    `json:"mobile"`
	Address            string `json:"address"`
	DOB                string `json:"dob"`
	Gender             string `json:"gender"`
	EmployeeStatus     int16  `json:"employee_status"`
	JobDetails      struct {
		JobCompanyName     string `json:"name"`
		JobCompanyLocation string `json:"location"`
	} `json:"job_details"`
	CivilStatus        int16  `json:"civil_status"`
	Kids               int    `json:"kids"`
	Configuration      struct {
		LoginStatus            int `json:"login_status"`
		PushNotificationStatus int `json:"push_notification_status"`
		SoundStatus            int `json:"sound_status"`
		LocationServiceStatus  int `json:"location_status"`
		AnyStatus              int `json:"any_status"`
	} `json:"configuration"`
	UserAdvertisementsCategories []entities.AdvertisementsCategories `json:"advertisement_cat_id"`
	UserBanks []entities.Banks `json:"bank_id_list"`
}
