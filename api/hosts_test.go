//go:build integration
// +build integration

package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHosts(t *testing.T) {
	c := NewClient("my-key")

	res, statusCode, err := c.GetHosts(nil)
	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %v", res)

	assert.NotEmpty(t, res[0].HostKey, "expecting non-empty hostkey")
}
