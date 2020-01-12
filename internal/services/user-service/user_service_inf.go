package user_service

import "github.com/udayangaac/mobile-api/internal/domain"

type UserService interface {
	AddMobileUser(param domain.SignUpRequest) (err error)
}
