package api

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	*resty.Client
}

func NewClient(apiToken string, apiEndpoint string) *Client {
	client := Client{}
	client.Client = resty.New()
	client.SetAuthToken(apiToken)
	client.SetHostURL(apiEndpoint)
	// Headers for all request
	client.SetHeader("Accept", "application/json")
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "dog_go_api",
	})
	return &client
}
