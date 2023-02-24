//go:build integration || ruleset

package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleIntegration(t *testing.T) {
	rulesetCreateResponse := DoTestCreateRuleset(t)
	t.Logf("Id: %v", rulesetCreateResponse.ID)
	DoTestGetRulesets(t)                    
	DoTestGetRuleset(t, rulesetCreateResponse.ID)
	DoTestGetRulesetByName(t, rulesetCreateResponse.Name) //R
	DoTestUpdateRuleset(t, rulesetCreateResponse.ID)      //U
	updatedRuleset := DoTestGetRuleset(t, rulesetCreateResponse.ID)
	assert.Equal(t, "name_update", updatedRuleset.Name)
	DoTestDeleteRuleset(t, rulesetCreateResponse.ID)
}

func DoTestGetRulesets(t *testing.T) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetRulesets(nil)
	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res[0].ID %s\n", res[0].ID)

	assert.NotEmpty(t, res[0].ID, "expecting non-empty Rules")
}

func DoTestGetRuleset(t *testing.T, RulesetID string) (rule Ruleset) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetRuleset(RulesetID, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, res.ID, RulesetID)
	return res
}

func DoTestGetRulesetByName(t *testing.T, RulesetName string) (Ruleset Ruleset) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetRulesetByName(RulesetName, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, res.Name, RulesetName)
	return res
}

func DoTestUpdateRuleset(t *testing.T, RulesetID string) (ruleset Ruleset) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

    rules := &Rules{
			Inbound: []*Rule{
				&Rule{
					Action:       "ACCEPT",
					Active:       true,
					Comment:      "",
					Environments: []string{},
					Group:        "any",
					GroupType:    "ANY",
					Interface:    "",
					Log:          false,
					LogPrefix:    "",
					Order:        1,
					Service:      "any",
					States:       []string{},
					Type:         "BASIC",
				},
			},
			Outbound: []*Rule{
				&Rule{
					Action:       "DROP",
					Active:       true,
					Comment:      "",
					Environments: []string{},
					Group:        "any",
					GroupType:    "ANY",
					Interface:    "",
					Log:          false,
					LogPrefix:    "",
					Order:        1,
					Service:      "any",
					States:       []string{},
					Type:         "BASIC",
				},
			},
		}

	update := RulesetUpdateRequest{
		Rules:   rules,
		Name:    "name_update",
	}

	t.Logf("RulesetID: %+v\n", RulesetID)
	t.Logf("rules: %+v\n", rules)
	t.Logf("update: %+v\n", update)
	res, statusCode, err := c.UpdateRuleset(RulesetID, update, nil)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, "name_update", res.Name)
	assert.Equal(t, 200, statusCode)
	return res
}

func DoTestCreateRuleset(t *testing.T) (rule Ruleset) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	newRule := RulesetCreateRequest{
		Rules: &Rules{
			Inbound: []*Rule{
				&Rule{
					Action:       "ACCEPT",
					Active:       true,
					Comment:      "",
					Environments: []string{},
					Group:        "any",
					GroupType:    "ANY",
					Interface:    "",
					Log:          false,
					LogPrefix:    "",
					Order:        1,
					Service:      "any",
					States:       []string{},
					Type:         "BASIC",
				},
			},
			Outbound: []*Rule{
				&Rule{
					Action:       "DROP",
					Active:       true,
					Comment:      "",
					Environments: []string{},
					Group:        "any",
					GroupType:    "ANY",
					Interface:    "",
					Log:          false,
					LogPrefix:    "",
					Order:        1,
					Service:      "any",
					States:       []string{},
					Type:         "BASIC",
				},
			},
		},
		Name:    "name",
	}

	res, statusCode, err := c.CreateRuleset(newRule, nil)
	assert.Equal(t, 201, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)
	return res
}

func DoTestDeleteRuleset(t *testing.T, RuleID string) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.DeleteRuleset(RuleID, nil)
	assert.Equal(t, 204, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.Empty(t, res, "expecting empty response")
}
