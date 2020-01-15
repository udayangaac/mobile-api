package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/udayangaac/mobile-api/internal/entities"
	"github.com/udayangaac/mobile-api/internal/lib/orm"
)

type mobileAppUserMySqlRepo struct {
	DB *gorm.DB
}

func NewMobileAppUser() MobileAppUserRepo {
	return &mobileAppUserMySqlRepo{
		DB: orm.DB,
	}
}

func (m mobileAppUserMySqlRepo) AddMobileUser(mobileUser entities.MobileAppUser) (isUpdate bool) {
	return m.DB.NewRecord(&mobileUser)
}

func (m mobileAppUserMySqlRepo) GetMobileUserByEmail(email string) (mobileUser entities.MobileAppUser, err error) {
	err = m.DB.Where("email=?", email).First(&mobileUser).Error
	return
}
