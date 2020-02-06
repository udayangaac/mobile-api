package domain

type SoundPermissionRequest struct {
	UserId    int `json:"userId"`
	Status    int `json:"status"`
}