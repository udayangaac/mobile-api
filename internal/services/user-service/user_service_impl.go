package user_service

import (
	"context"
	"errors"
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
		Name:         param.Name,
		Email:        param.Email,
		HashPassword: sha256.GetHashString(param.Password),
		DOB:          param.DOB,
	}
	mobileUserConfiguration := entities.MobileUserConfiguration{
		LoginStatus:            1,
		PushNotificationStatus: 1,
		SoundStatus:            1,
		LocationServiceStatus:  1,
		AnyStatus:              0,
	}

	err = u.RepoContainer.MobileUserRepo.AddMobileUser(ctx, mobileAppUser, mobileUserConfiguration)
	log.Info(err)
	return
}

func (u *userService) UpdateUserProfile(ctx context.Context, param domain.UserProfile, userId int) (err error) {
	mobileAppUser := entities.MobileAppUser{
		Name:               param.Name,
		Email:              param.Email,
		HashPassword:       sha256.GetHashString(param.Password),
		DOB:                param.DOB,
		Gender:             param.Gender,
		EmployeeStatus:     param.JobStatus,
		Status:             1,
		Address:            param.Address,
		CivilStatus:        param.CivilStatus,
		JobCompanyName:     param.JobDetails.Name,
		JobCompanyLocation: param.JobDetails.Address,
		Kids:               param.Kids,
	}

	mobileUserConfiguration := entities.MobileUserConfiguration{
		LoginStatus:            param.Configuration.LoginStatus,
		PushNotificationStatus: param.Configuration.PushNotificationStatus,
		SoundStatus:            param.Configuration.SoundStatus,
		LocationServiceStatus:  param.Configuration.LocationServiceStatus,
		AnyStatus:              param.Configuration.AnyStatus,
	}

	userAdvertisementCategories := entities.UserAdvertisementCategories{}

	// bank := entities.MobileUserBank{}

	err = u.RepoContainer.MobileUserRepo.UpdateUserProfile(ctx, mobileAppUser, mobileUserConfiguration, userAdvertisementCategories, userId)
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

func (u *userService) NotificationTypes(ctx context.Context, userId int) (resp interface{}, err error) {
	var notification interface{}
	notification, err = u.RepoContainer.MobileUserRepo.NotificationTypesList(ctx, userId)
	notificationTypes := []domain.NotificationTypes{}
	if err != nil {
		return
	}
	categories, ok := notification.([]entities.AdvertisementsCategories)
	if !ok {
		return nil, errors.New("cannot cast []entities.UserAdvertisementCategories")
	}
	for _, val := range categories {
		notificationType := domain.NotificationTypes{}
		notificationType.Id = int(val.ID)
		notificationType.CategoryName = val.CategoryName
		notificationTypes = append(notificationTypes, notificationType)
	}

	return notificationTypes, nil
}

func (u *userService) GetUserProfile(ctx context.Context, userId int) (resp entities.MobileAppUser, err error) {
	var user interface{}
	//userProfile := domain.UserProfileResponse{}
	//user := userId
	user, err = u.RepoContainer.MobileUserRepo.GetUserProfile(ctx, userId)
	log.Info(user)
	if err != nil {
		return
	}
	//userProfile.Name = user.Name

	return
}
