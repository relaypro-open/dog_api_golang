package api

import "strconv"

type HostsList []Host

type HostsListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func (c *Client) GetHosts(options *HostsListOptions) (hostsList HostsList, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&HostsList{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		Get("/hosts")

	result := (*resp.Result().(*HostsList))
	return result, resp.StatusCode(), err
}
