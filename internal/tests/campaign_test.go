package tests

import (
	campaign2 "emailn/internal/domain/campaign"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	name     = "Campaign X"
	content  = "body"
	contacts = []string{"mail@email.com", "email2@mail.com"}
)

func TestNewCampaign(t *testing.T) {
	assertions := assert.New(t)

	campaign, _ := campaign2.NewCampaign(name, content, contacts)

	assertions.Equal(campaign.Name, name)
	assertions.Equal(campaign.Content, content)
	assertions.Equal(len(campaign.Contacts), len(contacts))
}
func TestIDIsNotNull(t *testing.T) {
	assertions := assert.New(t)

	campaign, _ := campaign2.NewCampaign(name, content, contacts)

	assertions.NotNil(campaign.ID)
}
func TestNonParameterizableValues(t *testing.T) {
	assertions := assert.New(t)

	campaign, _ := campaign2.NewCampaign(name, content, contacts)

	assertions.NotNil(campaign.ID)
	assertions.GreaterOrEqual(campaign.CreatedOn, time.Now().Add(-time.Minute))
}

func TestMustValidateName(t *testing.T) {
	assertions := assert.New(t)

	_, exception := campaign2.NewCampaign("", content, contacts)

	assertions.Equal("name can not be empty", exception.Error())
}
func TestMustValidateContent(t *testing.T) {
	assertions := assert.New(t)

	_, exception := campaign2.NewCampaign(name, "", contacts)

	assertions.Equal("content is required", exception.Error())
}

func TestMustValidateEmails(t *testing.T) {
	assertions := assert.New(t)

	_, exception := campaign2.NewCampaign(name, content, []string{})

	assertions.Equal("emails are required", exception.Error())
}
