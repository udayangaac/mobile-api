package user_service

import (
	"context"
	"github.com/udayangaac/mobile-api/internal/domain"
	"github.com/udayangaac/mobile-api/internal/entities"
)

type Notification struct {
}

// User management related functions are done by this service
type UserService interface {
	// add mobile user to the system,
	AddMobileUser(ctx context.Context, param domain.SignUpRequest) (resp domain.LoginResponse, err error)

	// generate JWT token with validating user credentials.
	GenerateToken(ctx context.Context, param domain.LoginRequest) (resp domain.LoginResponse, err error)

	// Logout
	LogOut(ctx context.Context, param domain.LoginRequest) (resp domain.LogoutResponse, err error)

	// push notification
	PushNotification(ctx context.Context, userId int, lat float64, lon float64) (resp entities.Notification, err error)

	// Notification Types
	NotificationTypes(ctx context.Context, userId int) (resp interface{}, err error)

	// Bank List
	BankList(ctx context.Context, userId int) (resp interface{}, err error)

	// Get Mobile user
	GetUserProfile(ctx context.Context, userId int) (resp domain.UserProfileResponse, err error)

	// User Profile
	UpdateUserProfile(ctx context.Context, param domain.UserProfile, userId int, advertisementCategory []int, bankIdList []int) (err error)

	// Track User Location
	TrackUserLocation(ctx context.Context, param domain.TrackUserLocation) (err error)

	// Track User Reactions
	TrackUserReaction(ctx context.Context, param domain.TrackUserReaction) (err error)

	// Track User Push Reactions
	TrackUserPushReaction(ctx context.Context, param domain.TrackUserReaction) (err error)

	// Record User Viewed Notifications
	UserViewedNotifications(ctx context.Context, param domain.UserViewedNotification) (err error)

	// pull notification
	PullNotification(ctx context.Context, userId int, lat float64, lon float64) (resp interface{}, err error)

	// Pull search Notifications
	PullSearchNotification(ctx context.Context, userId int, text string) (resp interface{}, err error)

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
