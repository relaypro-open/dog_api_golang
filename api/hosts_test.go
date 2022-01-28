//go:build integration
// +build integration

package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHosts(t *testing.T) {
	//c := NewClient(os.Getenv("HostST_INTEGRATION_API_KEY"))
	c := NewClient("DUMMY_API_KEY")

	ctx := context.Background()
	res, err := c.GetHosts(ctx, nil)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %v", res)
	t.Logf("res.Hosts: %v", res.Hosts)

	assert.NotEmpty(t, res.Hosts[0].HostKey, "expecting non-empty hostkey")
}
