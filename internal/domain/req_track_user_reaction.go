package domain

type TrackUserReaction struct {
	UserId          int `json:"userId"`
	NotificationId  int `json:"notificationId"`
	Status          int `json:"status"`
}