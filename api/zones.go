package api

import "strconv"

type ZonesList []Zone

type ZonesListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
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
