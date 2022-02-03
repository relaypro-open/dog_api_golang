//go:build integration
// +build integration

package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHosts(t *testing.T) {
	c := NewClient("my-key")

	res, err := c.GetHosts(nil)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %v", res)
	t.Logf("res.Hosts: %v", res.Hosts)

	assert.NotEmpty(t, res.Hosts[0].HostKey, "expecting non-empty hostkey")
}
