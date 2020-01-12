package helper

import (
	"context"
	"github.com/udayangaac/mobile-api/internal/domain"
	"github.com/udayangaac/mobile-api/internal/lib/jwt"
)

func GetAuthenticationParam(ctx context.Context) (authParam domain.AuthenticationParam) {
	//todo: null must be handled from the location.
	values := ctx.Value("claims").(jwt.Claims)
	return
}
