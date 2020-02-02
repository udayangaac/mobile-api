package repositories

import (
	"context"
	"github.com/udayangaac/mobile-api/internal/entities"
)

type MobileAppUserRepo interface {
	AddMobileUser(ctx context.Context, mobileUser entities.MobileAppUser)
	GetMobileUserByEmail(ctx context.Context, email string) (mobileUser entities.MobileAppUser, err error)
	UserLogout(ctx context.Context) (err error)
	PushNotification(ctx context.Context, email string) (mobileUser entities.MobileAppUser, err error)
	PullNotification(ctx context.Context, email string) (mobileUser entities.MobileAppUser, err error)
	LocationTrack(ctx context.Context, email string) (mobileUser entities.MobileAppUser, err error)
	UserProfilePicture(ctx context.Context, email string) (mobileUser entities.MobileAppUser, err error)
	SoundSettingChange(ctx context.Context, email string) (mobileUser entities.MobileAppUser, err error)
}
