package dbl_go

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/resty.v1"
	"net/http"
	"time"
)

const (
	baseURL                  = "https://discordbots.org/api"
	errorWhileSendingRequest = "error while sending request:"
)

type DBLApi struct {
	accessToken    string
	requestTimeout time.Duration
}

// Returns new DBLApi struct initialized with optimal values
func NewDBLApi(accessToken string) DBLApi {
	return DBLApi{
		accessToken:    accessToken,
		requestTimeout: time.Second * 10,
	}
}

// Post bot guild count to the website
func (d DBLApi) PostStatsSimple(guildCount int) error {
	url := d.getRequestURL("/bots/stats")
	params := map[string]string{"server_count": fmt.Sprintf("%d", guildCount)}

	result, err := d.getBaseRequest().SetBody(params).Post(url)
	if err != nil {
		return errors.Wrap(err, errorWhileSendingRequest)
	}

	statusCode := result.StatusCode()
	if statusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("server returned %d http code", statusCode))
	}

	return nil
}

func (d DBLApi) getBaseRequest() *resty.Request {
	ctx, _ := context.WithTimeout(context.Background(), d.requestTimeout)

	return resty.R().SetHeader("Authorization", d.accessToken).SetContext(ctx)
}

// Appends endpoint to baseURL, and return full request URL
func (d DBLApi) getRequestURL(endpoint string) string {
	return fmt.Sprintf("%s%s", baseURL, endpoint)
}
