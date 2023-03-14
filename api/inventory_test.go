//go:build integration || inventory

package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInventoryIntegration(t *testing.T) {
	InventoryCreateResponse := DoTestCreateInventory(t) //C
	t.Logf("Id: %v", InventoryCreateResponse.ID)
	DoTestGetInventories(t)                               //R
	DoTestGetInventory(t, InventoryCreateResponse.ID)         //R
	DoTestGetInventoryByName(t, InventoryCreateResponse.Name) //R
	DoTestUpdateInventory(t, InventoryCreateResponse.ID)      //U
	updatedInventory := DoTestGetInventory(t, InventoryCreateResponse.ID)
	assert.Equal(t, "name_update", updatedInventory.Name)
	DoTestDeleteInventory(t, InventoryCreateResponse.ID) //D
}

func DoTestGetInventories(t *testing.T) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetInventories(nil)
	assert.Equal(t, 200, statusCode)
	t.Logf("err: %+v\n", err)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res[0].ID %s\n", res[0].ID)

	assert.NotEmpty(t, res[0].ID, "expecting non-empty Rules")
}

func DoTestGetInventory(t *testing.T, InventoryID string) (Inventory Inventory) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetInventory(InventoryID, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, res.ID, InventoryID)
	return res
}

func DoTestGetInventoryByName(t *testing.T, InventoryName string) (Inventory Inventory) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetInventoryByName(InventoryName, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, res.Name, InventoryName)
	return res
}

func DoTestUpdateInventory(t *testing.T, InventoryID string) (Inventory Inventory) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	Vars1 := map[string]string{
	    "environment": "mob_dev",
	    "dog_env": "dev",
        }

	Hosts1 := map[string]map[string]string{
	    "web.test.abc": 
	    	map[string]string{"os": "Linux"},
	    "db.test.abc":
	    	map[string]string{"db": "sql"},
	}


	Ig1 := &InventoryGroup{"name", Vars1, Hosts1 }

	update := InventoryUpdateRequest{
		Name:          "name_update",
		Groups:        []*InventoryGroup{ Ig1 },
	}

	res, statusCode, err := c.UpdateInventory(InventoryID, update, nil)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, "name_update", res.Name)
	assert.Equal(t, 200, statusCode)
	return res
}

func DoTestCreateInventory(t *testing.T) (inventory Inventory) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))
	
	Vars1 := map[string]string{
	    "environment": "mob_dev",
	    "dog_env": "dev",
        }

	Hosts1 := map[string]map[string]string{
	    "web.test.abc": 
	    	map[string]string{"os": "Linux"},
	    "db.test.abc":
	    	map[string]string{"db": "sql"},
	}

	Ig1 := &InventoryGroup{"name", Vars1, Hosts1 }

	newInventory := InventoryCreateRequest{
		Name:          "name",
		Groups:        []*InventoryGroup{ Ig1 },
	}

	res, statusCode, err := c.CreateInventory(newInventory, nil)
	assert.Equal(t, 201, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)
	return res
}

func DoTestDeleteInventory(t *testing.T, InventoryID string) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.DeleteInventory(InventoryID, nil)
	assert.Equal(t, 204, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.Empty(t, res, "expecting empty response")
}
