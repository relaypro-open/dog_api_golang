package api

import (
	"strconv"
)

type ZonesList []Zone

type ZonesListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type Zone struct {
	Created       string   `json:"created"`
	ID            string   `json:"id"`
	IPv4Addresses []string `json:"ipv4_addresses"`
	IPv6Addresses []string `json:"ipv6_addresses"`
	Name          string   `json:"name"`
}

type ZoneListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

// ZoneUpdateRequest is a struct for the request object required to update a Zone
type ZoneUpdateRequest struct {
	IPv4Addresses []string `json:"ipv4_addresses"`
	IPv6Addresses []string `json:"ipv6_addresses"`
	Name          string   `json:"name"`
}

type ZoneCreateRequest struct {
	IPv4Addresses []string `json:"ipv4_addresses"`
	IPv6Addresses []string `json:"ipv6_addresses"`
	Name          string   `json:"name"`
}

type ZoneCreateResponse struct {
	ID     string `json:"id"`
	Result string `json:"result"`
}

func (c *Client) GetZones(options *ZonesListOptions) (zonesList ZonesList, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&ZonesList{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		Get("/zones")

	result := (*resp.Result().(*ZonesList))
	return result, resp.StatusCode(), err
}

func (c *Client) GetZone(ZoneID string, options *ZoneListOptions) (zone Zone, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Zone{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"ZoneID": ZoneID,
		}).
		Get("/zone/{ZoneID}")

	result := (*resp.Result().(*Zone))
	return result, resp.StatusCode(), err

}

func (c *Client) GetZoneByName(ZoneName string, options *ZoneListOptions) (zone Zone, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Zone{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"ZoneName": ZoneName,
		}).
		Get("/zone?name={ZoneName}")

	result := (*resp.Result().(*Zone))
	return result, resp.StatusCode(), err

}

func (c *Client) UpdateZone(ZoneID string, ZoneUpdate ZoneUpdateRequest, options *ZoneListOptions) (zone Zone, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Zone{}).
		SetPathParams(map[string]string{
			"ZoneID": ZoneID,
		}).
		SetBody(ZoneUpdate).
		Put("/zone/{ZoneID}")

	result := (*resp.Result().(*Zone))
	return result, resp.StatusCode(), err
}

func (c *Client) CreateZone(zoneNew ZoneCreateRequest, options *ZoneListOptions) (zone Zone, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Zone{}).
		SetBody(zoneNew).
		Post("/zone")

	result := (*resp.Result().(*Zone))
	return result, resp.StatusCode(), err
}

func (c *Client) DeleteZone(ZoneID string, options *ZoneListOptions) (zone Zone, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Zone{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"ZoneID": ZoneID,
		}).
		Delete("/zone/{ZoneID}")

	result := (*resp.Result().(*Zone))
	return result, resp.StatusCode(), err
}
