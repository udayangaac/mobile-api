package entities

import (
	"github.com/jinzhu/gorm"
)

type AdvertisementsCategories struct {
	gorm.Model
	Id           int    // `json:"id"`
	Status       int    // `json:"status"`
	CategoryName string // `json:"name"`
}
