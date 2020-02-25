package entities

import "github.com/jinzhu/gorm"

type MobileUserConfiguration struct {
	gorm.Model
	UserId                 uint
	SoundStatus            int `gorm:"column:sound_status"`
	LocationServiceStatus  int `gorm:"column:location_service_status"`
	PushNotificationStatus int `gorm:"column:push_notification_status"`
	LoginStatus            int `gorm:"column:login_status"`
	AnyStatus              int `gorm:"column:any_status"`
	LastViewedAddId        int `gorm:"column:last_viewed_add_id"`
}
