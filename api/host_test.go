//go:build integration || host

package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHostIntegration(t *testing.T) {
	hostCreateResponse := DoTestCreateHost(t) //C
	t.Logf("Id: %v", hostCreateResponse.ID)
	DoTestGetHost(t, hostCreateResponse.ID)    //R
	DoTestUpdateHost(t, hostCreateResponse.ID) //U
	updatedHost := DoTestGetHost(t, hostCreateResponse.ID)
	assert.Equal(t, "name", updatedHost.Name)
	DoTestDeleteHost(t, hostCreateResponse.ID) //D
}

func DoTestGetHost(t *testing.T, hostID string) (host Host) {
	c := NewClient(os.Getenv("DOG_API_KEY"))

	res, statusCode, err := c.GetHost(hostID, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.HostKey, "expecting non-empty hostkey")
	assert.Equal(t, res.ID, hostID)
	return res
}

func DoTestUpdateHost(t *testing.T, hostID string) (host Host) {
	c := NewClient(os.Getenv("DOG_API_KEY"))

	update := HostUpdateRequest{Group: "group", HostKey: "hostkey", Name: "name"}
	res, statusCode, err := c.UpdateHost(hostID, update, nil)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.HostKey, "expecting non-empty hostkey")
	UpdatedHost, statusCode, err := c.GetHost(hostID, nil)
	assert.Equal(t, 200, statusCode)
	assert.Equal(t, UpdatedHost.HostKey, "hostkey")
	return res
}

func DoTestCreateHost(t *testing.T) (hostCreateResponse HostCreateResponse) {
	c := NewClient(os.Getenv("DOG_API_KEY"))

	newHost := HostCreateRequest{
		Active:      "active",
		Environment: "*",
		Group:       "new_group",
		HostKey:     "new_hostkey",
		Location:    "*",
		Name:        "new_name",
	}

	res, statusCode, err := c.CreateHost(newHost, nil)
	assert.Equal(t, 201, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)
	return res
}
func DoTestDeleteHost(t *testing.T, hostID string) {
	c := NewClient(os.Getenv("DOG_API_KEY"))

	res, statusCode, err := c.DeleteHost(hostID, nil)
	assert.Equal(t, 204, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.Empty(t, res, "expecting empty response")
}
