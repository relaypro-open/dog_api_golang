package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHost(t *testing.T) {
	c := NewClient("my-key")

	HostId := "eda000f6-0743-448f-874b-a7703ecddfaa"
	res, err := c.GetHost(HostId, nil)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %v", res)

	assert.NotEmpty(t, res.HostKey, "expecting non-empty hostkey")
	assert.Equal(t, res.Id, HostId)
}

func TestUpdateHost(t *testing.T) {
	c := NewClient("my-key")

	HostId := "eda000f6-0743-448f-874b-a7703ecddfaa"
	update := HostUpdateRequest{"group", "hostkey", "name"}
	res, err := c.UpdateHost(HostId, update, nil)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %v", res)

	assert.NotEmpty(t, res.HostKey, "expecting non-empty hostkey")
	UpdatedHost, err := c.GetHost(HostId, nil)
	assert.Equal(t, UpdatedHost.HostKey, "hostkey")
}
