package errors_custom

import (
	"context"
	log "github.com/sirupsen/logrus"
	log_traceable "github.com/udayangaac/mobile-api/internal/lib/log-traceable"
	"sync"
)

type ErrorContent struct {
	HttpStatusCode       int
	ApplicationErrorCode int
	CustomMessage        string
}

type CustomError struct {
	mu sync.RWMutex
}

func (ce CustomError) GetErrorContent(ctx context.Context, err error) (errorContent ErrorContent) {
	ce.mu.RLock()
	defer ce.mu.RUnlock()
	val, ok := customErrorMap[err]
	if ok {
		errorContent = val
		return
	}
	log.Warn(log_traceable.GetMessage(ctx, "Returned error as internal server error."))
	return UnknownErrorContent
}
