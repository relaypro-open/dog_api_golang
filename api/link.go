package api

import (
	"strconv"
)

type Link struct {
	ID              string      `json:"id"`
	AddressHandling string      `json:"address_handling"`
	Connection      *Connection `json:"connection"`
	ConnectionType  string      `json:"connection_type"`
	Direction       string      `json:"direction"`
	Enabled         bool        `json:"enabled"`
	Name            string      `json:"name"`
}

type Connection struct {
	ApiPort     int         `json:"api_port"`
	Host        string      `json:"host"`
	Password    string      `json:"password"`
	Port        int         `json:"port"`
	SSLOptions  *SSLOptions `json:"ssl_options"`
	User        string      `json:"user"`
	VirtualHost string      `json:"virtual_host"`
}

type SSLOptions struct {
	CaCertFile           string `json:"cacertfile"`
	CertFile             string `json:"certfile"`
	FailIfNoPeerCert     bool   `json:"fail_if_no_peer_cert"`
	KeyFile              string `json:"keyfile"`
	ServerNameIndication string `json:"server_name_indication"`
	Verify               string `json:"verify"`
}

type LinkListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

// LinkUpdateRequest is a struct for the request object required to update a Link
type LinkUpdateRequest struct {
	AddressHandling string      `json:"address_handling"`
	Connection      *Connection `json:"connection"`
	ConnectionType  string      `json:"connection_type"`
	Direction       string      `json:"direction"`
	Enabled         bool        `json:"enabled"`
	Name            string      `json:"name"`
}

type LinkCreateRequest struct {
	AddressHandling string      `json:"address_handling"`
	Connection      *Connection `json:"connection"`
	ConnectionType  string      `json:"connection_type"`
	Direction       string      `json:"direction"`
	Enabled         bool        `json:"enabled"`
	Name            string      `json:"name"`
}

type LinkCreateResponse struct {
	ID     string `json:"id"`
	Result string `json:"result"`
}

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

func (c *Client) GetLink(LinkID string, options *LinkListOptions) (link Link, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Link{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"LinkID": LinkID,
		}).
		Get("/link/{LinkID}")

	result := (*resp.Result().(*Link))
	return result, resp.StatusCode(), err

}

func (c *Client) UpdateLink(LinkID string, LinkUpdate LinkUpdateRequest, options *LinkListOptions) (link Link, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Link{}).
		SetPathParams(map[string]string{
			"LinkID": LinkID,
		}).
		SetBody(LinkUpdate).
		Put("/link/{LinkID}")

	result := (*resp.Result().(*Link))
	return result, resp.StatusCode(), err
}

func (c *Client) CreateLink(linkNew LinkCreateRequest, options *LinkListOptions) (link Link, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Link{}).
		SetBody(linkNew).
		Post("/link")

	result := (*resp.Result().(*Link))
	return result, resp.StatusCode(), err
}

func (c *Client) DeleteLink(LinkID string, options *LinkListOptions) (link Link, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Link{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"LinkID": LinkID,
		}).
		Delete("/link/{LinkID}")

	result := (*resp.Result().(*Link))
	return result, resp.StatusCode(), err
}
