package api

import (
	"strconv"
)

type InventoriesList []Inventory

type InventoriesListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type Inventory struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Groups        []*InventoryGroup `json:"groups"`
}

type InventoryGroup struct {
	Name          string   `json:"name"`
	Vars          map[string]string `json:"vars"`
	Hosts         map[string]map[string]string `json:"hosts"`
}

type InventoryListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

// InventoryUpdateRequest is a struct for the request object required to update an Inventory
type InventoryUpdateRequest struct {
	Name          string   `json:"name"`
	Groups        []*InventoryGroup `json:"groups"`
}

type InventoryCreateRequest struct {
	Name          string   `json:"name"`
	Groups        []*InventoryGroup `json:"groups"`
}

type InventoryCreateResponse struct {
	ID     string `json:"id"`
	Result string `json:"result"`
}

func (c *Client) GetInventories(options *InventoriesListOptions) (inventoriesList InventoriesList, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&InventoriesList{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		Get("/inventories")

	result := (*resp.Result().(*InventoriesList))
	return result, resp.StatusCode(), err
}

func (c *Client) GetInventory(InventoryID string, options *InventoryListOptions) (inventory Inventory, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Inventory{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"InventoryID": InventoryID,
		}).
		Get("/inventory/{InventoryID}")

	result := (*resp.Result().(*Inventory))
	return result, resp.StatusCode(), err

}

func (c *Client) GetInventoryByName(InventoryName string, options *InventoryListOptions) (inventory Inventory, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Inventory{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"InventoryName": InventoryName,
		}).
		Get("/inventory?name={InventoryName}")

	result := (*resp.Result().(*Inventory))
	return result, resp.StatusCode(), err

}

func (c *Client) UpdateInventory(InventoryID string, InventoryUpdate InventoryUpdateRequest, options *InventoryListOptions) (inventory Inventory, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Inventory{}).
		SetPathParams(map[string]string{
			"InventoryID": InventoryID,
		}).
		SetBody(InventoryUpdate).
		Put("/inventory/{InventoryID}")

	result := (*resp.Result().(*Inventory))
	return result, resp.StatusCode(), err
}

func (c *Client) CreateInventory(inventoryNew InventoryCreateRequest, options *InventoryListOptions) (inventory Inventory, statusCode int, Error error) {

	resp, err := c.Client.R().
		SetResult(&Inventory{}).
		SetBody(inventoryNew).
		Post("/inventory")

	result := (*resp.Result().(*Inventory))
	return result, resp.StatusCode(), err
}

func (c *Client) DeleteInventory(InventoryID string, options *InventoryListOptions) (inventory Inventory, statusCode int, Error error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	resp, err := c.Client.R().
		SetResult(&Inventory{}).
		SetQueryParams(map[string]string{
			"page_no": strconv.Itoa(page),
			"limit":   strconv.Itoa(limit),
		}).
		SetPathParams(map[string]string{
			"InventoryID": InventoryID,
		}).
		Delete("/inventory/{InventoryID}")

	result := (*resp.Result().(*Inventory))
	return result, resp.StatusCode(), err
}
