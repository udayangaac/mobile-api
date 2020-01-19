package repositories

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/mobile-api/internal/entities"
	log_traceable "github.com/udayangaac/mobile-api/internal/lib/log-traceable"
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

func (m mobileAppUserMySqlRepo) AddMobileUser(ctx context.Context, mobileUser entities.MobileAppUser) (isUpdate bool) {
	fmt.Printf("in the repo")
	log.Info(log_traceable.GetMessage(ctx, fmt.Sprintf("%v", mobileUser)))
	isUpdate = m.DB.NewRecord(&mobileUser)
	log.Info(log_traceable.GetMessage(ctx, fmt.Sprintf("%v", isUpdate)))
	return
}

func (m mobileAppUserMySqlRepo) GetMobileUserByEmail(ctx context.Context, email string) (mobileUser entities.MobileAppUser, err error) {
	err = m.DB.Where("email=?", email).First(&mobileUser).Error
	return
}
