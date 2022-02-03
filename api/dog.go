package api

import (
	"github.com/go-resty/resty/v2"
)

const (
	BaseURLV2 = "http://dog-ubuntu-server.lxd:8000/api"
)

type Client struct {
	*resty.Client
}

func NewClient(apiKey string) *Client {
	client := Client{}
	client.Client = resty.New()
	client.SetHeader("apikey", apiKey)
	client.SetHostURL(BaseURLV2)
	// Headers for all request
	client.SetHeader("Accept", "application/json")
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "dog_go_api",
	})
	return &client
}
