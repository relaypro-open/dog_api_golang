package api

import "strconv"

type ProfilesList []Profile

type ProfilesListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func (c *Client) GetProfiles(options *ProfilesListOptions) (profilesList ProfilesList, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&ProfilesList{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		Get("/profiles")

	result := (*resp.Result().(*ProfilesList))
	return result, resp.StatusCode(), err
}
