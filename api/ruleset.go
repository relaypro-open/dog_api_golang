package api

import (
	"strconv"
)

type Ruleset struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Rules       *Rules `json:"rules"`
}

type RulesetList []Ruleset

type RulesetUpdateRequest struct {
	Name        string `json:"name"`
	Rules       *Rules `json:"rules,omitempty"`
}

type RulesetCreateRequest struct {
	Name        string `json:"name"`
	Rules       *Rules `json:"rules,omitempty"`
}

type RulesList []Rule

type RulesListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type RuleListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type RulesetCreateResponse struct {
	ID     string `json:"id"`
	Result string `json:"result"`
}

type Rules struct {
	Inbound  []*Rule `json:"inbound"`
	Outbound []*Rule `json:"outbound"`
}

type Rule struct {
	Action       string   `json:"action"`
	Active       bool     `json:"active"`
	Comment      string   `json:"comment"`
	Environments []string `json:"environments"`
	Group        string   `json:"group"`
	GroupType    string   `json:"group_type"`
	Interface    string   `json:"interface"`
	Log          bool     `json:"log"`
	LogPrefix    string   `json:"log_prefix"`
	Order        int      `json:"order"`
	Service      string   `json:"service"`
	States       []string `json:"states"`
	Type         string   `json:"type"`
}

func (c *Client) GetRulesets(options *RulesListOptions) (rulesetList RulesetList, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&RulesetList{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		Get("/rulesets")

	result := (*resp.Result().(*RulesetList))
	return result, resp.StatusCode(), err
}

func (c *Client) GetRuleset(rulesetId string, options *RuleListOptions) (ruleset Ruleset, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Ruleset{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"rulesetID": rulesetId,
		}).
		Get("/ruleset/{rulesetID}")

	result := (*resp.Result().(*Ruleset))
	return result, resp.StatusCode(), err

}

func (c *Client) GetRulesetByName(rulesetName string, options *RuleListOptions) (ruleset Ruleset, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Ruleset{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"rulesetName": rulesetName,
		}).
		Get("/ruleset?name={rulesetName}")

	result := (*resp.Result().(*Ruleset))
	return result, resp.StatusCode(), err

}

func (c *Client) UpdateRuleset(rulesetId string, rulesetUpdate RulesetUpdateRequest, options *RuleListOptions) (ruleset Ruleset, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Ruleset{}).
		SetPathParams(map[string]string{
			"rulesetId": rulesetId,
		}).
		SetBody(rulesetUpdate).
		Put("/ruleset/{rulesetId}")

	result := (*resp.Result().(*Ruleset))
	return result, resp.StatusCode(), err
}

func (c *Client) CreateRuleset(rulesetNew RulesetCreateRequest, options *RuleListOptions) (ruleset Ruleset, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Ruleset{}).
		SetBody(rulesetNew).
		Post("/ruleset")

	result := (*resp.Result().(*Ruleset))
	return result, resp.StatusCode(), err
}

func (c *Client) DeleteRuleset(rulesetId string, options *RuleListOptions) (ruleset Ruleset, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Ruleset{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"rulesetId": rulesetId,
		}).
		Delete("/ruleset/{rulesetId}")

	result := (*resp.Result().(*Ruleset))
	return result, resp.StatusCode(), err
}
