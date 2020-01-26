package decoder

import (
	"context"
	"encoding/json"
	"github.com/udayangaac/mobile-api/internal/domain"
	domain_errors "github.com/udayangaac/mobile-api/internal/errors_custom"
	"net/http"
)

funcfunc SignUpDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
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
	loginRequestParam := domain.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&loginRequestParam)
	if err != nil {
		return nil, domain_errors.ErrBadRequest
	}
	return loginRequestParam, nil
}

