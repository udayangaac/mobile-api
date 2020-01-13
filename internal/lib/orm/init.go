package orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/udayangaac/mobile-api/internal/config"
)

var DB *gorm.DB

func InitDatabase(dbConf config.DatabaseConfig) (err error) {
	connectionString := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		dbConf.UserName,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
		dbConf.Database,
	)
	DB, err = gorm.Open("mysql", connectionString)
	return
}

func CloseDatabase() (err error) {
	err = DB.Close()
	return
}
