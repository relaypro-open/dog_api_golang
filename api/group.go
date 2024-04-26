package api

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Group struct {
	ID                  string                 `json:"id"`
	Description         string                 `json:"description"`
	Name                string                 `json:"name"`
	ProfileId           string                 `json:"profile_id"`
	ProfileName         string                 `json:"profile_name"`
	ProfileVersion      string                 `json:"profile_version"`
	Ec2SecurityGroupIds []*Ec2SecurityGroupIds `json:"ec2_security_group_ids"`
	Vars                *string                `json:"vars,omitempty"`
}

type GroupJson struct {
	ID                  string                 `json:"id,omitempty"`
	Description         string                 `json:"description,omitempty"`
	Name                string                 `json:"name,omitempty"`
	ProfileId           string                 `json:"profile_id,omitempty"`
	ProfileName         string                 `json:"profile_name,omitempty"`
	ProfileVersion      string                 `json:"profile_version,omitempty"`
	Ec2SecurityGroupIds []*Ec2SecurityGroupIds `json:"ec2_security_group_ids,omitempty"`
	Vars                map[string]any         `json:"vars,omitempty"` //parsed json
}

type GroupListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type Ec2SecurityGroupIds struct {
	Region string `json:"region"`
	SgId   string `json:"sgid"`
}

// GroupUpdateRequest is a struct for the request object required to update a Group
type GroupUpdateRequest struct {
	Description         string                 `json:"description,omitempty"`
	Name                string                 `json:"name,omitempty"`
	ProfileId           string                 `json:"profile_id,omitempty"`
	ProfileName         string                 `json:"profile_name,omitempty"`
	ProfileVersion      string                 `json:"profile_version,omitempty"`
	Ec2SecurityGroupIds []*Ec2SecurityGroupIds `json:"ec2_security_group_ids"`
	Vars                *string                 `json:"vars,omitempty"`
}

type GroupCreateRequest struct {
	Description         string                 `json:"description,omitempty"`
	Name                string                 `json:"name"`
	ProfileId           string                 `json:"profile_id,omitempty"`
	ProfileName         string                 `json:"profile_name,omitempty"`
	ProfileVersion      string                 `json:"profile_version,omitempty"`
	Ec2SecurityGroupIds []*Ec2SecurityGroupIds `json:"ec2_security_group_ids"`
	Vars                *string                 `json:"vars"`
}

type GroupCreateResponse struct {
	ID     string `json:"id"`
	Result string `json:"result"`
}

type GroupsList []Group

type GroupsListJson []GroupJson

type GroupsListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func encodeGroup(groupJson GroupJson) (group Group, marshalErr error) {
	var responseVars []byte
	if groupJson.Vars != nil {
		responseVars, marshalErr = json.Marshal(groupJson.Vars)
		varsString := string(responseVars)
		group.Vars = &varsString
	}
	group.ID = groupJson.ID
	group.Description = groupJson.Description
	group.ProfileId = groupJson.ProfileId
	group.ProfileName = groupJson.ProfileName
	group.ProfileVersion = groupJson.ProfileVersion
	group.Name = groupJson.Name
	group.Ec2SecurityGroupIds = groupJson.Ec2SecurityGroupIds
	return group, marshalErr
}

func decodeGroup(group Group) (groupJson GroupJson, unmarshalErr error) {
	if group.Vars != nil {
		var vars = map[string]any{}
		unmarshalErr = json.Unmarshal([]byte(*group.Vars), &vars)
		groupJson.Vars = map[string]any(vars)
	}
	groupJson.ID = group.ID
	groupJson.Description = group.Description
	groupJson.ProfileId = group.ProfileId
	groupJson.ProfileName = group.ProfileName
	groupJson.Name = group.Name
	groupJson.Ec2SecurityGroupIds = group.Ec2SecurityGroupIds
	return groupJson, unmarshalErr
}

func (c *Client) GetGroups(options *GroupsListOptions) (groupsList GroupsListJson, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&GroupsListJson{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		Get("/groups")

	result := (*resp.Result().(*GroupsListJson))
	return result, resp.StatusCode(), err
}

func (c *Client) GetGroupsEncode(options *GroupsListOptions) (groupsList GroupsList, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&GroupsListJson{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		Get("/groups")

	result := (*resp.Result().(*GroupsListJson))
	encodeGroups := GroupsList{}
	for _, groupJson := range result {
		ec, _ := encodeGroup(groupJson)
		encodeGroups = append(encodeGroups, ec)
	}
	return encodeGroups, resp.StatusCode(), err
}

func (c *Client) GetGroup(GroupID string, options *GroupListOptions) (group GroupJson, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&GroupJson{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"GroupID": GroupID,
		}).
		Get("/group/{GroupID}")

	result := (*resp.Result().(*GroupJson))
	return result, resp.StatusCode(), err

}

func (c *Client) GetGroupEncode(GroupID string, options *GroupListOptions) (group Group, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, respErr := c.Client.R().
		SetResult(&GroupJson{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"GroupID": GroupID,
		}).
		Get("/group/{GroupID}")

	result := (*resp.Result().(*GroupJson))
	group, responseVarsErr := encodeGroup(result)
	err := errors.Join(respErr, responseVarsErr)
	return group, resp.StatusCode(), err
}

func (c *Client) UpdateGroup(GroupID string, GroupUpdate GroupJson, options *GroupListOptions) (group GroupJson, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&GroupJson{}).
		SetPathParams(map[string]string{
			"GroupID": GroupID,
		}).
		SetBody(GroupUpdate).
		Put("/group/{GroupID}")

	result := (*resp.Result().(*GroupJson))
	return result, resp.StatusCode(), err
}

func (c *Client) UpdateGroupEncode(GroupID string, groupUpdate Group, options *GroupListOptions) (group Group, statusCode int, Error error) {
	requestGroup, responseVarsErr := decodeGroup(groupUpdate)

	resp, respErr := c.Client.R().
		SetResult(&GroupJson{}).
		SetPathParams(map[string]string{
			"GroupID": GroupID,
		}).
		SetBody(requestGroup).
		Put("/group/{GroupID}")
	result := (*resp.Result().(*GroupJson))
	group, responseVarsErr = encodeGroup(result)

	err := errors.Join(respErr, responseVarsErr)
	return group, resp.StatusCode(), err
}

func (c *Client) CreateGroup(groupNew GroupJson, options *GroupListOptions) (group GroupJson, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&GroupJson{}).
		SetBody(groupNew).
		Post("/group")

	result := (*resp.Result().(*GroupJson))
	return result, resp.StatusCode(), err
}

func (c *Client) CreateGroupEncode(groupNew Group, options *GroupListOptions) (group Group, statusCode int, Error error) {

	requestGroup, responseVarsErr := decodeGroup(groupNew)

	resp, respErr := c.Client.R().
		SetResult(&GroupJson{}).
		SetBody(requestGroup).
		Post("/group")

	result := (*resp.Result().(*GroupJson))
	group, responseVarsErr = encodeGroup(result)
	err := errors.Join(respErr, responseVarsErr)
	return group, resp.StatusCode(), err
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
