package user_service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/mobile-api/internal/config"
	"github.com/udayangaac/mobile-api/internal/domain"
	"github.com/udayangaac/mobile-api/internal/entities"
	jwt2 "github.com/udayangaac/mobile-api/internal/lib/jwt"
	"github.com/udayangaac/mobile-api/internal/lib/sha256"
	"github.com/udayangaac/mobile-api/internal/repositories"
)

type userService struct {
	RepoContainer repositories.RepoContainer
}

func NewUserService(repoContainer repositories.RepoContainer) UserService {
	return &userService{
		RepoContainer: repoContainer,
	}
}

func (u *userService) AddMobileUser(ctx context.Context, param domain.SignUpRequest) (err error) {
	mobileAppUser := entities.MobileAppUser{
		Name:           param.Name,
		Email:          param.Email,
		HashPassword:   sha256.GetHashString(param.Password),
		DOB:            param.DOB,
		Gender:         param.Gender,
		EmployeeStatus: param.JobStatus,
		Status:         1,
		Address:        param.Address,
		CivilStatus:    param.Married,
	}

	mobileUserConfiguration := entities.MobileUserConfiguration{
		LoginStatus:            1,
		PushNotificationStatus: 1,
		SoundStatus:            1,
		LocationServiceStatus:  1,
		AnyStatus:              0,
	}
	err = u.RepoContainer.MobileUserRepo.AddMobileUser(ctx, mobileAppUser, mobileUserConfiguration)
	return
}

func (u *userService) GenerateToken(ctx context.Context, param domain.LoginRequest) (resp domain.LoginResponse, err error) {
	mobileAppUser := entities.MobileAppUser{}
	jwt := jwt2.Resolver{
		SecretKey:     config.ServerConf.Jwt.Key,
		ValidDuration: config.ServerConf.Jwt.Duration,
	}

	mobileAppUser, err = u.RepoContainer.MobileUserRepo.GetMobileUserByEmail(ctx, param.Email)
	log.Info(mobileAppUser)
	if err != nil {
		return
	}

	if mobileAppUser.HashPassword == sha256.GetHashString(param.Password) {
		resp.Email = mobileAppUser.Email
		resp.ID = int(mobileAppUser.ID)
		resp.Name = mobileAppUser.Name
		claims := jwt2.Claims{Role: "user", UserId: mobileAppUser.ID}
		resp.Token, err = jwt.GenerateToken(claims)
	}

	return
}

func (u *userService) LogOut(ctx context.Context, param domain.LoginRequest) (resp domain.LogoutResponse, err error) {
	err = u.RepoContainer.MobileUserRepo.UserLogout(ctx)
	return
}

func (u *userService) PushNotification(ctx context.Context, userId int, lat float64, lon float64) (resp entities.Notification, err error) {
	notification := entities.Notification{}
	notification, err = u.RepoContainer.MobileUserRepo.PushNotification(ctx, userId, lat, lon)
	if err != nil {
		return
	}

	return notification, err
}

func (u *userService) PullNotification(ctx context.Context, userId int, lat float64, lon float64) (resp entities.Notification, err error) {
	notification := entities.Notification{}
	notification, err = u.RepoContainer.MobileUserRepo.PullNotification(ctx, userId, lat, lon)
	if err != nil {
		return
	}

	return notification, err
}

func (u *userService) UserProfilePicture(ctx context.Context, userId int16) (resp domain.SettingsChangeResponse, err error) {
	/*mobileAppUser := entities.MobileAppUser{}
	mobileAppUser, err = u.RepoContainer.MobileUserRepo.UserProfilePicture(ctx, userId)
	if err != nil {
		return
	}*/

	return
}

func (u *userService) SetLocationPermission(ctx context.Context, userId int, status int) (resp domain.SettingsChangeResponse, err error) {
	err = u.RepoContainer.MobileUserRepo.LocationTrack(ctx, userId, status)
	return
}

func (u *userService) SetSoundStatus(ctx context.Context, userId int, status int) (resp domain.SettingsChangeResponse, err error) {
	err = u.RepoContainer.MobileUserRepo.SoundSettingChange(ctx, userId, status)
	return
}

func (u *userService) SetPushNotificationPermission(ctx context.Context, userId int, status int) (resp domain.SettingsChangeResponse, err error) {
	err = u.RepoContainer.MobileUserRepo.PushNotificationSetting(ctx, userId, status)
	return
}

func (u *userService) SetLoginStatus(ctx context.Context, userId int, status int) (resp domain.SettingsChangeResponse, err error) {
	err = u.RepoContainer.MobileUserRepo.SetLoginStatus(ctx, userId, status)
	return
}

func (u *userService) NotificationTypes(ctx context.Context, userId int) (resp entities.AdvertismentsCategories, err error) {
	notificationTypes := entities.AdvertismentsCategories{}

	notificationTypes, err = u.RepoContainer.MobileUserRepo.NotificationTypes(ctx, userId)
	log.Info(notificationTypes)
	if err != nil {
		return
	}

	resp.Id = notificationTypes.Id
	resp.CategoryName = notificationTypes.CategoryName
	resp.Status = notificationTypes.Status

	return //notificationTypes, err
}

func (u *userService) UserProfile(ctx context.Context, userId int) (resp entities.MobileAppUser, err error) {
	notificationTypes := entities.MobileAppUser{}
	notificationTypes, err = u.RepoContainer.MobileUserRepo.UserProfile(ctx, userId)
	if err != nil {
		return
	}

	return notificationTypes, err
}