package api

import (
	"strconv"
)

type Host struct {
	Active      string `json:"active"`
	Environment string `json:"environment"`
	Group       string `json:"group"`
	Id          string `json:"id"`
	HostKey     string `json:"hostkey"`
	Location    string `json:"location"`
	Name        string `json:"name"`
}

type HostListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

// HostUpdateRequest is a struct for the request object required to update a Host
type HostUpdateRequest struct {
	Active      string `json:"active,omitempty"`
	Environment string `json:"environment,omitempty"`
	Group       string `json:"group,omitempty"`
	HostKey     string `json:"hostkey,omitempty"`
	Location    string `json:"location,omitempty"`
	Name        string `json:"name,omitempty"`
}

type HostCreateRequest struct {
	Active      string `json:"active"`
	Environment string `json:"environment"`
	Group       string `json:"group"`
	HostKey     string `json:"hostkey"`
	Location    string `json:"location"`
	Name        string `json:"name"`
}

func (c *Client) GetHost(HostId string, options *HostListOptions) (Host, int, error) {
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
	return result, resp.StatusCode(), err

}

func (c *Client) UpdateHost(HostId string, hostUpdate HostUpdateRequest, options *HostListOptions) (Host, int, error) {

	resp, err := c.Client.R().
		SetResult(&Host{}).
		SetPathParams(map[string]string{
			"HostId": HostId,
		}).
		SetBody(hostUpdate).
		Put("/host/{HostId}")

	result := (*resp.Result().(*Host))
	return result, resp.StatusCode(), err
}

func (c *Client) CreateHost(hostNew HostCreateRequest, options *HostListOptions) (Host, int, error) {

	resp, err := c.Client.R().
		SetResult(&Host{}).
		SetBody(hostNew).
		Post("/host")

	result := (*resp.Result().(*Host))
	return result, resp.StatusCode(), err
}

func (c *Client) DeleteHost(HostId string, options *HostListOptions) (Host, int, error) {
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
		Delete("/host/{HostId}")

	result := (*resp.Result().(*Host))
	return result, resp.StatusCode(), err

}
