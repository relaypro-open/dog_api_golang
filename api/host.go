package api

import (
	"strconv"
)

type Host struct {
	Environment string `json:"environment"`
	Group       string `json:"group"`
	ID          string `json:"id"`
	HostKey     string `json:"hostkey"`
	Location    string `json:"location"`
	Name        string `json:"name"`
	Vars	    map[string]any `json:"vars"`
}

type HostListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

// HostUpdateRequest is a struct for the request object required to update a Host
type HostUpdateRequest struct {
	Environment string `json:"environment,omitempty"`
	Group       string `json:"group,omitempty"`
	ID          string `json:"id,omitempty"`
	HostKey     string `json:"hostkey,omitempty"`
	Location    string `json:"location,omitempty"`
	Name        string `json:"name,omitempty"`
	Vars	    map[string]any `json:"vars"`
}

type HostCreateRequest struct {
	Environment string `json:"environment"`
	Group       string `json:"group"`
	HostKey     string `json:"hostkey"`
	Location    string `json:"location"`
	Name        string `json:"name"`
	Vars	    map[string]any `json:"vars"`
}

type HostCreateResponse struct {
	Environment string `json:"environment"`
	Group       string `json:"group"`
	ID          string `json:"id"`
	HostKey     string `json:"hostkey"`
	Location    string `json:"location"`
	Name        string `json:"name"`
	Vars	    map[string]any `json:"vars"`
}

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

func (c *Client) GetHost(hostID string, options *HostListOptions) (host Host, statusCode int, Error error) {
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
			"hostID": hostID,
		}).
		Get("/host/{hostID}")

	result := (*resp.Result().(*Host))
	return result, resp.StatusCode(), err

}

func (c *Client) UpdateHost(hostID string, hostUpdate HostUpdateRequest, options *HostListOptions) (host Host, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Host{}).
		SetPathParams(map[string]string{
			"hostID": hostID,
		}).
		SetBody(hostUpdate).
		Put("/host/{hostID}")

	result := (*resp.Result().(*Host))
	return result, resp.StatusCode(), err
}

func (c *Client) CreateHost(hostNew HostCreateRequest, options *HostListOptions) (host Host, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Host{}).
		SetBody(hostNew).
		Post("/host")

	result := (*resp.Result().(*Host))
	return result, resp.StatusCode(), err
}

func (c *Client) DeleteHost(hostID string, options *HostListOptions) (host Host, statusCode int, Error error) {
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
			"hostID": hostID,
		}).
		Delete("/host/{hostID}")

	result := (*resp.Result().(*Host))
	return result, resp.StatusCode(), err

}
