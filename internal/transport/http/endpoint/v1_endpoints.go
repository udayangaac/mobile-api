package endpoint

import (
	"context"
	endpoint2 "github.com/go-kit/kit/endpoint"
	"github.com/udayangaac/mobile-api/internal/services"
)

func SignUpEndpoints(service services.Services) endpoint2.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return
	}
}
