//go:build integration || group

package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupIntegration(t *testing.T) {
	GroupEncodeCreateResponse := DoTestCreateGroupEncode(t) //C
	t.Logf("Id: %v", GroupEncodeCreateResponse.ID)
	DoTestGetGroupsEncode(t)                           //R
	DoTestGetGroupEncode(t, GroupEncodeCreateResponse.ID)    //R
	DoTestUpdateGroupEncode(t, GroupEncodeCreateResponse.ID) //U
	updatedGroupEncode := DoTestGetGroupEncode(t, GroupEncodeCreateResponse.ID)
	assert.Equal(t, "name_update", updatedGroupEncode.Name)
	DoTestDeleteGroup(t, GroupEncodeCreateResponse.ID) //D

	GroupCreateResponse := DoTestCreateGroup(t) //C
	t.Logf("Id: %v", GroupCreateResponse.ID)
	DoTestGetGroups(t)                           //R
	DoTestGetGroup(t, GroupCreateResponse.ID)    //R
	DoTestUpdateGroup(t, GroupCreateResponse.ID) //U
	updatedGroup := DoTestGetGroup(t, GroupCreateResponse.ID)
	assert.Equal(t, "name_update", updatedGroup.Name)
	DoTestDeleteGroup(t, GroupCreateResponse.ID) //D
}

func DoTestGetGroups(t *testing.T) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetGroups(nil)
	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	//t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res[0].ID, "expecting non-empty Rules")
}

func DoTestGetGroup(t *testing.T, GroupID string) (group GroupJson) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetGroup(GroupID, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, res.ID, GroupID)
	return res
}

func DoTestUpdateGroup(t *testing.T, GroupID string) (group GroupJson) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	update := GroupJson{
		Description:    "description_update",
		Name:           "name_update",
		ProfileId:      "profile_id_update",
		ProfileName:    "profile_name_update",
		ProfileVersion: "profile_version_update",
		Ec2SecurityGroupIds:  []*Ec2SecurityGroupIds{
			&Ec2SecurityGroupIds{
				Region: "us-test-region",
				SgId: "sg-test",
			},
		},
		Vars: map[string]any{
			"test": "group_test",
			"boolean": true,
			"integer": 1,
		},
	}
	res, statusCode, err := c.UpdateGroup(GroupID, update, nil)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, "name_update", res.Name)
	assert.Equal(t, 200, statusCode)
	return res
}

func DoTestCreateGroup(t *testing.T) (group GroupJson) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	newGroup := GroupJson{
		Description:    "description",
		Name:           "name",
		ProfileId:      "profile_id",
		ProfileName:    "profile_name",
		ProfileVersion: "profile_version",
		Ec2SecurityGroupIds:  []*Ec2SecurityGroupIds{
			&Ec2SecurityGroupIds{
				Region: "us-test-region",
				SgId: "sg-test",
			},
		},
		Vars: map[string]any{
			"test": "group_test",
			"boolean": true,
			"integer": 1,
		},
	}
	res, statusCode, err := c.CreateGroup(newGroup, nil)
	assert.Equal(t, 201, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)
	return res
}

func DoTestGetGroupsEncode(t *testing.T) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetGroupsEncode(nil)
	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	//t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res[0].ID, "expecting non-empty Rules")
}

func DoTestGetGroupEncode(t *testing.T, GroupID string) (Group Group) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetGroupEncode(GroupID, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, res.ID, GroupID)
	return res
}

func DoTestUpdateGroupEncode(t *testing.T, GroupID string) (group Group) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	update := Group{
		Description:    "description_update",
		Name:           "name_update",
		ProfileId:      "profile_id_update",
		ProfileName:    "profile_name_update",
		ProfileVersion: "profile_version_update",
		Ec2SecurityGroupIds:  []*Ec2SecurityGroupIds{
			&Ec2SecurityGroupIds{
				Region: "us-test-region",
				SgId: "sg-test",
			},
		},
		Vars: `{
			"test": "group_test",
			"boolean": true,
			"integer": 1
		}`,
	}
	res, statusCode, err := c.UpdateGroupEncode(GroupID, update, nil)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, "name_update", res.Name)
	assert.Equal(t, 200, statusCode)
	return res
}

func DoTestCreateGroupEncode(t *testing.T) (group Group) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	newGroup := Group{
		Description:    "description",
		Name:           "name",
		ProfileId:      "profile_id",
		ProfileName:    "profile_name",
		ProfileVersion: "profile_version",
		Ec2SecurityGroupIds:  []*Ec2SecurityGroupIds{
			&Ec2SecurityGroupIds{
				Region: "us-test-region",
				SgId: "sg-test",
			},
		},
		Vars: `{
			"test": "group_test",
			"boolean": true,
			"integer": 1
		}`,
	}
	res, statusCode, err := c.CreateGroupEncode(newGroup, nil)
	assert.Equal(t, 201, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)
	return res
}

func DoTestDeleteGroup(t *testing.T, GroupID string) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.DeleteGroup(GroupID, nil)
	assert.Equal(t, 204, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.Empty(t, res, "expecting empty response")
}
