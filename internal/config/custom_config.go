package config

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/udayangaac/mobile-api/internal/lib/file-manager"
	log_traceable "github.com/udayangaac/mobile-api/internal/lib/log-traceable"
)

var CustomConf CustomConfig

type CustomConfig struct {
	NSIUrl string `yaml:"nsi_url"`
}

func (cs *CustomConfig) Read(fm file_manager.FileManager) {
	path := fmt.Sprintf(`config/custom.yaml`)
	err := fm.Read(path, &CustomConf)
	if err != nil {
		log.Fatal(log_traceable.GetMessage(context.Background(), "Unable to read the custom.yaml file"))
	}
}
