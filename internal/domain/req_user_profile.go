package domain

type UserProfile struct {
	UserId   int    `json:"userId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// Mobile     int    `json:"mobile"`
	Address    string `json:"address"`
	DOB        string `json:"dob"`
	Gender     string `json:"gender"`
	JobStatus  int16  `json:"job_status"`
	JobDetails struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"job_details"`
	Configuration struct {
		LoginStatus            int `json:"login_status"`
		PushNotificationStatus int `json:"push_notification_status"`
		SoundStatus            int `json:"sound_status"`
		LocationServiceStatus  int `json:"location_status"`
		AnyStatus              int `json:"any_status"`
	} `json:"configuration"`
	UserAdvertisementsCategories struct {
		AdvertisementCatId int `json:"advertisement_cat_id"`
	} `json:"user_advertisement_categories"`
	CivilStatus int16 `json:"civil_status"`
	Kids        int   `json:"kids"`
}
