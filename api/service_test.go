//go:build integration || service

package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceIntegration(t *testing.T) {
	ServiceCreateResponse := DoTestCreateService(t) //C
	t.Logf("Id: %v", ServiceCreateResponse.ID)
	DoTestGetService(t, ServiceCreateResponse.ID)    //R
	DoTestUpdateService(t, ServiceCreateResponse.ID) //U
	updatedService := DoTestGetService(t, ServiceCreateResponse.ID)
	assert.Equal(t, "name_update", updatedService.Name)
	DoTestDeleteService(t, ServiceCreateResponse.ID) //D
}

func DoTestGetService(t *testing.T, ServiceID string) (Service Service) {
	c := NewClient(os.Getenv("DOG_API_KEY"))

	res, statusCode, err := c.GetService(ServiceID, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, res.ID, ServiceID)
	return res
}

func DoTestUpdateService(t *testing.T, ServiceID string) (Service Service) {
	c := NewClient(os.Getenv("DOG_API_KEY"))

	update := ServiceUpdateRequest{
		Services: []Services{
			Services{
				Ports:    []string{"1:65535"},
				Protocol: "tcp",
			},
		},
		Name:    "name_update",
		Version: 2,
	}
	res, statusCode, err := c.UpdateService(ServiceID, update, nil)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, "name_update", res.Name)
	assert.Equal(t, 200, statusCode)
	return res
}

func DoTestCreateService(t *testing.T) (serviceCreateReponse ServiceCreateResponse) {
	c := NewClient(os.Getenv("DOG_API_KEY"))

	newService := ServiceCreateRequest{
		Services: []Services{
			Services{
				Ports:    []string{"2:65534"},
				Protocol: "udp",
			},
		},
		Name:    "name",
		Version: 1,
	}
	res, statusCode, err := c.CreateService(newService, nil)
	assert.Equal(t, 201, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)
	return res
}

func DoTestDeleteService(t *testing.T, ServiceID string) {
	c := NewClient(os.Getenv("DOG_API_KEY"))

	res, statusCode, err := c.DeleteService(ServiceID, nil)
	assert.Equal(t, 204, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.Empty(t, res, "expecting empty response")
}