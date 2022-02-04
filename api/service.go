package api

import (
	"strconv"
)

type Service struct {
	Created  int        `json:"created"`
	ID       string     `json:"id"`
	Services []Services `json:"services"`
	Name     string     `json:"name"`
	Version  int        `json:"version"`
}

type Services struct {
	Ports    []string `json:"ports"`
	Protocol string   `json:"protocol"`
}

type ServiceListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type ServiceUpdateRequest struct {
	Services []Services `json:"services"`
	Name     string     `json:"name"`
	Version  int        `json:"version"`
}

type ServiceCreateRequest struct {
	Services []Services `json:"services"`
	Name     string     `json:"name"`
	Version  int        `json:"version"`
}

type ServiceCreateResponse struct {
	ID     string `json:"id"`
	Result string `json:"result"`
}

func (c *Client) GetService(ServiceID string, options *ServiceListOptions) (service Service, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Service{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"ServiceID": ServiceID,
		}).
		Get("/service/{ServiceID}")

	result := (*resp.Result().(*Service))
	return result, resp.StatusCode(), err

}

func (c *Client) UpdateService(ServiceID string, ServiceUpdate ServiceUpdateRequest, options *ServiceListOptions) (service Service, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Service{}).
		SetPathParams(map[string]string{
			"ServiceID": ServiceID,
		}).
		SetBody(ServiceUpdate).
		Put("/service/{ServiceID}")

	result := (*resp.Result().(*Service))
	return result, resp.StatusCode(), err
}

func (c *Client) CreateService(serviceNew ServiceCreateRequest, options *ServiceListOptions) (serviceCreateResponse ServiceCreateResponse, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&ServiceCreateResponse{}).
		SetBody(serviceNew).
		Post("/service")

	result := (*resp.Result().(*ServiceCreateResponse))
	return result, resp.StatusCode(), err
}

func (c *Client) DeleteService(ServiceID string, options *ServiceListOptions) (service Service, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Service{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"ServiceID": ServiceID,
		}).
		Delete("/service/{ServiceID}")

	result := (*resp.Result().(*Service))
	return result, resp.StatusCode(), err
}
