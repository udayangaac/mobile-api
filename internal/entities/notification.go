package entities

import (
	"github.com/jinzhu/gorm"
)

type Notification struct {
	gorm.Model
	Content        string // `json:"content"`
	Status         int16  // `json:"status"`
	Address        string // json:"address"`
}