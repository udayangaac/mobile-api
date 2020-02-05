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
		response = domain.SuccessResponse{
			Message: "successfully added the user",
		}
		return
	}
}

func LoginEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		/*param := request.(domain.LoginRequest)
		err, _ = service.UserService.GenerateToken(ctx, param)
		if err != nil {
			return
		}
		response = domain.SuccessResponse{
			Message: "successfully Login",
		}*/
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

func PullNotificationEndpoints (service services.Services) endpoint2.Endpoint {
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

func PushNotificationEndpoints (service services.Services) endpoint2.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		notification := entities.Notification{}
		param := request.(domain.PushRequest)
		notification, err = service.UserService.PushNotification(ctx, param.UserId, param.Location.Lat, param.Location.Lon)
		log.Info(param.UserId)
		if err != nil {
			return
		}
		response = notification
		return
	}
}

func UserProfilePictureEndpoints (service services.Services) endpoint2.Endpoint {
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

func TrackLocationPermissionEndpoints (service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		/*param := request.(domain.LocationPermissionRequest)
		err, _ = service.UserService.SetLocationPermission(ctx, param)
		if err != nil {
			return
		}
		response = domain.SuccessResponse{
			Message: "successfully Login ",
		}*/
		return
	}
}

func SoundPermissionEndpoints (service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		/*param := request.(domain.SoundPermissionRequest)
		err, _ = service.UserService.SetSoundStatus(ctx, param)
		if err != nil {
			return
		}
		response = domain.SuccessResponse{
			Message: "successfully Login ",
		}*/
		return
	}
}
