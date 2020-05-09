package entities

import "github.com/jinzhu/gorm"

type MobileUserViewedAdvertisementList struct {
	gorm.Model
	UserId          	int
	NotificationId      int
}