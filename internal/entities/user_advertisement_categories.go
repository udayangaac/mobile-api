package entities

import (
	"github.com/jinzhu/gorm"
)

type UserAdvertisementCategories struct {
	gorm.Model
	UserId                       uint
	AdvertisementCatId           int
}