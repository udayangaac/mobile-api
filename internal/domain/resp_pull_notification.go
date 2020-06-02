package domain

import (
	nsi_client "github.com/udayangaac/mobile-api/internal/ext_services/nsi-client"
)

type PullResponse struct {
	Error     bool   `json:"error"`
	Offers    []nsi_client.Notification `json:"offers"`
	
}