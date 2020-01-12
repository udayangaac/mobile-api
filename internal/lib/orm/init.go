package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/udayangaac/mobile-api/internal/config"
)

var DB *gorm.DB

func InitDatabase(dbConf config.DatabaseConfig) (err error) {
	DB, err = gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	return
}

func CloseDatabase() (err error) {
	err = DB.Close()
	return
}
