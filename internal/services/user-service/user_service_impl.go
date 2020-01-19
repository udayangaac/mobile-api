package user_service

import (
	"context"
	"github.com/udayangaac/mobile-api/internal/domain"
	"github.com/udayangaac/mobile-api/internal/entities"
	"github.com/udayangaac/mobile-api/internal/errors_custom"
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
		Gender:         "Unknown",
		EmployeeStatus: 0,
		Status:         param.JobStatus,
	}
	if isAdded := u.RepoContainer.MobileUserRepo.AddMobileUser(ctx, mobileAppUser); isAdded {
		return
	} else {
		err = errors_custom.ErrUnableToAddMobileAppUser
		return
	}
}

func (u userService) GenerateToken(ctx context.Context, param domain.LoginRequest) (resp domain.LoginResponse, err error) {
	mobileAppUser := entities.MobileAppUser{}
	mobileAppUser, err = u.RepoContainer.MobileUserRepo.GetMobileUserByEmail(ctx, param.Email)
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
