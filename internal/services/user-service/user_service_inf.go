package user_service

import (
	"context"
	"github.com/udayangaac/mobile-api/internal/domain"
	"github.com/udayangaac/mobile-api/internal/entities"
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
	PushNotification(ctx context.Context, userId int, lat float64, lon float64) (resp entities.Notification, err error)

	// Notification Types
	NotificationTypes(ctx context.Context, userId int) (resp [12]domain.NotificationTypes, err error)

	// Get Mobile user
	GetUserProfile(ctx context.Context, userId int) (resp domain.UserProfileResponse, err error)

	// User Profile
	UpdateUserProfile(ctx context.Context, param domain.UserProfile) (err error)

	// pull notification
	PullNotification(ctx context.Context, userId int, lat float64, lon float64) (resp entities.Notification, err error)

	// Set Location Permission
	SetLocationPermission(ctx context.Context, userId int, Status int) (resp domain.SettingsChangeResponse, err error)

	// Set Push Notification
	SetPushNotificationPermission(ctx context.Context, userId int, Status int) (resp domain.SettingsChangeResponse, err error)

	// Set Sound Permission
	SetSoundStatus(ctx context.Context, userId int, Status int) (resp domain.SettingsChangeResponse, err error)

	// user Profile Picture
	UserProfilePicture(ctx context.Context, userId int16) (resp domain.SettingsChangeResponse, err error)

	// Set Login Status
	SetLoginStatus(ctx context.Context, userId int, Status int) (resp domain.SettingsChangeResponse, err error)
}
