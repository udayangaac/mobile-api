package endpoint

import (
	"context"
	endpoint2 "github.com/go-kit/kit/endpoint"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/mobile-api/internal/domain"
	"github.com/udayangaac/mobile-api/internal/entities"
	"github.com/udayangaac/mobile-api/internal/services"
)

func SignUpEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		param := request.(domain.SignUpRequest)
		err = service.UserService.AddMobileUser(ctx, param)
		if err != nil {
			return
		}

		log.Info(err)
		response = domain.SuccessResponse{
			Message: "successfully added the user",
		}
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
		notification := entities.Notification{}
		param := request.(domain.PullRequest)
		notification, err = service.UserService.PullNotification(ctx, param.UserId, param.Location.Lat, param.Location.Lon)
		log.Info(param.UserId)
		if err != nil {
			return
		}
		response = notification
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