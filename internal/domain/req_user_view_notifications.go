package domain

type UserViewedNotification struct {
	UserId          int `json:"userId"`
	NotificationId  int `json:"notificationId"`
}