package api

import (
	"context"
	"fmt"
	"net/http"
)

type HostsList struct {
	//Count      int    `json:"count"`
	//PagesCount int    `json:"pages_count"`
	Hosts []Host `json:"hosts"`
}

type Host struct {
	Active  string `json:"active"`
	Group   string `json:"group"`
	HostKey string `json:"hostkey"`
	Id      string `json:"id"`
	Name    string `json:"name"`
}

type HostsListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func (c *Client) GetHosts(ctx context.Context, options *HostsListOptions) (*HostsList, error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/hosts?limit=%d&page=%d", c.BaseURL, limit, page), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := HostsList{}
	fmt.Sprintf("res: %v", res)
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
