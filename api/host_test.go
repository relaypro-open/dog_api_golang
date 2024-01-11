//go:build integration || host

package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//func TestFail(t *testing.T) {
//	DoTestCreateHostFail(t) //C
//}

func TestHostIntegration(t *testing.T) {
	hostCreateResponse := DoTestCreateHost(t) //C
	t.Logf("Id: %v", hostCreateResponse.ID)
	DoTestGetHosts(t)                          //R
	DoTestGetHost(t, hostCreateResponse.ID)    //R
	DoTestUpdateHost(t, hostCreateResponse.ID) //U
	updatedHost := DoTestGetHost(t, hostCreateResponse.ID)
	assert.Equal(t, "update_name", updatedHost.Name)
	DoTestDeleteHost(t, hostCreateResponse.ID) //D
}

func DoTestGetHosts(t *testing.T) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetHosts(nil)
	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res[0].HostKey, "expecting non-empty hostkey")
}

func DoTestGetHost(t *testing.T, hostID string) (host Host) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetHost(hostID, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.HostKey, "expecting non-empty hostkey")
	assert.Equal(t, res.ID, hostID)
	t.Logf("res: %+v\n", res)
	return res
}

func DoTestUpdateHost(t *testing.T, hostID string) (host Host) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	updateHost := HostUpdateRequest{
		Environment: "*",
		Group:       "update_group",
		HostKey:     "update-hostkey",
		Location:    "*",
		Name:        "update_name",
		Vars: 	 `{
			"test": "host_test",
			"boolean": true,
			"integer": 1
		}`,
	}
	res, statusCode, err := c.UpdateHost(hostID, updateHost, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)

	return res
}

func DoTestCreateHost(t *testing.T) (host Host) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	newHost := HostCreateRequest{
		Environment: "*",
		Group:       "new_group",
		HostKey:     "new_hostkey",
		Location:    "*",
		Name:        "new_name",
		Vars: 	     `{
			"test": "host_test",
			"boolean": true,
			"integer": 1
		}`,
	}

	res, statusCode, err := c.CreateHost(newHost, nil)

	assert.Equal(t, 201, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)
	return res
}

//func DoTestCreateHostFail(t *testing.T) (hostCreateResponse HostCreateResponse) {
//	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))
//
//	newHost := HostCreateRequest{
//		Environment: "*",
//		Group:       "new_group",
//		HostKey:     "new_hostkey",
//		Location:    "*",
//		Name:        "new_name",
//	}
//
//	res, statusCode, err := c.CreateHost(newHost, nil)
//	assert.Equal(t, 201, statusCode)
//	assert.Nil(t, err, "expecting nil error")
//	assert.NotNil(t, res, "expecting non-nil result")
//	t.Logf("err: %v", err)
//	t.Logf("res: %+v\n", res)
//	return res
//}
func DoTestDeleteHost(t *testing.T, hostID string) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.DeleteHost(hostID, nil)
	assert.Equal(t, 204, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.Empty(t, res, "expecting empty response")
}
