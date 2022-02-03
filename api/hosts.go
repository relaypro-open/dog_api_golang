package api

import "strconv"

type HostsList struct {
	//Count      int    `json:"count"`
	//PagesCount int    `json:"pages_count"`
	Hosts []Host `json:"hosts"`
}

type HostsListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func (c *Client) GetHosts(options *HostsListOptions) (HostsList, error) {
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
	return result, err
}
