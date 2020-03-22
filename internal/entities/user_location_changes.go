package entities

import (
	"github.com/jinzhu/gorm"
)

type UserLocationChanges struct {
	gorm.Model
	UserId          int
	Lat				string
	Lon             string
	//Date            string
	CellId          int
}
