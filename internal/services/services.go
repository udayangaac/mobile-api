package services

import (
	user_service "github.com/udayangaac/mobile-api/internal/services/user-service"
)

type Services struct {
	UserService user_service.UserService
}
