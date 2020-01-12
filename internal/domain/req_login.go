package domain

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Location struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	} `json:"location"`
}
