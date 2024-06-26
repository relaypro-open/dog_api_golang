package api

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Host struct {
	Environment string `json:"environment"`
	Group       string `json:"group"`
	ID          string `json:"id"`
	HostKey     string `json:"hostkey"`
	Location    string `json:"location"`
	Name        string `json:"name"`
	Vars        string `json:"vars"` //raw json for Terraform
	AlertEnable *bool  `json:"alert_enable,omitempty"`
}

type HostJson struct {
	Environment string         `json:"environment,omitempty"`
	Group       string         `json:"group,omitempty"`
	ID          string         `json:"id,omitempty"`
	HostKey     string         `json:"hostkey,omitempty"`
	Location    string         `json:"location,omitempty"`
	Name        string         `json:"name,omitempty"`
	Vars        map[string]any `json:"vars,omitempty"` //parsed json
	AlertEnable *bool          `json:"alert_enable,omitempty"`
}

type HostListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

// HostUpdateRequest is a struct for the request object required to update a Host
type HostUpdateRequest struct {
	Environment string `json:"environment,omitempty"`
	Group       string `json:"group,omitempty"`
	ID          string `json:"id,omitempty"`
	HostKey     string `json:"hostkey,omitempty"`
	Location    string `json:"location,omitempty"`
	Name        string `json:"name,omitempty"`
	Vars        string `json:"vars,omitempty"`
	AlertEnable *bool  `json:"alert_enable,omitempty"`
}

type HostCreateRequest struct {
	Environment string `json:"environment"`
	Group       string `json:"group"`
	HostKey     string `json:"hostkey"`
	Location    string `json:"location"`
	Name        string `json:"name"`
	Vars        string `json:"vars"`
	AlertEnable *bool  `json:"alert_enable,omitempty"`
}

type HostCreateResponse struct {
	Environment string `json:"environment"`
	Group       string `json:"group"`
	ID          string `json:"id"`
	HostKey     string `json:"hostkey"`
	Location    string `json:"location"`
	Name        string `json:"name"`
	Vars        string `json:"vars"`
	AlertEnable *bool  `json:"alert_enable,omitempty"`
}

type HostsList []Host

type HostsListJson []HostJson

type HostsListOptions struct {
	Limit  int    `json:"limit,omitempty"`
	Page   int    `json:"page,omitempty"`
	Active string `json:"active,omitempty"`
}

func encodeHost(hostJson HostJson) (host Host, marshalErr error) {
	var responseVars []byte
	if hostJson.Vars != nil {
		responseVars, marshalErr = json.Marshal(hostJson.Vars)
		varsString := string(responseVars)
		host.Vars = varsString
	}
	if hostJson.AlertEnable != nil {
		host.AlertEnable = hostJson.AlertEnable
	}
	host.Environment = hostJson.Environment
	host.Group = hostJson.Group
	host.HostKey = hostJson.HostKey
	host.ID = hostJson.ID
	host.Location = hostJson.Location
	host.Name = hostJson.Name
	return host, marshalErr
}

func decodeHost(host Host) (hostJson HostJson, unmarshalErr error) {
	if host.Vars != "" {
		var vars = map[string]any{}
		unmarshalErr = json.Unmarshal([]byte(host.Vars), &vars)
		hostJson.Vars = map[string]any(vars)
	}
	hostJson.AlertEnable = host.AlertEnable //add even if nil
	hostJson.Environment = host.Environment
	hostJson.Group = host.Group
	hostJson.HostKey = host.HostKey
	hostJson.Location = host.Location
	hostJson.Name = host.Name
	hostJson.ID = host.ID
	return hostJson, unmarshalErr
}

func (c *Client) GetHosts(options *HostsListOptions) (hostsList HostsListJson, statusCode int, Error error) {
	limit := 100
	page := 1
	active := "false"
	if options != nil {
		limit = options.Limit
		page = options.Page
		active = options.Active
	}

	resp, err := c.Client.R().
		SetResult(&HostsListJson{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
			"active":  active,
		}).
		Get("/hosts")

	result := (*resp.Result().(*HostsListJson))
	return result, resp.StatusCode(), err
}

