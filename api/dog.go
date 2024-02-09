package api
	

import (
	"github.com/go-resty/resty/v2"
    _ "embed"
)

//go:generate sh -c "printf %s $(git rev-parse HEAD) > VERSION.txt"
//go:embed VERSION.txt
var Commit string

type Client struct {
	*resty.Client
}

func NewClient(apiToken string, apiEndpoint string) *Client {
	client := Client{}
	client.Client = resty.New()
	client.SetAuthToken(apiToken)
	client.SetBaseURL(apiEndpoint)
	// Headers for all request
	client.SetHeader("Accept", "application/json")
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "dog_go_api_" + Commit,
	})
	return &client
}
