//go:build integration || rule

package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleIntegration(t *testing.T) {
	ruleCreateResponse := DoTestCreateRule(t) //C
	t.Logf("Id: %v", ruleCreateResponse.ID)
	DoTestGetRules(t)                                               //R
	DoTestGetRule(t, ruleCreateResponse.ID)                      //R
	ruleUpdated := DoTestUpdateRule(t, ruleCreateResponse.ID) //U
	updatedRule := DoTestGetRule(t, ruleUpdated.ID)           //Updating Rules create new Rules
	assert.Equal(t, "name_update", updatedRule.Name)
	DoTestDeleteRule(t, ruleCreateResponse.ID) //D
}

func DoTestGetRules(t *testing.T) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetRules(nil)
	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res[0].ID %s\n", res[0].ID)

	assert.NotEmpty(t, res[0].ID, "expecting non-empty Rules")
}

func DoTestGetRule(t *testing.T, RuleID string) (rule Rule) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetRule(RuleID, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, res.ID, RuleID)
	return res
}

func DoTestUpdateRule(t *testing.T, RuleID string) (rule Rule) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	update := RuleUpdateRequest{
		Rules: &FwRules{
			Inbound: []*FwRule{
				&FwRule{
					Action:       "ACCEPT",
					Active:       true,
					Comment:      "",
					Environments: []string{},
					Group:        "cpz_any",
					GroupType:    "ZONE",
					Interface:    "",
					Log:          false,
					LogPrefix:    "",
					Order:        1,
					Service:      "any",
					States:       []string{},
					Type:         "BASIC",
				},
			},
			Outbound: []*FwRule{},
		},
		Name:    "name_update",
	}
	res, statusCode, err := c.UpdateRule(RuleID, update, nil)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, "name_update", res.Name)
	assert.Equal(t, 200, statusCode)
	return res
}

func DoTestCreateRule(t *testing.T) (rule Rule) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	newRule := RuleCreateRequest{
		Rules: &FwRules{
			Inbound: []*FwRule{
				&FwRule{
					Action:       "ACCEPT",
					Active:       true,
					Comment:      "",
					Environments: []string{},
					Group:        "cpz_any",
					GroupType:    "ZONE",
					Interface:    "",
					Log:          false,
					LogPrefix:    "",
					Order:        1,
					Service:      "any",
					States:       []string{},
					Type:         "BASIC",
				},
			},
			Outbound: []*FwRule{},
		},
		Name:    "name",
	}

	res, statusCode, err := c.CreateRule(newRule, nil)
	assert.Equal(t, 201, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)
	return res
}

func DoTestDeleteRule(t *testing.T, RuleID string) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.DeleteRule(RuleID, nil)
	assert.Equal(t, 204, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.Empty(t, res, "expecting empty response")
}
