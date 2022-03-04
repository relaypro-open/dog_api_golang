package api

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	*resty.Client
}

func NewClient(apiKey string, apiEndpoint string) *Client {
	client := Client{}
	client.Client = resty.New()
	client.SetHeader("apikey", apiKey)
	client.SetHostURL(apiEndpoint)
	// Headers for all request
	client.SetHeader("Accept", "application/json")
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "dog_go_api",
	})
	return &client
}
