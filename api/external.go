package api

import (
	"strconv"
)

type ExternalList []External

type ExternalListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type External struct {
	AddressHandling string `json:"address_handling"`
	Enabled         bool   `json:"enabled"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	State           string `json:"state"`
	Timestamp       int    `json:"timestamp"`
	V4              V4     `json:"v4"`
	V6              V6     `json:"v6"`
}

type V4 struct {
	Groups map[string]map[string][]string `json:"groups"`
	Zones  map[string]map[string][]string `json:"zones"`
}

type V6 struct {
	Groups map[string]map[string][]string `json:"groups"`
	Zones  map[string]map[string][]string `json:"zones"`
}

func (c *Client) GetExternals(options *ExternalListOptions) (externalList ExternalList, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&ExternalList{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		Get("/externals")

	result := (*resp.Result().(*ExternalList))
	return result, resp.StatusCode(), err
}

func (c *Client) GetExternal(ExternalID string, options *ExternalListOptions) (external External, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&External{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"ExternalID": ExternalID,
		}).
		Get("/external/{ExternalID}")

	result := (*resp.Result().(*External))
	return result, resp.StatusCode(), err

}
