//go:build integration || profile

package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfileIntegration(t *testing.T) {
	profileCreateResponse := DoTestCreateProfile(t) //C
	t.Logf("Id: %v", profileCreateResponse.ID)
	DoTestGetProfiles(t)                                               //R
	DoTestGetProfile(t, profileCreateResponse.ID)                      //R
	profileUpdated := DoTestUpdateProfile(t, profileCreateResponse.ID) //U
	updatedProfile := DoTestGetProfile(t, profileUpdated.ID)           //Updating Profiles create new Profiles
	assert.Equal(t, "name_update", updatedProfile.Name)
	DoTestDeleteProfile(t, profileCreateResponse.ID) //D
}

func DoTestGetProfiles(t *testing.T) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetProfiles(nil)
	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res[0].ID %s\n", res[0].ID)

	assert.NotEmpty(t, res[0].ID, "expecting non-empty Rules")
}

func DoTestGetProfile(t *testing.T, ProfileID string) (Profile Profile) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.GetProfile(ProfileID, nil)

	assert.Equal(t, 200, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, res.ID, ProfileID)
	return res
}

func DoTestUpdateProfile(t *testing.T, ProfileID string) (Profile Profile) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	update := ProfileUpdateRequest{
		Rules: &Rules{
			Inbound: []*Rule{
				&Rule{
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
			Outbound: []*Rule{},
		},
		Name:    "name_update",
		Version: "version_update",
	}
	res, statusCode, err := c.UpdateProfile(ProfileID, update, nil)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.NotEmpty(t, res.ID, "expecting non-empty ID")
	assert.Equal(t, "name_update", res.Name)
	assert.Equal(t, 200, statusCode)
	return res
}

func DoTestCreateProfile(t *testing.T) (profile Profile) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	newProfile := ProfileCreateRequest{
		Rules: &Rules{
			Inbound: []*Rule{
				&Rule{
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
			Outbound: []*Rule{},
		},
		Name:    "name",
		Version: "version",
	}

	res, statusCode, err := c.CreateProfile(newProfile, nil)
	assert.Equal(t, 201, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("err: %v", err)
	t.Logf("res: %+v\n", res)
	return res
}

func DoTestDeleteProfile(t *testing.T, ProfileID string) {
	c := NewClient(os.Getenv("DOG_API_TOKEN"), os.Getenv("DOG_API_ENDPOINT"))

	res, statusCode, err := c.DeleteProfile(ProfileID, nil)
	assert.Equal(t, 204, statusCode)
	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, res, "expecting non-nil result")
	t.Logf("res: %+v\n", res)

	assert.Empty(t, res, "expecting empty response")
}
