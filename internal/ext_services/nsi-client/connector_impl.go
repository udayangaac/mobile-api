package nsi_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	log_traceable "github.com/udayangaac/mobile-api/internal/lib/log-traceable"
	"net/http"
	"time"
)

type nsiConnector struct {
	BaseUrl string
	Client  *http.Client
}

func NewNSIConnector(baseUrl string) NSIConnector {
	n := nsiConnector{
		BaseUrl: baseUrl,
	}
	n.init()
	return &n
}

func (n *nsiConnector) init() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	n.Client = &http.Client{Transport: tr}
}

func (n *nsiConnector) GetNotifications(ctx context.Context, param RequestBody) (notifications []Notification, geoRefId string, err error) {
	var (
		req     *http.Request
		res     *http.Response
		payload []byte
	)

	body := ResponseBody{}

	url := fmt.Sprintf("%s/tnsi/notifications", n.BaseUrl)

	if payload, err = json.Marshal(param); err != nil {
		return
	}

	if req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload)); err != nil {
		return
	}
	if n.Client == nil {
		n.init()
		log.Warn(log_traceable.GetMessage(ctx, "Re initialize http client"))
	}

	if res, err = n.Client.Do(req); err != nil {
		return
	}

	defer func() {
		err = res.Body.Close()
		if err != nil {
			log.Error(log_traceable.GetMessage(ctx, "Unable close the body of the request"))
		}
	}()

	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return
	}

	return body.Data.Notifications, body.Data.GeoRefID, err
}

func (n *nsiConnector) UpdateUserNotificationReaction(ctx context.Context, param TrackUserReaction) (err error) {
	var (
		req     *http.Request
		payload []byte
	)

	reqBody := UserReactionRequest{
		UserReaction: param.Status,
		Status:       0,
	}

	url := fmt.Sprintf("%s/tnsi/user/%v/notification/%v", n.BaseUrl, param.UserId, param.NotificationId)

	if payload, err = json.Marshal(reqBody); err != nil {
		return
	}

	if req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload)); err != nil {
		return
	}

	if n.Client == nil {
		n.init()
		log.Warn(log_traceable.GetMessage(ctx, "Re initialize http client"))
	}

	if _, err = n.Client.Do(req); err != nil {
		return
	}

	return err
}
