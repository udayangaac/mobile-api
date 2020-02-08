package domain

type SettingChangeRequest struct {
	UserId    int  `json:"userId"`
	Status    int  `json:"status"`
}