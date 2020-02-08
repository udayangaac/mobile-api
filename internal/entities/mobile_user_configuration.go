package entities

import "github.com/jinzhu/gorm"

type MobileUserConfiguration struct {
	gorm.Model
	UserId int
	Status int `gorm:"column:sound_status"`
}
