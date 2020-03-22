package decoder

import (
	"context"
	"encoding/json"
	"github.com/udayangaac/mobile-api/internal/domain"
	domain_errors "github.com/udayangaac/mobile-api/internal/errors_custom"
	"net/http"
)

func SignUpDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	signUpRequestParam := domain.SignUpRequest{}
	err := json.NewDecoder(r.Body).Decode(&signUpRequestParam)
	if err != nil {
		return nil, domain_errors.ErrBadRequest
	}
	return signUpRequestParam, nil
}

func LoginDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	loginRequestParam := domain.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&loginRequestParam)
	if err != nil {
		return nil, domain_errors.ErrBadRequest
	}
	return loginRequestParam, nil
}

func LogoutDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	logoutRequestParam := domain.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&logoutRequestParam)
	if err != nil {
		return nil, domain_errors.ErrBadRequest
	}
	return logoutRequestParam, nil
}

func PullNotificationDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	pullNotificationRequestParam := domain.PullRequest{}
	err := json.NewDecoder(r.Body).Decode(&pullNotificationRequestParam)
	if err != nil {
		return nil, domain_errors.ErrBadRequest
	}
	return  pullNotificationRequestParam, nil
}

func PushNotificationDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	pushNotificationRequestParam := domain.PushRequest{}
	err := json.NewDecoder(r.Body).Decode(&pushNotificationRequestParam)
	if err != nil {
		return nil, domain_errors.ErrBadRequest
	}
	return  pushNotificationRequestParam, nil
}

func UserProfileDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	userProfile := domain.UserProfile{}
	err := json.NewDecoder(r.Body).Decode(&userProfile)
	if err != nil {
		return nil, domain_errors.ErrBadRequest
	}
	return  userProfile, nil
}

func NotificationTypeDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	notificationType := domain.NotificationType{}
	err := json.NewDecoder(r.Body).Decode(&notificationType)
	if err != nil {
		return nil, domain_errors.ErrBadRequest
	}
	return  notificationType, nil
}

func BankListDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	bankList := domain.BankList{}
	err := json.NewDecoder(r.Body).Decode(&bankList)
	if err != nil {
		return nil, domain_errors.ErrBadRequest
	}
	return  bankList, nil
}

func ProfilePictureDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	profilePictureRequestParam := domain.ProfilePictureRequest{}
	err := json.NewDecoder(r.Body).Decode(&profilePictureRequestParam)
	if err != nil {
		return nil, domain_errors.ErrBadRequest
	}
	return  profilePictureRequestParam, nil

}

func LocationStatusDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	locationRequestParam := domain.LocationPermissionRequest{}
	err := json.NewDecoder(r.Body).Decode(&locationRequestParam)
	if err != nil {
		return nil, domain_errors.ErrBadRequest
	}
	return  locationRequestParam, nil
}

func PushNotificationStatusDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	pushRequestParam := domain.SettingChangeRequest{}
	err := json.NewDecoder(r.Body).Decode(&pushRequestParam)
	if err != nil {
		return nil, domain_errors.ErrBadRequest
	}
	return  pushRequestParam, nil
}

func LoginStatusDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	loginStatusRequestParam := domain.SettingChangeRequest{}
	err := json.NewDecoder(r.Body).Decode(&loginStatusRequestParam)
	if err != nil {
		return nil, domain_errors.ErrBadRequest
	}
	return  loginStatusRequestParam, nil
}

func SoundStatusDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	soundRequestParam := domain.SoundPermissionRequest{}
	err := json.NewDecoder(r.Body).Decode(&soundRequestParam)
	if err != nil {
		return nil, domain_errors.ErrBadRequest
	}
	return soundRequestParam , nil
}
func TrackUserDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	userLocation := domain.TrackUserLocation{}
	err := json.NewDecoder(r.Body).Decode(&userLocation)
	if err != nil {
		return nil, domain_errors.ErrBadRequest
	}
	return  userLocation, nil
}
