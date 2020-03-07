package entities

import (
	"github.com/jinzhu/gorm"
)

type Banks struct {
	gorm.Model
	// UserId          int
	Name		string
}
