package encoder

import (
	"context"
	"encoding/json"
	"github.com/udayangaac/mobile-api/internal/domain"
	domain_errors "github.com/udayangaac/mobile-api/internal/errors_custom"
	"net/http"
)

var domainError *domain_errors.CustomError

func init() {
	domainError = new(domain_errors.CustomError)
}

func MainEncoder(_ context.Context, w http.ResponseWriter, response interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	errorContent := domainError.GetErrorContent(ctx, err)
	errResp := domain.ErrorResponse{
		Code:    errorContent.ApplicationErrorCode,
		Message: errorContent.CustomMessage,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorContent.HttpStatusCode)
	err = json.NewEncoder(w).Encode(errResp)
	if err != nil {
		return
	}
}
