package api

import "strconv"

type LinksList []Link

type LinksListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func (c *Client) GetLinks(options *LinksListOptions) (linksList LinksList, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&LinksList{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		Get("/links")

	result := (*resp.Result().(*LinksList))
	return result, resp.StatusCode(), err
}
