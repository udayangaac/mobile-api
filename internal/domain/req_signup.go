package domain

type SignUpRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Mobile     int    `json:"mobile"`
	Address    string `json:"Address"`
	DOB        string `json:"DOB"`
	JobStatus  int    `json:"job_status"`
	JobDetails struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	} `json:"job_details"`
	Married int `json:"married"`
	Family  struct {
		Kids int `json:"kids"`
	} `json:"family"`
	Location struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	} `json:"location"`
}
