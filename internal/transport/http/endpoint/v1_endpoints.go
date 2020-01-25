package endpoint

import (
	"context"
	endpoint2 "github.com/go-kit/kit/endpoint"
	"github.com/udayangaac/mobile-api/internal/domain"
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
		param := request.(domain.LoginRequest)
		err; = service.UserService.GenerateToken(ctx, param)
		if err != nil {
			return
		}
		response = domain.SuccessResponse{
			Message: "successfully Login",
		}
		return
	}
}