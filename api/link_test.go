//go:build integration || link

package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkIntegration(t *testing.T) {
	linkCreateResponse := DoTestCreateLink(t) //C
	t.Logf("Id: %v", linkCreateResponse.ID)
	DoTestGetExternals(t)                                     //R
	DoTestGetLinks(t)                                         //R
	DoTestGetLink(t, linkCreateResponse.ID)                   //R
	linkUpdated := DoTestUpdateLink(t, linkCreateResponse.ID) //U
	updatedLink := DoTestGetLink(t, linkUpdated.ID)           //Updating Links create new Links
	assert.Equal(t, "z2", updatedLink.Name)
	DoTestDeleteLink(t, linkCreateResponse.ID) //D
}

func DoTestGetLinks(t *testing.T) {
	c := NewClient(os.Getenv("DOG_API_KEY"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetLinks(nil)
	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res[0].ID %s\n", res[0].ID)

	assert.NotEmpty(t, res[0].ID, "expecting non-empty Rules")
}

func DoTestGetLink(t *testing.T, LinkID string) (Link Link) {
	c := NewClient(os.Getenv("DOG_API_KEY"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetLink(LinkID, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, res.ID, LinkID)
	return res
}

func DoTestUpdateLink(t *testing.T, LinkID string) (Link Link) {
	c := NewClient(os.Getenv("DOG_API_KEY"), os.Getenv("DOG_API_ENDPOINT"))

	update := LinkUpdateRequest{
		AddressHandling: "union",
		Connection: Connection{
			ApiPort:  1234,
			Host:     "host",
			Password: "password",
			Port:     2345,
			SSLOptions: SSLOptions{
				CaCertFile:           "cacertfile",
				CertFile:             "certfile",
				FailIfNoPeerCert:     true,
				KeyFile:              "keyfile",
				ServerNameIndication: "disable",
				Verify:               "verify_peer",
			},
			User:        "user",
			VirtualHost: "virtual_host",
		},
		ConnectionType: "thumper",
		Direction:      "bidirectional",
		Enabled:        false, //Do not test with Enabled: true
		Name:           "z2",
	}
	res, statusCode, err := c.UpdateLink(LinkID, update, nil)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, "z2", res.Name)
	assert.Equal(t, 200, statusCode)
	return res
}

func DoTestCreateLink(t *testing.T) (link Link) {
	c := NewClient(os.Getenv("DOG_API_KEY"), os.Getenv("DOG_API_ENDPOINT"))

	newLink := LinkCreateRequest{
		AddressHandling: "union",
		Connection: Connection{
			ApiPort:  3456,
			Host:     "host",
			Password: "password",
			Port:     7890,
			SSLOptions: SSLOptions{
				CaCertFile:           "cacertfile",
				CertFile:             "certfile",
				FailIfNoPeerCert:     true,
				KeyFile:              "keyfile",
				ServerNameIndication: "disable",
				Verify:               "verify_peer",
			},
			User:        "user",
			VirtualHost: "virtual_host",
		},
		ConnectionType: "thumper",
		Direction:      "bidirectional",
		Enabled:        false, //Do not test with Enabled: true
		Name:           "z1",
	}

	res, statusCode, err := c.CreateLink(newLink, nil)
	assert.Equal(t, 201, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)
	return res
}

func DoTestDeleteLink(t *testing.T, LinkID string) {
	c := NewClient(os.Getenv("DOG_API_KEY"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.DeleteLink(LinkID, nil)
	assert.Equal(t, 204, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.Empty(t, res, "expecting empty response")
}

func DoTestGetExternals(t *testing.T) {
	c := NewClient(os.Getenv("DOG_API_KEY"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetExternals(nil)
	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res[0]: %+v\n", res[0])

	assert.NotEmpty(t, res[0].ID, "expecting non-empty ID")
}

func DoTestGetExternal(t *testing.T, ExternalID string) (External External) {
	c := NewClient(os.Getenv("DOG_API_KEY"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetExternal(ExternalID, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, res.ID, ExternalID)
	return res
}
