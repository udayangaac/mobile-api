package boot

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/mobile-api/internal/config"
	file_manager "github.com/udayangaac/mobile-api/internal/lib/file-manager"
	log_traceable "github.com/udayangaac/mobile-api/internal/lib/log-traceable"
	"github.com/udayangaac/mobile-api/internal/lib/orm"
	"github.com/udayangaac/mobile-api/internal/repositories"
	"github.com/udayangaac/mobile-api/internal/services"
	user_service "github.com/udayangaac/mobile-api/internal/services/user-service"
	http2 "github.com/udayangaac/mobile-api/internal/transport/http"
	"os"
	"os/signal"
	"syscall"
)

func Init(ctx context.Context) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	config.Configurations{
		new(config.ServerConfig),
		new(config.DatabaseConfig),
	}.Init(file_manager.NewYamlManager())

	if err := orm.InitDatabase(config.DatabaseConf); err != nil {
		log.Fatal(log_traceable.GetMessage(ctx, "Unable to open the database error :"+err.Error()))
	}

	// initialize the repositories
	repoContainer := repositories.RepoContainer{}
	repoContainer.MobileUserRepo = repositories.NewMobileAppUser()

	// initialized the services
	serviceContainer := services.Services{}
	serviceContainer.UserService = user_service.NewUserService(repoContainer)

	webService := http2.WebService{}
	webService.Port = config.ServerConf.Port
	webService.Services = serviceContainer
	webService.Init()

	select {
	case <-sigs:
		log.Info(log_traceable.GetMessage(ctx, "Shutting down Application"))
		// graceful shutdown code here
		if err := orm.CloseDatabase(); err != nil {
			log.Fatal(log_traceable.GetMessage(ctx, "Unable to close the database error :"+err.Error()))
		}
		log.Info(log_traceable.GetMessage(ctx, "Application stopped"))
		os.Exit(0)
	}
}
