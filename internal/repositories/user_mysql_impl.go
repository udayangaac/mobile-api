package repositories

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/mobile-api/internal/entities"
	"github.com/udayangaac/mobile-api/internal/errors_custom"
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

func (m mobileAppUserMySqlRepo) AddMobileUser(ctx context.Context, mobileUser entities.MobileAppUser, mobileUserConfiguration entities.MobileUserConfiguration) (err error) {
	log.Info(log_traceable.GetMessage(ctx, fmt.Sprintf("%v", mobileUser)))
	rowAffected := m.DB.Create(&mobileUser).RowsAffected
	if rowAffected == 0 {
		err = errors_custom.ErrDuplicateUserEntry
		return
	}
	mobileUserConfiguration.UserId = mobileUser.ID
	err = m.DB.Create(&mobileUserConfiguration).Error
	return
}

func (m mobileAppUserMySqlRepo) GetMobileUserByEmail(ctx context.Context, email string) (mobileUser entities.MobileAppUser, err error) {
	err = m.DB.Where("email=?", email).First(&mobileUser).Error
	return
}

func (m mobileAppUserMySqlRepo) UserLogout(ctx context.Context) (err error) {
	return
}

func (m mobileAppUserMySqlRepo) PushNotification(ctx context.Context, userId int, lat float64, lon float64) (PushNotification entities.Notification, err error) {
	err = m.DB.Where("id=?", userId).First(&PushNotification).Error
	return
}

func (m mobileAppUserMySqlRepo) PullNotification(ctx context.Context, userId int, lat float64, lon float64) (Notification entities.Notification, err error) {
	err = m.DB.Where("id=?", userId).First(&Notification).Error
	return

}

func (m mobileAppUserMySqlRepo) LocationTrack(ctx context.Context, userId int, status int) (err error) {
	err = m.DB.Model(entities.MobileUserConfiguration{}).Where("user_id = 1").Update("location_service_status", "0").Error
	return

}

func (m mobileAppUserMySqlRepo) UserProfilePicture(ctx context.Context, userId int16) (mobileUser entities.MobileAppUser, err error) {
	err = m.DB.Where("email=?").First(&mobileUser).Error
	return

}

func (m mobileAppUserMySqlRepo) SoundSettingChange(ctx context.Context, userId int, status int) (err error) {
	err = m.DB.Model(entities.MobileUserConfiguration{}).Where("user_id = ?", userId).Update("sound_status", status).Error
	log.Info(err)
	return
}

func (m mobileAppUserMySqlRepo) PushNotificationSetting(ctx context.Context, userId int, status int) (err error) {
	err = m.DB.Model(entities.MobileUserConfiguration{}).Where("user_id = ?", userId).Update("push_notification_status", status).Error
	log.Info(err)
	return
}

func (m mobileAppUserMySqlRepo) SetLoginStatus(ctx context.Context, userId int, status int) (err error) {
	err = m.DB.Model(entities.MobileUserConfiguration{}).Where("user_id = ?", userId).Update("login_status", status).Error
	log.Info(err)
	return
}
