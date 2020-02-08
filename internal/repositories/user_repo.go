package repositories

import (
	"context"
	"github.com/udayangaac/mobile-api/internal/entities"
)

type MobileAppUserRepo interface {
	AddMobileUser(ctx context.Context, mobileUser entities.MobileAppUser)
	GetMobileUserByEmail(ctx context.Context, email string) (mobileUser entities.MobileAppUser, err error)
	UserLogout(ctx context.Context) (err error)
	PushNotification(ctx context.Context, userId int, lat float64, lon float64) (notification entities.Notification, err error)
	PullNotification(ctx context.Context, userId int, lat float64, lon float64) (mobileUser entities.Notification, err error)
	LocationTrack(ctx context.Context, userId int, status int) (err error)
	UserProfilePicture(ctx context.Context, userId int16) (mobileUser entities.MobileAppUser, err error)
	SoundSettingChange(ctx context.Context, userId int, status int) (err error)
	PushNotificationSetting(ctx context.Context, userId int, status int) ( err error)
}
