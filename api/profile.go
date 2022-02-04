package api

import (
	"strconv"
)

type Profile struct {
	ID          string `json:"id"`
	Created     int    `json:"created"`
	Description string `json:"description"`
	Rules       Rules  `json:"rules"`
	Name        string `json:"name"`
	Version     string `json:"version"`
}

type Rules struct {
	Inbound  []Rule `json:"inbound"`
	Outbound []Rule `json:"outbound"`
}

type Rule struct {
	Action      string   `json:"action"`
	Active      bool     `json:"active"`
	Comment     string   `json:"comment"`
	Environment []string `json:"environments"`
	Group       string   `json:"group"`
	GroupType   string   `json:"group_type"`
	Interface   string   `json:"interface"`
	Log         bool     `json:"log"`
	LogPrefix   string   `json:"log_prefix"`
	Order       int      `json:"order"`
	Service     string   `json:"service"`
	States      []string `json:"states"`
	Type        string   `json:"type"`
}

type ProfileListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

// ProfileUpdateRequest is a struct for the request object required to update a Profile
type ProfileUpdateRequest struct {
	Description string `json:"description,omitempty"`
	Rules       Rules  `json:"rules,omitempty"`
	Name        string `json:"name,omitempty"`
	Version     string `json:"version,omitempty"`
}

type ProfileCreateRequest struct {
	Description string `json:"description,omitempty"`
	Rules       Rules  `json:"rules,omitempty"`
	Name        string `json:"name"`
	Version     string `json:"version,omitempty"`
}

type ProfileCreateResponse struct {
	ID     string `json:"id"`
	Result string `json:"result"`
}

func (c *Client) GetProfile(ProfileID string, options *ProfileListOptions) (profile Profile, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Profile{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"ProfileID": ProfileID,
		}).
		Get("/profile/{ProfileID}")

	result := (*resp.Result().(*Profile))
	return result, resp.StatusCode(), err

}

func (c *Client) UpdateProfile(ProfileID string, ProfileUpdate ProfileUpdateRequest, options *ProfileListOptions) (profile Profile, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Profile{}).
		SetPathParams(map[string]string{
			"ProfileID": ProfileID,
		}).
		SetBody(ProfileUpdate).
		Put("/profile/{ProfileID}")

	result := (*resp.Result().(*Profile))
	return result, resp.StatusCode(), err
}

func (c *Client) CreateProfile(profileNew ProfileCreateRequest, options *ProfileListOptions) (profileCreateResponse ProfileCreateResponse, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&ProfileCreateResponse{}).
		SetBody(profileNew).
		Post("/profile")

	result := (*resp.Result().(*ProfileCreateResponse))
	return result, resp.StatusCode(), err
}

func (c *Client) DeleteProfile(ProfileID string, options *ProfileListOptions) (profile Profile, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Profile{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"ProfileID": ProfileID,
		}).
		Delete("/profile/{ProfileID}")

	result := (*resp.Result().(*Profile))
	return result, resp.StatusCode(), err
}