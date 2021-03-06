package api

import (
	"strconv"
)

type Group struct {
	ID string `json:"id"`
	//Created        int    `json:"created,omitempty"` //TODO: created has both int and string entries
	Description    string `json:"description"`
	Name           string `json:"name"`
	ProfileName    string `json:"profile_name"`
	ProfileVersion string `json:"profile_version"`
}

type GroupListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

// GroupUpdateRequest is a struct for the request object required to update a Group
type GroupUpdateRequest struct {
	Description    string `json:"description,omitempty"`
	Name           string `json:"name,omitempty"`
	ProfileName    string `json:"profile_name,omitempty"`
	ProfileVersion string `json:"profile_version,omitempty"`
}

type GroupCreateRequest struct {
	Description    string `json:"description,omitempty"`
	Name           string `json:"name"`
	ProfileName    string `json:"profile_name,omitempty"`
	ProfileVersion string `json:"profile_version,omitempty"`
}

type GroupCreateResponse struct {
	ID     string `json:"id"`
	Result string `json:"result"`
}

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

func (c *Client) GetGroup(GroupID string, options *GroupListOptions) (group Group, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Group{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"GroupID": GroupID,
		}).
		Get("/group/{GroupID}")

	result := (*resp.Result().(*Group))
	return result, resp.StatusCode(), err

}

func (c *Client) UpdateGroup(GroupID string, GroupUpdate GroupUpdateRequest, options *GroupListOptions) (group Group, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Group{}).
		SetPathParams(map[string]string{
			"GroupID": GroupID,
		}).
		SetBody(GroupUpdate).
		Put("/group/{GroupID}")

	result := (*resp.Result().(*Group))
	return result, resp.StatusCode(), err
}

func (c *Client) CreateGroup(groupNew GroupCreateRequest, options *GroupListOptions) (group Group, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Group{}).
		SetBody(groupNew).
		Post("/group")

	result := (*resp.Result().(*Group))
	return result, resp.StatusCode(), err
}

func (c *Client) DeleteGroup(GroupID string, options *GroupListOptions) (group Group, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Group{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"GroupID": GroupID,
		}).
		Delete("/group/{GroupID}")

	result := (*resp.Result().(*Group))
	return result, resp.StatusCode(), err
}
