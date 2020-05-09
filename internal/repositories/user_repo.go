package repositories

import (
	"context"
	"github.com/udayangaac/mobile-api/internal/entities"
)

type MobileAppUserRepo interface {
	AddMobileUser(ctx context.Context, mobileUser entities.MobileAppUser, mobileUserConfiguration entities.MobileUserConfiguration) (user entities.MobileAppUser,err error)
	GetMobileUserByEmail(ctx context.Context, email string) (mobileUser entities.MobileAppUser, err error)
	UserLogout(ctx context.Context) (err error)
	PushNotification(ctx context.Context, userId int, lat float64, lon float64) (notification entities.Notification, err error)
	PullNotification(ctx context.Context, userId int, lat float64, lon float64) (notification entities.Notification, err error)
	LocationTrack(ctx context.Context, userId int, status int) (err error)
	UserProfilePicture(ctx context.Context, userId int16) (mobileUser entities.MobileAppUser, err error)
	SoundSettingChange(ctx context.Context, userId int, status int) (err error)
	PushNotificationSetting(ctx context.Context, userId int, status int) (err error)
	SetLoginStatus(ctx context.Context, userId int, status int) (err error)
	NotificationTypesList(ctx context.Context, userId int) (notificationType interface{}, err error)
	BankList(ctx context.Context, userId int) (banks interface{}, err error)
	GetUserProfile(ctx context.Context, userId int) (userProfile entities.MobileAppUser, bankSelected int, err error)
	UpdateUserProfile(ctx context.Context, mobileUser entities.MobileAppUser, mobileUserConfiguration entities.MobileUserConfiguration, userAdvertisementCategories []int, userBankList []int, userId int) (err error)
	TrackUserLocation(ctx context.Context, location entities.UserLocationChanges) (err error)
	TrackUserReaction(ctx context.Context, reaction entities.MobileUserResponse) (err error)
	UserViewedNotifications(ctx context.Context, reaction entities.MobileUserViewedAdvertisementList) (err error)
}
