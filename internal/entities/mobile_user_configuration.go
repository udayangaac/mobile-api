package entities

import "github.com/jinzhu/gorm"

type MobileUserConfiguration struct {
	gorm.Model
	UserId                int
	SoundStatus           int `gorm:"column:sound_status"`
	LocationServiceStatus int `gorm:"column:location_service_status"`
}
