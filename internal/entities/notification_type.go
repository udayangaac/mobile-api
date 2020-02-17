package entities

import "github.com/jinzhu/gorm"

type NotificationType struct {
	gorm.Model
	Id                         int     // `json:"id"`
	Status                     int     // `json:"status"`
	Name                       string  // `json:"image"`
}