package endpoint

import (
	"context"
	endpoint2 "github.com/go-kit/kit/endpoint"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/mobile-api/internal/domain"
	"github.com/udayangaac/mobile-api/internal/entities"
	"github.com/udayangaac/mobile-api/internal/errors_custom"
	"github.com/udayangaac/mobile-api/internal/services"
)

func SignUpEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		token := domain.LoginResponse{}
		param := request.(domain.SignUpRequest)
		token, err = service.UserService.AddMobileUser(ctx, param)
		log.Info(err)
		if err != nil {
			return
		}
		response = token
		return
	}
}

func LoginEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		token := domain.LoginResponse{}
		param := request.(domain.LoginRequest)
		token, err = service.UserService.GenerateToken(ctx, param)
		if err != nil {
			return
		}
		response = token
		if token.ID == 0 {
			err = errors_custom.ErrInvalidCredentials
		}
		return
	}
}

func LogoutEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		/*param := request.(domain.LoginRequest)
		err, _ = service.UserService.Logout(ctx, param)
		if err != nil {
			return
		}
		response = domain.SuccessResponse{
			Message: "successfully Logout",
		}*/
		return
	}
}

func PullNotificationEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(domain.PullRequest)
		log.Info("pull request")
		response, err = service.UserService.PullNotification(ctx, param.UserId, param.Location.Lat, param.Location.Lon)
		log.Info(param.UserId)
		if err != nil {
			return
		}
		return
	}
}

func PushNotificationEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		notification := entities.Notification{}
		param := request.(domain.PushRequest)
		notification, err = service.UserService.PushNotification(ctx, param.UserId, param.Location.Lat, param.Location.Lon)
		if err != nil {
			return
		}
		response = notification
		return
	}
}

func NotificationTypeEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(entities.UserParam)
		response, err = service.UserService.NotificationTypes(ctx, param.UserId)
		if err != nil {
			return
		}
		return
	}
}

func BankListEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(entities.UserParam)
		response, err = service.UserService.BankList(ctx, param.UserId)
		if err != nil {
			return
		}
		return
	}
}

func UserProfileEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		userProfile := domain.UserProfileResponse{}
		param := request.(domain.UserProfile)
		userProfile, err = service.UserService.GetUserProfile(ctx, param.UserId)
		if err != nil {
			return
		}
		response = userProfile
		return
	}
}

func UserProfileUpdateEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(domain.UserProfile)
		err = service.UserService.UpdateUserProfile(ctx, param, param.UserId, param.AdvertisementCatId, param.BankIdList)
		log.Info(err)
		if err != nil {
			return
		}
		response = domain.SuccessResponse{
			Message: "successfully Updated the User",
		}
		return
	}
}

func TrackUserLocationEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(domain.TrackUserLocation)

		err = service.UserService.TrackUserLocation(ctx, param)
		log.Info(err)
		if err != nil {
			return
		}
		response = domain.SuccessResponse{
			Message: "success",
		}
		return
	}
}

func TrackUserReactionEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(domain.TrackUserReaction)
		log.Info(param)

		err = service.UserService.TrackUserReaction(ctx, param)
		log.Info(err)
		if err != nil {
			return
		}
		response = domain.SuccessResponse{
			Message: "success",
		}
		return
	}
}

func UserNotificationViewedStatsEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(domain.UserViewedNotification)
		log.Info(param)

		err = service.UserService.UserViewedNotifications(ctx, param)
		log.Info(err)
		if err != nil {
			return
		}
		response = domain.SuccessResponse{
			Message: "success",
		}
		return
	}
}

func UserProfilePictureEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		/*param := request.(domain.PullRequest)
		err, _ = service.UserService.PushNotification(ctx, userId, lat, lon)
		if err != nil {
			return
		}
		response = domain.SuccessResponse{
			Message: "successfully Push",
		}*/
		return
	}
}

func TrackLocationPermissionEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(domain.LocationPermissionRequest)
		_, err = service.UserService.SetLocationPermission(ctx, param.UserId, param.Status)
		if err != nil {
			return
		}

		response = domain.SuccessResponse{
			Message: "successfully Updated",
		}
		return
	}
}

func SoundPermissionEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(domain.SoundPermissionRequest)
		_, err = service.UserService.SetSoundStatus(ctx, param.UserId, param.Status)

		if err != nil {
			return
		}

		response = domain.SuccessResponse{
			Message: "successfully Updated",
		}

		return
	}
}

func PushPermissionEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(domain.SettingChangeRequest)
		_, err = service.UserService.SetPushNotificationPermission(ctx, param.UserId, param.Status)

		if err != nil {
			return
		}

		response = domain.SuccessResponse{
			Message: "successfully Updated",
		}

		return
	}
}

func LoginStatusEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(domain.SettingChangeRequest)
		_, err = service.UserService.SetLoginStatus(ctx, param.UserId, param.Status)

		if err != nil {
			return
		}

		response = domain.SuccessResponse{
			Message: "successfully Updated",
		}

		return
	}
}
