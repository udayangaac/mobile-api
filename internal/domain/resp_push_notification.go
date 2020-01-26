package domain

type PushResponse struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	ClientID        int    `json:"client_id"`
	Email           string `json:"email"`
	Avatar          string `json:"avatar"`
	LbsNotification struct {
		ID      int    `json:"id"`
		Content string `json:"content"`
	} `json:"lbs_notification"`
	Token string `json:"token"`
}