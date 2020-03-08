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

func (u *userService) AddMobileUser(ctx context.Context, param domain.SignUpRequest) (resp domain.LoginResponse, err error) {

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

	jwt := jwt2.Resolver{
		SecretKey:     config.ServerConf.Jwt.Key,
		ValidDuration: config.ServerConf.Jwt.Duration,
	}
	
	userDetails := entities.MobileAppUser{}

	userDetails, err = u.RepoContainer.MobileUserRepo.AddMobileUser(ctx, mobileAppUser, mobileUserConfiguration)

	if err != nil{
		return
	}
	
	resp.Email = userDetails.Email
	resp.ID = int(userDetails.ID)
	resp.Name = userDetails.Name
	claims := jwt2.Claims{Role: "user", UserId: mobileAppUser.ID}
	resp.Token, err = jwt.GenerateToken(claims)

	return
}

func (u *userService) UpdateUserProfile(ctx context.Context, param domain.UserProfile, userId int, advertisementCategory []int, bankList []int) (err error) {
	mobileAppUser := entities.MobileAppUser{
		Name:               	param.Name,
		Email:              	param.Email,
		HashPassword:       	sha256.GetHashString(param.Password),
		DOB:                	param.DOB,
		Gender:             	param.Gender,
		EmployeeStatus:     	param.JobStatus,
		Status:             	1,
		Address:            	param.Address,
		CivilStatus:        	param.CivilStatus,
		JobCompanyName:     	param.JobDetails.Name,
		JobCompanyLocation: 	param.JobDetails.Address,
		Kids:               	param.Kids,
		LoginStatus:            param.Configuration.LoginStatus,
		PushNotificationStatus: param.Configuration.PushNotificationStatus,
		SoundStatus:            param.Configuration.SoundStatus,
		LocationServiceStatus:  param.Configuration.LocationServiceStatus,
		AnyStatus:              param.Configuration.AnyStatus,
	}

	mobileUserConfiguration := entities.MobileUserConfiguration{
		LoginStatus:            param.Configuration.LoginStatus,
		PushNotificationStatus: param.Configuration.PushNotificationStatus,
		SoundStatus:            param.Configuration.SoundStatus,
		LocationServiceStatus:  param.Configuration.LocationServiceStatus,
		AnyStatus:              param.Configuration.AnyStatus,
	}

	//userAdvertisement :=  entities.UserAdvertisementCategories{}

	err = u.RepoContainer.MobileUserRepo.UpdateUserProfile(ctx, mobileAppUser, mobileUserConfiguration, advertisementCategory, bankList,userId)
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

func (u *userService) BankList(ctx context.Context, userId int) (resp interface{}, err error) {
	var bank interface{}
	bank, err = u.RepoContainer.MobileUserRepo.BankList(ctx, userId)
	bankLists := []domain.BankListResponse{}
	if err != nil {
		return
	}
	categories, ok := bank.([]entities.Banks)
	if !ok {
		return nil, errors.New("cannot cast []entities.UserBankList")
	}
	for _, val := range categories {
		bankList := domain.BankListResponse{}
		bankList.Id = int(val.ID)
		bankList.BankName = val.Name
		bankLists = append(bankLists, bankList)
	}

	return bankLists, nil
}

func (u *userService) GetUserProfile(ctx context.Context, userId int) (resp domain.UserProfileResponse, err error) {

	userProfile := entities.MobileAppUser{}

	userProfile, err = u.RepoContainer.MobileUserRepo.GetUserProfile(ctx, userId)
	log.Info(userProfile.LoginStatus)
	if err != nil {
		return
	}
	resp.UserId = userProfile.ID
	resp.Name = userProfile.Name
	resp.Email = userProfile.Email
	resp.DOB = userProfile.DOB
	resp.Gender = userProfile.Gender
	resp.EmployeeStatus = userProfile.EmployeeStatus
	resp.Address = userProfile.Address
	resp.CivilStatus = userProfile.CivilStatus
	resp.JobCompanyName = userProfile.JobCompanyName
	resp.JobCompanyLocation = userProfile.JobCompanyLocation
	resp.Kids = userProfile.Kids
	resp.Configuration.LoginStatus = userProfile.LoginStatus
	resp.Configuration.PushNotificationStatus = userProfile.PushNotificationStatus
	resp.Configuration.SoundStatus =  userProfile.SoundStatus
	resp.Configuration.LocationServiceStatus = userProfile.LocationServiceStatus
	resp.Configuration.AnyStatus = userProfile.AnyStatus
	resp.UserAdvertisementsCategories   = userProfile.UserAdvertisementCategories
	
	return resp,err
}
