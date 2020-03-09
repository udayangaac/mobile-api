package entities

import "github.com/jinzhu/gorm"

type Banks struct {
	gorm.Model
	//Id          int
	Name		string
}
