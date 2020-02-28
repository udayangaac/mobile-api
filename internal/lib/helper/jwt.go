package helper

import (
	"context"
	"github.com/udayangaac/mobile-api/internal/domain"
	"github.com/udayangaac/mobile-api/internal/lib/jwt"
)

func GetAuthenticationParam(ctx context.Context) (authParam domain.AuthenticationParam) {
	//todo: null must be handled from the location.
	values, ok := ctx.Value("claims").(jwt.Claims)
	if ok {
		return domain.AuthenticationParam{UserId: int(values.UserId), Role: 1}
	}
	return
}
