//go:build integration || zone

package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetZones(t *testing.T) {
	c := NewClient(os.Getenv("DOG_API_KEY"))

	res, statusCode, err := c.GetZones(nil)
	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res[0].ID %s\n", res[0].ID)

	assert.NotEmpty(t, res[0].ID, "expecting non-empty Rules")
}
