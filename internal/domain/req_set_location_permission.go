package domain

type LocationPermissionRequest struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}
