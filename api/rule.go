package api

import (
	"strconv"
)

type Rule struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Rules       *FwRules `json:"rules"`
}

type RuleUpdateRequest struct {
	Name        string `json:"name"`
	Rules       *FwRules `json:"rules"`
}

type RuleCreateRequest struct {
	Name        string `json:"name"`
	Rules       *FwRules `json:"rules"`
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

// ruleUpdateRequest is a struct for the request object required to update a rule
type FwRuleUpdateRequest struct {
	IPv4Addresses []string `json:"ipv4_addresses"`
	IPv6Addresses []string `json:"ipv6_addresses"`
	Name          string   `json:"name"`
}

type FwRuleCreateRequest struct {
	IPv4Addresses []string `json:"ipv4_addresses"`
	IPv6Addresses []string `json:"ipv6_addresses"`
	Name          string   `json:"name"`
}

type RuleCreateResponse struct {
	ID     string `json:"id"`
	Result string `json:"result"`
}

type FwRules struct {
	Inbound  []*FwRule `json:"inbound"`
	Outbound []*FwRule `json:"outbound"`
}

type FwRule struct {
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

func (c *Client) GetRules(options *RulesListOptions) (rulesList RulesList, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&RulesList{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		Get("/rules")

	result := (*resp.Result().(*RulesList))
	return result, resp.StatusCode(), err
}

func (c *Client) GetRule(ruleID string, options *RuleListOptions) (rule Rule, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Rule{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"ruleID": ruleID,
		}).
		Get("/rule/{ruleID}")

	result := (*resp.Result().(*Rule))
	return result, resp.StatusCode(), err

}

func (c *Client) GetRuleByName(ruleName string, options *RuleListOptions) (rule Rule, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Rule{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"ruleName": ruleName,
		}).
		Get("/rule?name={ruleName}")

	result := (*resp.Result().(*Rule))
	return result, resp.StatusCode(), err

}

func (c *Client) UpdateRule(ruleID string, ruleUpdate RuleUpdateRequest, options *RuleListOptions) (rule Rule, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Rule{}).
		SetPathParams(map[string]string{
			"ruleID": ruleID,
		}).
		SetBody(ruleUpdate).
		Put("/rule/{ruleID}")

	result := (*resp.Result().(*Rule))
	return result, resp.StatusCode(), err
}

func (c *Client) CreateRule(ruleNew RuleCreateRequest, options *RuleListOptions) (rule Rule, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Rule{}).
		SetBody(ruleNew).
		Post("/rule")

	result := (*resp.Result().(*Rule))
	return result, resp.StatusCode(), err
}

func (c *Client) DeleteRule(ruleID string, options *RuleListOptions) (rule Rule, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Rule{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"ruleID": ruleID,
		}).
		Delete("/rule/{ruleID}")

	result := (*resp.Result().(*Rule))
	return result, resp.StatusCode(), err
}
