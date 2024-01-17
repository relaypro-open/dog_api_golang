package api

import (
	"strconv"
//	"encoding/json"
)

type FactsList []Fact

type FactsListJson []FactJson

type FactsListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type Fact struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Groups        map[string]*FactGroup `json:"groups"`
}

type FactJson struct {
	ID           string   `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	Groups        map[string]*FactGroupJson `json:"groups,omitempty"`
}

type FactGroup struct {
	Vars          string   `json:"vars"`
	Hosts         map[string]map[string]string `json:"hosts"`
	Children      []string `json:"children"`
}

type FactGroupJson struct {
	Vars          map[string]any   `json:"vars,omitempty"`
	Hosts         map[string]map[string]string `json:"hosts,omitempty"`
	Children      []string `json:"children,omitempty"`
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

func (c *Client) GetFacts(options *FactsListOptions) (factsList FactsListJson, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&FactsListJson{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		Get("/facts")

	result := (*resp.Result().(*FactsListJson))
	//for _, factJson := range result {
	//	groups := factJson.Groups
	//	for groupName, group := range groups {
	//		PrettyPrint(groupName, group)
	//		vars := group.Vars
	//		PrettyPrint("vars", vars)
	//	}
	//	//var vars = map[string]any{}
	//	//unmarshalErr := json.Unmarshal([]byte(groups.Vars), &vars)
	//	//groups.Vars = groupsJson.Vars
	//	//var vars = map[string]any{}
	//	//json.Unmarshal([]byte(groupsJson.Vars), &vars)
	//}
	return result, resp.StatusCode(), err
}

func (c *Client) GetFact(FactID string, options *FactListOptions) (fact FactJson, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&FactJson{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"FactID": FactID,
		}).
		Get("/fact/{FactID}")

	result := (*resp.Result().(*FactJson))
	return result, resp.StatusCode(), err

}

func (c *Client) GetFactByName(FactName string, options *FactListOptions) (fact FactJson, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&FactJson{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"FactName": FactName,
		}).
		Get("/fact?name={FactName}")

	result := (*resp.Result().(*FactJson))
	return result, resp.StatusCode(), err

}

func (c *Client) UpdateFact(FactID string, FactUpdate FactJson, options *FactListOptions) (fact FactJson, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&FactJson{}).
		SetPathParams(map[string]string{
			"FactID": FactID,
		}).
		SetBody(FactUpdate).
		Put("/fact/{FactID}")

	result := (*resp.Result().(*FactJson))
	return result, resp.StatusCode(), err
}

func (c *Client) CreateFact(factNew FactJson, options *FactListOptions) (fact FactJson, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&FactJson{}).
		SetBody(factNew).
		Post("/fact")

	result := (*resp.Result().(*FactJson))
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
