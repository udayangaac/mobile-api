package domain

type SignUpRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	DOB        string `json:"dob"`
}
