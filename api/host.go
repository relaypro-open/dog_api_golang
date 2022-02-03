package api

import "strconv"

type Host struct {
	Active  string `json:"active"`
	Group   string `json:"group"`
	HostKey string `json:"hostkey"`
	Id      string `json:"id"`
	Name    string `json:"name"`
}

type HostListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

// HostUpdateRequest is a struct for the request object required to update a Host
type HostUpdateRequest struct {
	Group   string `json:"group"`
	HostKey string `json:"hostkey"`
	Name    string `json:"name"`
}

func (c *Client) GetHost(HostId string, options *HostListOptions) (Host, error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Host{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"HostId": HostId,
		}).
		Get("/host/{HostId}")

	result := (*resp.Result().(*Host))
	return result, err
}

func (c *Client) UpdateHost(HostId string, hostUpdate HostUpdateRequest, options *HostListOptions) (Host, error) {

	resp, err := c.Client.R().
		SetResult(&Host{}).
		SetPathParams(map[string]string{
			"HostId": HostId,
		}).
		SetBody(hostUpdate).
		Put("/host/{HostId}")

	result := (*resp.Result().(*Host))
	return result, err
}
