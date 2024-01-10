package api

import (
	"strconv"
)

type FactsList []Fact

type FactsListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type Fact struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Groups        map[string]*FactGroup `json:"groups"`
}

type FactGroup struct {
	Vars          map[string]any `json:"vars"`
	Hosts         map[string]map[string]string `json:"hosts"`
	Children      []string `json:"children"`
}

type FactListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

// FactUpdateRequest is a struct for the request object required to update an Fact
type FactUpdateRequest struct {
	Name          string   `json:"name"`
	Groups        map[string]*FactGroup `json:"groups"`
}

type FactCreateRequest struct {
	Name          string   `json:"name"`
	Groups        map[string]*FactGroup `json:"groups"`
}

type FactCreateResponse struct {
	ID     string `json:"id"`
	Result string `json:"result"`
}

func (c *Client) GetFacts(options *FactsListOptions) (factsList FactsList, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&FactsList{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		Get("/facts")

	result := (*resp.Result().(*FactsList))
	return result, resp.StatusCode(), err
}

func (c *Client) GetFact(FactID string, options *FactListOptions) (fact Fact, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Fact{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"FactID": FactID,
		}).
		Get("/fact/{FactID}")

	result := (*resp.Result().(*Fact))
	return result, resp.StatusCode(), err

}

func (c *Client) GetFactByName(FactName string, options *FactListOptions) (fact Fact, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Fact{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"FactName": FactName,
		}).
		Get("/fact?name={FactName}")

	result := (*resp.Result().(*Fact))
	return result, resp.StatusCode(), err

}

func (c *Client) UpdateFact(FactID string, FactUpdate FactUpdateRequest, options *FactListOptions) (fact Fact, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Fact{}).
		SetPathParams(map[string]string{
			"FactID": FactID,
		}).
		SetBody(FactUpdate).
		Put("/fact/{FactID}")

	result := (*resp.Result().(*Fact))
	return result, resp.StatusCode(), err
}

func (c *Client) CreateFact(factNew FactCreateRequest, options *FactListOptions) (fact Fact, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Fact{}).
		SetBody(factNew).
		Post("/fact")

	result := (*resp.Result().(*Fact))
	return result, resp.StatusCode(), err
}

func (c *Client) DeleteFact(FactID string, options *FactListOptions) (fact Fact, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Fact{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"FactID": FactID,
		}).
		Delete("/fact/{FactID}")

	result := (*resp.Result().(*Fact))
	return result, resp.StatusCode(), err
}
