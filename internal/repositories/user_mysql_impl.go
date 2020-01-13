package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/udayangaac/mobile-api/internal/entities"
)

type mobileAppUserMySqlRepo struct {
	DB *gorm.DB
}

func NewMobileAppUser() MobileAppUserRepo {
	return &mobileAppUserMySqlRepo{}
}

func (m mobileAppUserMySqlRepo) AddMobileUser(mobileUser entities.MobileAppUser) (isUpdate bool) {
	return m.DB.NewRecord(&mobileUser)
}

func (m mobileAppUserMySqlRepo) GetMobileUserByEmail(email string) (mobileUser entities.MobileAppUser, err error) {
	err = m.DB.Where("email=?", email).First(&mobileUser).Error
	return
}
