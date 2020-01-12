package middleware

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/mobile-api/internal/config"
	"github.com/udayangaac/mobile-api/internal/domain"
	domain_errors "github.com/udayangaac/mobile-api/internal/errors_custom"
	"github.com/udayangaac/mobile-api/internal/lib/jwt"
	log_traceable "github.com/udayangaac/mobile-api/internal/lib/log-traceable"
	"net/http"
	"strings"
)

var domainError *domain_errors.CustomError

func init() {
	domainError = new(domain_errors.CustomError)
}

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorizationHeader := r.Header.Get("Authorization")
		// note : Bearer<space> must be defined
		splitToken := strings.Split(authorizationHeader, "Bearer ")

		if len(splitToken) < 2 {
			errorContent := domainError.GetErrorContent(r.Context(), domain_errors.ErrEmptyToken)
			errResp := domain.ErrorResponse{
				Code:    errorContent.ApplicationErrorCode,
				Message: errorContent.CustomMessage,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(errorContent.HttpStatusCode)
			err := json.NewEncoder(w).Encode(errResp)
			if err != nil {
				log.Warn(log_traceable.GetMessage(r.Context(), "Invalid token error:", err.Error()))
			}
			return
		}
		reqToken := splitToken[1]
		// validate the token
		isValid, extractedClaims, err := jwt.Resolver{
			SecretKey:     config.ServerConf.Jwt.Key,
			ValidDuration: config.ServerConf.Jwt.Duration,
		}.ValidateToken(reqToken)

		if err != nil {
			log.Error(log_traceable.GetMessage(r.Context(), "Token validation error:", err.Error()))
			err = domain_errors.ErrInvalidToken
			errorContent := domainError.GetErrorContent(r.Context(), err)
			errResp := domain.ErrorResponse{
				Code:    errorContent.ApplicationErrorCode,
				Message: errorContent.CustomMessage,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(errorContent.HttpStatusCode)
			err = json.NewEncoder(w).Encode(errResp)
			if err != nil {
				log.Warn(log_traceable.GetMessage(r.Context(), "Invalid token error:", err.Error()))
			}
			return
		}

		if !isValid {
			err = domain_errors.ErrInvalidToken
			errorContent := domainError.GetErrorContent(r.Context(), err)
			errResp := domain.ErrorResponse{
				Code:    errorContent.ApplicationErrorCode,
				Message: errorContent.CustomMessage,
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(errorContent.HttpStatusCode)
			err = json.NewEncoder(w).Encode(errResp)
			if err != nil {
				log.Warn(log_traceable.GetMessage(r.Context(), "Invalid token error:", err.Error()))
			}
			return
		}

		ctx := context.WithValue(r.Context(), "claims", extractedClaims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			return
		}
		next.ServeHTTP(w, r)
	})
}
