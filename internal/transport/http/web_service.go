package http

import (
	"context"
	"fmt"
	transportHttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	log_traceable "github.com/udayangaac/mobile-api/internal/lib/log-traceable"
	"github.com/udayangaac/mobile-api/internal/services"
	"github.com/udayangaac/mobile-api/internal/transport/http/decoder"
	"github.com/udayangaac/mobile-api/internal/transport/http/encoder"
	"github.com/udayangaac/mobile-api/internal/transport/http/endpoint"
	"github.com/udayangaac/mobile-api/internal/transport/http/middleware"
	"net/http"
	"time"
)

var server *http.Server

var serverOpts = []transportHttp.ServerOption{
	transportHttp.ServerErrorEncoder(encoder.ErrorEncoder),
	transportHttp.ServerBefore(reqFunc),
	transportHttp.ServerAfter(respFunc),
}

func reqFunc(ctx context.Context, r *http.Request) (ctxR context.Context) {
	ctxWithTimeout, _ := context.WithTimeout(ctx, 10*time.Second)
	uuidStr := uuid.New().String()
	ctxR = context.WithValue(ctxWithTimeout, "uuid_str", uuidStr)
	log.Trace(log_traceable.GetMessage(ctxR, "Started to process request URL:", r.URL, "Method:", r.Method))
	return
}

func respFunc(ctx context.Context, w http.ResponseWriter) (ctxR context.Context) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return ctx
}

type WebService struct {
	Port     int
	Services services.Services
}

func (ws *WebService) Init() {
	routerRoot := mux.NewRouter()
	routerRoot.Use(middleware.CORSMiddleware)
	router := routerRoot.PathPrefix("/api/1.0").Subrouter()

	router.Handle("/signup",
		transportHttp.NewServer(
			endpoint.SignUpEndpoints(ws.Services),
			decoder.SignUpDecoder,
			encoder.MainEncoder,
			serverOpts...)).Methods(http.MethodPost)

	authSubRouter := router.PathPrefix("/auth").Subrouter()
	authSubRouter.Use(middleware.JwtMiddleware)

	server = &http.Server{
		Addr:         fmt.Sprintf(":%v", ws.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      middleware.CORSMiddleware(router),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Error(log_traceable.GetMessage(context.Background(), "Unable to start the server error.", err.Error()))
		}
	}()
}

func (ws *WebService) Stop(ctx context.Context) {
	err := server.Shutdown(ctx)
	if err != nil {
		log.Error(log_traceable.GetMessage(ctx, "Error shutting down application error.", err.Error()))
	}
}
