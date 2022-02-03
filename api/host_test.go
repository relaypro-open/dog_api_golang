package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHost(t *testing.T) {
	c := NewClient("my-key")

	HostId := "eda000f6-0743-448f-874b-a7703ecddfaa"
	res, statusCode, err := c.GetHost(HostId, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %v", res)

	assert.NotEmpty(t, res.HostKey, "expecting non-empty hostkey")
	assert.Equal(t, res.Id, HostId)
}

func TestUpdateHost(t *testing.T) {
	c := NewClient("my-key")

	HostId := "eda000f6-0743-448f-874b-a7703ecddfaa"
	update := HostUpdateRequest{Group: "group", HostKey: "hostkey", Name: "name"}
	res, statusCode, err := c.UpdateHost(HostId, update, nil)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %v", res)

	assert.NotEmpty(t, res.HostKey, "expecting non-empty hostkey")
	UpdatedHost, statusCode, err := c.GetHost(HostId, nil)
	assert.Equal(t, 200, statusCode)
	assert.Equal(t, UpdatedHost.HostKey, "hostkey")
}

func TestCreateDeleteHost(t *testing.T) {
	c := NewClient("my-key")

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
	t.Logf("res: %v", res)
	//Test DeleteHost
	HostId := res.Id
	DeleteRes, DeleteStatusCode, DeleteErr := c.DeleteHost(HostId, nil)
	assert.Equal(t, 204, DeleteStatusCode)
	assert.Nil(t, DeleteErr, "expecting nil error")
	assert.NotNil(t, DeleteRes, "expecting non-nil result")
}
