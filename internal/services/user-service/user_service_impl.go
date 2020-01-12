package user_service

import "github.com/udayangaac/mobile-api/internal/domain"

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (u *userService) AddMobileUser(param domain.SignUpRequest) (err error) {
	return
}
