//go:build integration || external

package api

//import (
//	"os"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//)

//func TestExternalIntegration(t *testing.T) {
//	DoTestGetExternals(t)                                     //R
//}
//
//func DoTestGetExternals(t *testing.T) {
//	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))
//
//	res, statusCode, err := c.GetExternals(nil)
//	assert.Equal(t, 200, statusCode)
//	assert.Nil(t, err, "expecting nil error")
//	assert.NotNil(t, res, "expecting non-nil result")
//	t.Logf("res[0]: %+v\n", res[0])
//
//	assert.NotEmpty(t, res[0].ID, "expecting non-empty ID")
//}
//
//func DoTestGetExternal(t *testing.T, ExternalID string) (External External) {
//	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))
//
//	res, statusCode, err := c.GetExternal(ExternalID, nil)
//
//	assert.Equal(t, 200, statusCode)
//	assert.Nil(t, err, "expecting nil error")
//	assert.NotNil(t, res, "expecting non-nil result")
//	t.Logf("res: %+v\n", res)
//
//	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
//	assert.Equal(t, res.ID, ExternalID)
//	return res
//}
