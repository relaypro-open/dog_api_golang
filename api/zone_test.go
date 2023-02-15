//go:build integration || zone

package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZoneIntegration(t *testing.T) {
	ZoneCreateResponse := DoTestCreateZone(t) //C
	t.Logf("Id: %v", ZoneCreateResponse.ID)
	DoTestGetZones(t)                               //R
	DoTestGetZone(t, ZoneCreateResponse.ID)         //R
	DoTestGetZoneByName(t, ZoneCreateResponse.Name) //R
	DoTestUpdateZone(t, ZoneCreateResponse.ID)      //U
	updatedZone := DoTestGetZone(t, ZoneCreateResponse.ID)
	assert.Equal(t, "name_update", updatedZone.Name)
	DoTestDeleteZone(t, ZoneCreateResponse.ID) //D
}

func DoTestGetZones(t *testing.T) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetZones(nil)
	assert.Equal(t, 200, statusCode)
	t.Logf("err: %+v\n", err)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res[0].ID %s\n", res[0].ID)

	assert.NotEmpty(t, res[0].ID, "expecting non-empty Rules")
}

func DoTestGetZone(t *testing.T, ZoneID string) (Zone Zone) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetZone(ZoneID, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, res.ID, ZoneID)
	return res
}

func DoTestGetZoneByName(t *testing.T, ZoneName string) (Zone Zone) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetZoneByName(ZoneName, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, res.Name, ZoneName)
	return res
}

func DoTestUpdateZone(t *testing.T, ZoneID string) (Zone Zone) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	update := ZoneUpdateRequest{
		IPv4Addresses: []string{"1.2.3.4"},
		IPv6Addresses: []string{"fe80::2dfc:c1fc:eded:fb97/64"},
		Name:          "name_update",
	}
	res, statusCode, err := c.UpdateZone(ZoneID, update, nil)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, "name_update", res.Name)
	assert.Equal(t, 200, statusCode)
	return res
}

func DoTestCreateZone(t *testing.T) (zone Zone) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	newZone := ZoneCreateRequest{
		IPv4Addresses: []string{"1.2.3.4"},
		IPv6Addresses: []string{"fe80::2dfc:c1fc:eded:fb97/64"},
		Name:          "name",
	}

	res, statusCode, err := c.CreateZone(newZone, nil)
	assert.Equal(t, 201, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)
	return res
}

func DoTestDeleteZone(t *testing.T, ZoneID string) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.DeleteZone(ZoneID, nil)
	assert.Equal(t, 204, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.Empty(t, res, "expecting empty response")
}
