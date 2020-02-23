package entities

import (
	"github.com/jinzhu/gorm"
)

type UserAdvertisementCategories struct {
	gorm.Model
	UserId             int
	AdvertisementCatId int
}
