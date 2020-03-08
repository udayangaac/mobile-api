package entities

import "github.com/jinzhu/gorm"

type MobileUserBank struct {
	gorm.Model
	MobileUserId int
	BankId       int
}
