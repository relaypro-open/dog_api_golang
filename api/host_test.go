
//go:build integration || host

package api

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHostIntegration(t *testing.T) {
	hostEncodeCreateNoVarsResponse := DoTestCreateHostEncodeNoVars(t) //C
	t.Logf("Id: %v", hostEncodeCreateNoVarsResponse.ID)
	DoTestDeleteHost(t, hostEncodeCreateNoVarsResponse.ID) //D
	hostEncodeCreateResponse := DoTestCreateHostEncode(t) //C
	t.Logf("Id: %v", hostEncodeCreateResponse.ID)
	DoTestGetHostsEncode(t)                          //R
	DoTestGetHostEncode(t, hostEncodeCreateResponse.ID)    //R
	DoTestUpdateHostEncode(t, hostEncodeCreateResponse.ID) //U
	updatedHostEncode := DoTestGetHostEncode(t, hostEncodeCreateResponse.ID)
	assert.Equal(t, "update_name", updatedHostEncode.Name)
	DoTestDeleteHost(t, hostEncodeCreateResponse.ID) //D
	
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

func DoTestGetHost(t *testing.T, hostID string) (host HostJson) {
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

func DoTestUpdateHost(t *testing.T, hostID string) (host HostJson) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	updateHost := HostJson{
		Environment: "*",
		Group:       "update_group",
		HostKey:     "update-hostkey",
		Location:    "*",
		Name:        "update_name",
		Vars: 	     map[string]any{
			"test": "host_test",
			"boolean":  true,
			"integer": 1,
		},
	}
	res, statusCode, err := c.UpdateHost(hostID, updateHost, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)

	return res
}

func DoTestUpdateHostNoLocation(t *testing.T, hostID string) (host HostJson) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	updateHost := HostJson{
		Environment: "*",
		Group:       "update_group",
		HostKey:     "update-hostkey",
		Name:        "update_name",
		Vars: 	     map[string]any{
			"test": "host_test",
			"boolean":  true,
			"integer": 1,
		},
	}
	res, statusCode, err := c.UpdateHost(hostID, updateHost, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)

	return res
}
func DoTestUpdateHostNoVars(t *testing.T, hostID string) (host HostJson) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	updateHost := HostJson{
		Environment: "*",
		Group:       "update_group",
		HostKey:     "update-hostkey",
		Name:        "update_name",
		Location:    "*",
	}
	res, statusCode, err := c.UpdateHost(hostID, updateHost, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)

	return res
}

func DoTestCreateHost(t *testing.T) (host HostJson) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	newHost := HostJson{
		Environment: "*",
		Group:       "new_group",
		HostKey:     "new_hostkey",
		Location:    "*",
		Name:        "new_name",
		Vars: 	     map[string]any{
			"test": "host_test",
			"boolean":  true,
			"integer": 1,
		},
	}

	res, statusCode, err := c.CreateHost(newHost, nil)

	assert.Equal(t, 201, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)
	return res
}

func DoTestGetHostsEncode(t *testing.T) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetHostsEncode(nil)
	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res[0].HostKey, "expecting non-empty hostkey")
}

func DoTestGetHostsEncodeActive(t *testing.T) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))
	hla := HostsListOptions{
		Active: "true",
	}
	res, statusCode, err := c.GetHostsEncode(&hla)
	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res[0].HostKey, "expecting non-empty hostkey")
}

func DoTestGetHostEncode(t *testing.T, hostID string) (host Host) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetHostEncode(hostID, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.HostKey, "expecting non-empty hostkey")
	assert.Equal(t, res.ID, hostID)
	t.Logf("res: %+v\n", res)
	return res
}

func DoTestUpdateHostEncode(t *testing.T, hostID string) (host Host) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	updateHost := Host{
		Environment: "*",
		Group:       "update_group",
		HostKey:     "update-hostkey",
		Location:    "*",
		Name:        "update_name",
		Vars: 	     `{
			"test": "host_test",
			"boolean":  true,
			"integer": 1
		}`,
	}
	res, statusCode, err := c.UpdateHostEncode(hostID, updateHost, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)

	return res
}

func DoTestUpdateHostEncodeEmptyVars(t *testing.T, hostID string) (host Host) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	updateHost := Host{
		Environment: "*",
		Group:       "update_group",
		HostKey:     "update-hostkey",
		Location:    "*",
		Name:        "update_name",
		Vars: 	     "null",
	}
	res, statusCode, err := c.UpdateHostEncode(hostID, updateHost, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	assert.Equal(t, "{}", res.Vars)
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)

	return res
}

func DoTestCreateHostEncode(t *testing.T) (host Host) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	newHost := Host{
		Environment: "*",
		Group:       "new_group",
		HostKey:     "new_hostkey",
		Location:    "*",
		Name:        "new_name",
		Vars: 	     `{
			"test": "host_test",
			"boolean":  true,
			"integer": 1
		}`,
	}

	res, statusCode, err := c.CreateHostEncode(newHost, nil)

	assert.Equal(t, 201, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)
	return res
}

func DoTestCreateHostEncodeNoVars(t *testing.T) (host Host) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	newHost := Host{
		Environment: "*",
		Group:       "new_group",
		HostKey:     "new_hostkey",
		Location:    "*",
		Name:        "new_name",
	}

	res, statusCode, err := c.CreateHostEncode(newHost, nil)

	assert.Equal(t, 201, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)
	return res
}

func DoTestDeleteHost(t *testing.T, hostID string) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.DeleteHost(hostID, nil)
	assert.Equal(t, 204, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.Empty(t, res, "expecting empty response")
}
