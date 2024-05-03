//go:build integration || fact

package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactIntegration(t *testing.T) {
	FactCreateResponse := DoTestCreateFact(t) //C
	t.Logf("Id: %v", FactCreateResponse.ID)
	DoTestGetFacts(t)                               //R
	DoTestGetFact(t, FactCreateResponse.ID)         //R
	DoTestGetFactByName(t, FactCreateResponse.Name) //R
	DoTestUpdateFact(t, FactCreateResponse.ID)      //U
	updatedFact := DoTestGetFact(t, FactCreateResponse.ID)
	assert.Equal(t, "name_update", updatedFact.Name)
	DoTestDeleteFact(t, FactCreateResponse.ID) //D

	FactCreateResponseEncode := DoTestCreateFactEncode(t) //C
	t.Logf("Id: %v", FactCreateResponseEncode.ID)
	DoTestGetFactsEncode(t)                                     //R
	DoTestGetFactEncode(t, FactCreateResponseEncode.ID)         //R
	DoTestGetFactByNameEncode(t, FactCreateResponseEncode.Name) //R
	DoTestUpdateFactEncode(t, FactCreateResponseEncode.ID)      //U
	updatedFactEncode := DoTestGetFactEncode(t, FactCreateResponseEncode.ID)
	assert.Equal(t, "name_update", updatedFactEncode.Name)
	DoTestDeleteFact(t, FactCreateResponseEncode.ID) //D
}

// *Encode functions are for use by Terraform

func DoTestGetFactsEncode(t *testing.T) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetFactsEncode(nil)
	assert.Equal(t, 200, statusCode)
	t.Logf("err: %+v\n", err)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res[0].ID %s\n", res[0].ID)

	assert.NotEmpty(t, res[0].ID, "expecting non-empty Rules")
}

func DoTestGetFactEncode(t *testing.T, FactID string) (Fact Fact) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetFactEncode(FactID, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, res.ID, FactID)
	return res
}

func DoTestGetFactByNameEncode(t *testing.T, FactName string) (Fact Fact) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetFactByNameEncode(FactName, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, res.Name, FactName)
	return res
}

func DoTestUpdateFactEncode(t *testing.T, FactID string) (fact Fact) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	Vars1 := `{
	    "environment": "mob_dev",
	    "dog_env": "dev",
	    "boolean": true,
	    "integer": 1
        }`

	Hosts1 := `{
		"web.test.abc": {"os": "Linux"},
		"db.test.abc":  {"db": "sql"}
	}`

	Children1 := []string{"test"}

	Ig1 := &FactGroup{
		Vars:     &Vars1,
		Hosts:    &Hosts1,
		Children: Children1}

	update := Fact{
		Name:   "name_update",
		Groups: map[string]*FactGroup{"mob_dev": Ig1},
	}

	res, statusCode, err := c.UpdateFactEncode(FactID, update, nil)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, "name_update", res.Name)
	assert.Equal(t, 200, statusCode)
	return res
}

func DoTestCreateFactEncode(t *testing.T) (fact Fact) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	//Vars1 := `{
	//    "environment": "mob_dev",
	//    "dog_env": "dev",
	//    "boolean": true,
	//    "integer": 1
	//    }`

	Hosts1 := `{
		"web.test.abc": {"os": "Linux"},
		"db.test.abc":  {"db": "sql"}
	}`

	Children1 := []string{"test"}

	Ig1 := &FactGroup{
		Hosts:    &Hosts1,
		Children: Children1}

	newFact := Fact{
		Name:   "name",
		Groups: map[string]*FactGroup{"mob_dev": Ig1},
	}

	res, statusCode, err := c.CreateFactEncode(newFact, nil)
	assert.Equal(t, 201, statusCode)
	assert.NotEmpty(t, res.ID, "expected non-empty ID")
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)
	return res
}

//Non-Terraform functions

func DoTestGetFacts(t *testing.T) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetFacts(nil)
	assert.Equal(t, 200, statusCode)
	t.Logf("err: %+v\n", err)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res[0].ID %s\n", res[0].ID)

	assert.NotEmpty(t, res[0].ID, "expecting non-empty Rules")
}

func DoTestGetFact(t *testing.T, FactID string) (Fact FactJson) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetFact(FactID, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, res.ID, FactID)
	return res
}

func DoTestGetFactByName(t *testing.T, FactName string) (Fact FactJson) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetFactByName(FactName, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, res.Name, FactName)
	return res
}

func DoTestUpdateFact(t *testing.T, FactID string) (Fact FactJson) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	Vars1 := map[string]any{
		"environment": "mob_dev",
		"dog_env":     "dev",
		"boolean":     true,
		"integer":     1,
	}

	Hosts1 := map[string]map[string]any{
		"web.test.abc": map[string]any{"os": "Linux"},
		"db.test.abc":  map[string]any{"db": "sql"},
	}

	Children1 := []string{"test"}

	Ig1 := &FactGroupJson{Vars1, Hosts1, Children1}

	update := FactJson{
		Name:   "name_update",
		Groups: map[string]*FactGroupJson{"mob_dev": Ig1},
	}

	res, statusCode, err := c.UpdateFact(FactID, update, nil)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, "name_update", res.Name)
	assert.Equal(t, 200, statusCode)
	return res
}

func DoTestCreateFact(t *testing.T) (fact FactJson) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	Vars1 := map[string]any{
		"environment": "mob_dev",
		"dog_env":     "dev",
		"boolean":     true,
		"integer":     1,
	}

	Hosts1 := map[string]map[string]any{
		"web.test.abc": map[string]any{"os": "Linux"},
		"db.test.abc":  map[string]any{"db": "sql"},
	}

	Children1 := []string{"test"}

	Ig1 := &FactGroupJson{Vars1, Hosts1, Children1}

	newFact := FactJson{
		Name:   "name",
		Groups: map[string]*FactGroupJson{"mob_dev": Ig1},
	}

	res, statusCode, err := c.CreateFact(newFact, nil)
	assert.Equal(t, 201, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)
	return res
}

func DoTestDeleteFact(t *testing.T, FactID string) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.DeleteFact(FactID, nil)
	assert.Equal(t, 204, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.Empty(t, res, "expecting empty response")
}
