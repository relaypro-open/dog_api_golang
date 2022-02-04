package api

import "strconv"

type GroupsList []Group

type GroupsListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func (c *Client) GetGroups(options *GroupsListOptions) (groupsList GroupsList, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&GroupsList{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		Get("/groups")

	result := (*resp.Result().(*GroupsList))
	return result, resp.StatusCode(), err
}
