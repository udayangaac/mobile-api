package entities

import "github.com/jinzhu/gorm"

type MobileUserPushNotificationsResponse struct {
	gorm.Model
	UserId          	int
	Status				int
	NotificationId      int
}
