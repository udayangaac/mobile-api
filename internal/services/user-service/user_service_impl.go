package user_service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/mobile-api/internal/domain"
	"github.com/udayangaac/mobile-api/internal/entities"
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
		//Job:            param.JobDetails,
	}
	//if isAdded := u.RepoContainer.MobileUserRepo.AddMobileUser(ctx, mobileAppUser); isAdded {
	//	return
	//} else {
	//	err = errors_custom.ErrUnableToAddMobileAppUser
	//	return
	//}
	u.RepoContainer.MobileUserRepo.AddMobileUser(ctx, mobileAppUser)
	return
}

func (u *userService) GenerateToken(ctx context.Context, param domain.LoginRequest) (resp domain.LoginResponse, err error) {
	mobileAppUser := entities.MobileAppUser{}
	mobileAppUser, err = u.RepoContainer.MobileUserRepo.GetMobileUserByEmail(ctx, param.Email)
	log.Info(mobileAppUser)
	if err != nil {
		return
	}
	if mobileAppUser.HashPassword == sha256.GetHashString(param.Password) {
		resp.Email = mobileAppUser.Email
		resp.ID = int(mobileAppUser.ID)
		resp.Name = mobileAppUser.Name
		// resp.Avatar
		// resp.LbsNotification.ID
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
	notification, err = u.RepoContainer.MobileUserRepo.PullNotification(ctx, userId, lat, lon )
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

func (u *userService) SetLocationPermission(ctx context.Context, userId int) (resp domain.SettingsChangeResponse, err error) {
	/*mobileAppUser := entities.MobileAppUser{}
	mobileAppUser, err = u.RepoContainer.MobileUserRepo.LocationTrack(ctx, userId)*/

	return
}

func (u *userService) SetSoundStatus(ctx context.Context, userId int) (resp domain.SettingsChangeResponse, err error) {
	/*mobileAppUser := entities.MobileAppUser{}
	mobileAppUser, err = u.RepoContainer.MobileUserRepo.SoundSettingChange(ctx, userId)*/

	return
}

func (u *userService) SetPushNotificationPermission(ctx context.Context, userId int) (resp domain.SettingsChangeResponse, err error) {
	/*mobileAppUser := entities.MobileAppUser{}
	mobileAppUser, err = u.RepoContainer.MobileUserRepo.PushNotificationSetting(ctx, userId)*/

	return
}
