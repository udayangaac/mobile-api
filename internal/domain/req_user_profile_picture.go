package domain

type ProfilePictureRequest struct {
	UserId    string `json:"uid"`
	Image struct {
		Image string `json:"image"`
	} `json:"image"`
}