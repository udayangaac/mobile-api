package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type AdvertismentsCategories struct {
	gorm.Model
    Id                   int     	// `json:"id"`
    Status               int     	// `json:"status"`
	CategoryName         string  	// `json:"name"`
	CreatedAt            time.Time  `json:"-"`
	UpdatedAt            time.Time  `json:"-"`
	DeletedAt            *time.Time `json:"-";sql:"index"`
}