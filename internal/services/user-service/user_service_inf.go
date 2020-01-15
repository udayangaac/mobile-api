package user_service

import "github.com/udayangaac/mobile-api/internal/domain"

// User management related functions are done by this service
type UserService interface {
	// add mobile user to the system,
	AddMobileUser(param domain.SignUpRequest) (err error)
	// generate JWT token with validating user credentials.
	GenerateToken(param domain.LoginRequest) (resp domain.LoginResponse, err error)
	// logout
	// add profile picture
	//
}
