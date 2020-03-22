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

	router.Handle("/register",
		transportHttp.NewServer(
			endpoint.SignUpEndpoints(ws.Services),
			decoder.SignUpDecoder,
			encoder.MainEncoder,
			serverOpts...)).Methods(http.MethodPost)

	router.Handle("/login",
		transportHttp.NewServer(
			endpoint.LoginEndpoints(ws.Services),
			decoder.LoginDecoder,
			encoder.MainEncoder,
			serverOpts...)).Methods(http.MethodPost)

	authSubRouter := router.PathPrefix("/auth").Subrouter()
	authSubRouter.Use(middleware.JwtMiddleware)

	authSubRouter.Handle("/logout",
		transportHttp.NewServer(
			endpoint.LogoutEndpoints(ws.Services),
			decoder.LogoutDecoder,
			encoder.MainEncoder,
			serverOpts...)).Methods(http.MethodPost)

	authSubRouter.Handle("/push",
		transportHttp.NewServer(
			endpoint.PushNotificationEndpoints(ws.Services),
			decoder.PushNotificationDecoder,
			encoder.MainEncoder,
			serverOpts...)).Methods(http.MethodPost)

	authSubRouter.Handle("/pull",
		transportHttp.NewServer(
			endpoint.PullNotificationEndpoints(ws.Services),
			decoder.PullNotificationDecoder,
			encoder.MainEncoder,
			serverOpts...)).Methods(http.MethodPost)

	authSubRouter.Handle("/settings/profile-picture",
		transportHttp.NewServer(
			endpoint.UserProfilePictureEndpoints(ws.Services),
			decoder.ProfilePictureDecoder,
			encoder.MainEncoder,
			serverOpts...)).Methods(http.MethodPost)

	authSubRouter.Handle("/settings/track-location",
		transportHttp.NewServer(
			endpoint.TrackLocationPermissionEndpoints(ws.Services),
			decoder.LocationStatusDecoder,
			encoder.MainEncoder,
			serverOpts...)).Methods(http.MethodPost)

	authSubRouter.Handle("/settings/push-status",
		transportHttp.NewServer(
			endpoint.PushPermissionEndpoints(ws.Services),
			decoder.PushNotificationStatusDecoder,
			encoder.MainEncoder,
			serverOpts...)).Methods(http.MethodPost)

	authSubRouter.Handle("/settings/login-status",
		transportHttp.NewServer(
			endpoint.LoginStatusEndpoints(ws.Services),
			decoder.LoginStatusDecoder,
			encoder.MainEncoder,
			serverOpts...)).Methods(http.MethodPost)

	authSubRouter.Handle("/settings/sound_notification",
		transportHttp.NewServer(
			endpoint.SoundPermissionEndpoints(ws.Services),
			decoder.SoundStatusDecoder,
			encoder.MainEncoder,
			serverOpts...)).Methods(http.MethodPost)

	authSubRouter.Handle("/notification-type",
		transportHttp.NewServer(
			endpoint.NotificationTypeEndpoints(ws.Services),
			decoder.NotificationTypeDecoder,
			encoder.MainEncoder,
			serverOpts...)).Methods(http.MethodGet)

	authSubRouter.Handle("/bank-list",
		transportHttp.NewServer(
			endpoint.BankListEndpoints(ws.Services),
			decoder.BankListDecoder,
			encoder.MainEncoder,
			serverOpts...)).Methods(http.MethodGet)

	authSubRouter.Handle("/user-profile",
		transportHttp.NewServer(
			endpoint.UserProfileEndpoints(ws.Services),
			decoder.UserProfileDecoder,
			encoder.MainEncoder,
			serverOpts...)).Methods(http.MethodGet)

	authSubRouter.Handle("/user-profile",
		transportHttp.NewServer(
			endpoint.UserProfileUpdateEndpoints(ws.Services),
			decoder.UserProfileDecoder,
			encoder.MainEncoder,
			serverOpts...)).Methods(http.MethodPost)

	authSubRouter.Handle("/track-user-location",
		transportHttp.NewServer(
			endpoint.TrackUserLocationEndpoints(ws.Services),
			decoder.TrackUserDecoder,
			encoder.MainEncoder,
			serverOpts...)).Methods(http.MethodPost)

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
