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
	LogOut(ctx context.Context, param domain.LoginRequest) (resp domain.LogoutResponse, err error)

	// push notification
	PushNotification(ctx context.Context, param domain.PushRequest) (resp domain.PushResponse, err error)

	// pull notification
	PullNotification(ctx context.Context, param domain.PullRequest) (resp domain.PullResponse, err error)

	// Set Location Permission
	SetLocationPermission(ctx context.Context, param domain.LocationPermissionRequest) (resp domain.SettingsChangeResponse, err error)

	// Set Push Notification
	SetPushNotificationPermission(ctx context.Context, param domain.SettingChangeRequest) (resp domain.SettingsChangeResponse, err error)

	// Set Sound Permission
	SetSoundStatus(ctx context.Context, param domain.SoundPermissionRequest) (resp domain.SettingsChangeResponse, err error)

	// user Profile Picture
	UserProfilePicture(ctx context.Context, param domain.ProfilePictureRequest) (resp domain.SettingsChangeResponse, err error)
}
