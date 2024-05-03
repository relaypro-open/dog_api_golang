package api

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strconv"
)

type FactsList []Fact

type FactsListJson []FactJson

type FactsListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type Fact struct {
	ID     string                `json:"id"`
	Name   string                `json:"name"`
	Groups map[string]*FactGroup `json:"groups"`
}

type FactJson struct {
	ID     string                    `json:"id,omitempty"`
	Name   string                    `json:"name"`
	Groups map[string]*FactGroupJson `json:"groups"`
}

type FactGroup struct {
	Vars     *string                      `json:"vars,omitempty"`
	Hosts    *string                      `json:"hosts"`
	Children []string                     `json:"children"`
}

type FactGroupJson struct {
	Vars     map[string]any               `json:"vars,omitempty"`
	Hosts    map[string]map[string]any    `json:"hosts"`
	Children []string                     `json:"children"`
}

type FactListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

// FactUpdateRequest is a struct for the request object required to update an Fact
type FactUpdateRequest struct {
	Name   string                `json:"name"`
	Groups map[string]*FactGroup `json:"groups"`
}

type FactCreateRequest struct {
	Name   string                `json:"name"`
	Groups map[string]*FactGroup `json:"groups"`
}

type FactCreateResponse struct {
	ID     string `json:"id"`
	Result string `json:"result"`
}

func EncodeFact(factJson FactJson) (fact Fact) {
	factEncoded := Fact{}
	encodedGroups := map[string]*FactGroup{}
	for name, group := range factJson.Groups {
		responseVars, _ := json.Marshal(group.Hosts)
		hostsString := string(responseVars)
		if group.Vars == nil {
			encodedGroup := FactGroup{
				Hosts:    &hostsString,
				Children: group.Children,
			}
			encodedGroups[name] = &encodedGroup
		} else {
			responseVars, _ := json.Marshal(group.Vars)
			varsString := string(responseVars)
			encodedGroup := FactGroup{
				Vars:     &varsString,
				Hosts:    &hostsString,
				Children: group.Children,
			}
			encodedGroups[name] = &encodedGroup
		}
	}
	factEncoded.Groups = encodedGroups
	factEncoded.Name = factJson.Name
	factEncoded.ID = factJson.ID
	return factEncoded
}

func DecodeFact(fact Fact) (factJson FactJson, err error) {
	newGroups := map[string]*FactGroupJson{}
	var varsUnmarshallErr error
	var hostsUnmarshallErr error 
	for name, group := range fact.Groups {
		newGroup := FactGroupJson{}
		if group.Vars != nil {
			var vars = map[string]any{}
			varsUnmarshallErr = json.Unmarshal([]byte(*group.Vars), &vars)
			newGroup.Vars = vars
		}
		if group.Hosts != nil {
			var vars = map[string]map[string]any{}
			hostsUnmarshallErr = json.Unmarshal([]byte(*group.Hosts), &vars)
			newGroup.Hosts = vars
		}
		newGroup.Children = group.Children
		newGroups[name] = &newGroup
	}
	factDecoded := FactJson{}
	factDecoded.Groups = newGroups
	factDecoded.Name = fact.Name
	factDecoded.ID = fact.ID
	err = errors.Join(varsUnmarshallErr, hostsUnmarshallErr)

	return factDecoded, err
}

func (c *Client) GetFactsEncode(options *FactsListOptions) (factsList FactsList, statusCode int, Error error) {
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
	EncodeFacts := FactsList{}
	for _, factJson := range result {
		EncodeFacts = append(EncodeFacts, EncodeFact(factJson))
	}
	return EncodeFacts, resp.StatusCode(), err
}

func (c *Client) GetFactEncode(FactID string, options *FactListOptions) (fact Fact, statusCode int, Error error) {
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
	factEncoded := EncodeFact(result)
	return factEncoded, resp.StatusCode(), err

}

func (c *Client) GetFactByNameEncode(FactName string, options *FactListOptions) (fact Fact, statusCode int, Error error) {
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
	factEncoded := EncodeFact(result)
	return factEncoded, resp.StatusCode(), err

}

func (c *Client) UpdateFactEncode(FactID string, factUpdate Fact, options *FactListOptions) (fact Fact, statusCode int, Error error) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	factDecoded, decodeErr := DecodeFact(factUpdate)
	resp, clientErr := c.Client.R().
		SetResult(&FactJson{}).
		SetPathParams(map[string]string{
			"FactID": FactID,
		}).
		SetBody(factDecoded).
		Put("/fact/{FactID}")

	result := (*resp.Result().(*FactJson))
	factEncoded := EncodeFact(result)
	err = errors.Join(decodeErr, clientErr)

	return factEncoded, resp.StatusCode(), err
}

func (c *Client) CreateFactEncode(factNew Fact, options *FactListOptions) (fact Fact, statusCode int, Error error) {

	PrettyPrint("factNew", factNew)
	factDecoded, decodeErr := DecodeFact(factNew)
	PrettyPrint("factDecoded", factDecoded)
	resp, respErr := c.Client.R().
		SetResult(&FactJson{}).
		SetBody(factDecoded).
		Post("/fact")

	result := (*resp.Result().(*FactJson))
	PrettyPrint("fact result", result)
	factEncoded := EncodeFact(result)
	PrettyPrint("factEncoded", factEncoded)
	err := errors.Join(decodeErr, respErr)

	return factEncoded, resp.StatusCode(), err
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
