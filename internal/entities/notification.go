package entities

import (
	"github.com/jinzhu/gorm"
)

type Notification struct {
	gorm.Model
	NotificationContent        string  // `json:"content"`
	Status                     int     // `json:"status"`
	Image                      string  // `json:"image"`
	Company                    int     // `json:"company"`
}