package entities

import "github.com/jinzhu/gorm"

type MobileUserConfiguration struct {
	gorm.Model
	UserId         int  // `json:"userId"`
	Status         int  // `json:"status"`
}