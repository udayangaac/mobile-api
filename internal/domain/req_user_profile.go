package domain

type UserProfile struct {
	UserId     int    `json:"userId"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Mobile     int    `json:"mobile"`
	Address    string `json:"Address"`
	DOB        string `json:"DOB"`
	Gender     string `json:"Gender"`
	JobStatus  int16  `json:"job_status"`
	JobDetails struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"job_details"`
	Married int16 `json:"married"`
	Family  struct {
		Kids int `json:"kids"`
	} `json:"family"`
	Location struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	} `json:"location"`
}
