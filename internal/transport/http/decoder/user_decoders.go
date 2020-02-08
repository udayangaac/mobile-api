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

func SoundStatusDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	soundRequestParam := domain.SoundPermissionRequest{}
	err := json.NewDecoder(r.Body).Decode(&soundRequestParam)
	if err != nil {
		return nil, domain_errors.ErrBadRequest
	}
	return soundRequestParam , nil
}