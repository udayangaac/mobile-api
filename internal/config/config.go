package config

import (
	"github.com/udayangaac/mobile-api/internal/lib/file-manager"
)

type Configuration interface {
	Read(fm file_manager.FileManager)
}
type Configurations []Configuration

func (configs Configurations) Init(fm file_manager.FileManager) {
	for _, c := range configs {
		c.Read(fm)
	}
}
