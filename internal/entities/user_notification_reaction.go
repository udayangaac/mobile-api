package entities

import "github.com/jinzhu/gorm"

type MobileUserResponse struct {
	gorm.Model
	UserId          	int
	Status				int
	NotificationId      int
}
