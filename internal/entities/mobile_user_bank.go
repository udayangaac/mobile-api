package entities

import "github.com/jinzhu/gorm"

type MobileUserBank struct {
	gorm.Model
	MobileUserId string
	BankId       int
}
