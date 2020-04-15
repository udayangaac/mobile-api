package entities

import (
	"github.com/jinzhu/gorm"
)

type AdvertisementsCategories struct {
	gorm.Model
	CategoryName string
}
