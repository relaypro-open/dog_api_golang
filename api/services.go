package api

import "strconv"

type ServicesList []Service

type ServicesListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func (c *Client) GetServices(options *ServicesListOptions) (servicesList ServicesList, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&ServicesList{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		Get("/services")

	result := (*resp.Result().(*ServicesList))
	return result, resp.StatusCode(), err
}
