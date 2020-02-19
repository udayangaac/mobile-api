package domain

import "github.com/jinzhu/gorm"

type NotificationTypes struct {
	gorm.Model
	// Id                   int     // `json:"id"`
	// Status               int     // `json:"status"`
	CategoryName         string     // `json:"name"`
}