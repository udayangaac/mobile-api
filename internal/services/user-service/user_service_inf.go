package user_service

import (
	"context"
	"github.com/udayangaac/mobile-api/internal/domain"
)

// User management related functions are done by this service
type UserService interface {
	// add mobile user to the system,
	AddMobileUser(ctx context.Context, param domain.SignUpRequest) (err error)

	// generate JWT token with validating user credentials.
	GenerateToken(ctx context.Context, param domain.LoginRequest) (resp domain.LoginResponse, err error)

	// Logout
	Logout(ctx context.Context, param domain.LoginRequest) (resp domain.LoginResponse, err error)

	// push notification
	PushNotification(ctx context.Context, param domain.PullRequest) (resp domain.PushResponse, err error)
	// add profile picture

}
