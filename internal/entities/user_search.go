package entities

import "github.com/jinzhu/gorm"

type MobileUserSearch struct {
	gorm.Model
	UserId          int
	SearchText		string
}
