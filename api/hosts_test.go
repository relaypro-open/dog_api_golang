//go:build integration || host

package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHosts(t *testing.T) {
	c := NewClient(os.Getenv("DOG_API_KEY"))

	res, statusCode, err := c.GetHosts(nil)
	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res[0].HostKey, "expecting non-empty hostkey")
}