func (c *Client) GetHostsEncode(options *HostsListOptions) (hostsList HostsList, statusCode int, Error error) {
	limit := 100
	page := 1
	active := "false"
	if options != nil {
		limit = options.Limit
		page = options.Page
		active = options.Active
	}

	resp, err := c.Client.R().
		SetResult(&HostsListJson{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
			"active":  active,
		}).
		Get("/hosts")

	result := (*resp.Result().(*HostsListJson))
	encodeHosts := HostsList{}
	for _, hostJson := range result {
		ec, _ := encodeHost(hostJson)
		encodeHosts = append(encodeHosts, ec)
	}
	return encodeHosts, resp.StatusCode(), err
}

func (c *Client) GetHost(hostID string, options *HostListOptions) (host HostJson, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&HostJson{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"hostID": hostID,
		}).
		Get("/host/{hostID}")

	result := (*resp.Result().(*HostJson))
	return result, resp.StatusCode(), err

}

func (c *Client) GetHostEncode(hostID string, options *HostListOptions) (host Host, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, respErr := c.Client.R().
		SetResult(&HostJson{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"hostID": hostID,
		}).
		Get("/host/{hostID}")

	result := (*resp.Result().(*HostJson))
	host, responseVarsErr := encodeHost(result)
	err := errors.Join(respErr, responseVarsErr)
	return host, resp.StatusCode(), err
}

func (c *Client) UpdateHost(hostID string, hostUpdate HostJson, options *HostListOptions) (host HostJson, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&HostJson{}).
		SetPathParams(map[string]string{
			"hostID": hostID,
		}).
		SetBody(hostUpdate).
		Put("/host/{hostID}")

	result := (*resp.Result().(*HostJson))
	return result, resp.StatusCode(), err
}

func (c *Client) UpdateHostEncode(hostID string, hostUpdate Host, options *HostListOptions) (host Host, statusCode int, Error error) {

	requestHost, decodeErr := decodeHost(hostUpdate)
	resp, respErr := c.Client.R().
		SetResult(&HostJson{}).
		SetPathParams(map[string]string{
			"hostID": hostID,
		}).
		SetBody(requestHost).
		Put("/host/{hostID}")

	result := (*resp.Result().(*HostJson))
	host, responseVarsErr := encodeHost(result)

	err := errors.Join(respErr, decodeErr, responseVarsErr)
	return host, resp.StatusCode(), err
}

func (c *Client) CreateHost(hostNew HostJson, options *HostListOptions) (host HostJson, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&HostJson{}).
		SetBody(hostNew).
		Post("/host")

	result := (*resp.Result().(*HostJson))
	return result, resp.StatusCode(), err
}

func (c *Client) CreateHostEncode(hostNew Host, options *HostListOptions) (host Host, statusCode int, Error error) {

	//PrettyPrint("hostNew", hostNew)
	decodedHost, decodeErr := decodeHost(hostNew)
	//PrettyPrint("decodedHost", decodedHost)

	resp, respErr := c.Client.R().
		SetResult(&HostJson{}).
		SetBody(decodedHost).
		Post("/host")

	result := (*resp.Result().(*HostJson))
	//PrettyPrint("result", result)
	hostEncoded, responseVarsErr := encodeHost(result)
	//PrettyPrint("hostEncoded", hostEncoded)
	err := errors.Join(respErr, decodeErr,responseVarsErr)
	return hostEncoded, resp.StatusCode(), err
}

func (c *Client) DeleteHost(hostID string, options *HostListOptions) (host Host, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Host{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"hostID": hostID,
		}).
		Delete("/host/{hostID}")

	result := (*resp.Result().(*Host))
	return result, resp.StatusCode(), err

}
