package tests

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/enums"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCampaign(t *testing.T) {
	var (
		name      = "Campaign X"
		content   = "Body Hi!"
		contacts  = []string{"email1@e.com", "email2@e.com"}
		createdBy = "mail@mail.com"
		fake      = faker.New()
	)

	t.Run("should create a new campaign", func(t *testing.T) {
		assertions := assert.New(t)
		data, _ := campaign.NewCampaign(name, content, contacts, createdBy)

		assertions.Equal(data.Name, name)
		assertions.Equal(data.Content, content)
		assertions.Equal(len(data.Contacts), len(contacts))
		assertions.Equal(data.CreatedBy, createdBy)

	})

	t.Run("should create and id can not be null", func(t *testing.T) {
		assertions := assert.New(t)

		data, _ := campaign.NewCampaign(name, content, contacts, createdBy)

		assertions.NotNil(data.ID)

	})
	t.Run("should add the correct value to CreatedOn field", func(t *testing.T) {
		assertions := assert.New(t)
		now := time.Now().Add(-time.Minute)

		data, _ := campaign.NewCampaign(name, content, contacts, createdBy)

		assertions.Greater(data.CreatedOn, now)

	})
	t.Run("should validate the min length of name", func(t *testing.T) {
		assertions := assert.New(t)

		_, err := campaign.NewCampaign("", content, contacts, createdBy)

		assertions.Equal("Name is required with min 5", err.Error())

	})
	t.Run("should validate the max length of name", func(t *testing.T) {
		assertions := assert.New(t)

		_, err := campaign.NewCampaign(fake.Lorem().Text(30), content, contacts, createdBy)

		assertions.Equal("Name is required with max 24", err.Error())

	})
	t.Run("should validate the min length of contacts array", func(t *testing.T) {
		assertions := assert.New(t)

		_, err := campaign.NewCampaign(name, content, nil, createdBy)

		assertions.Equal("Contacts is required with min 1", err.Error())

	})
	t.Run("should validate if the passed emails are valid", func(t *testing.T) {
		assertions := assert.New(t)

		_, err := campaign.NewCampaign(name, content, []string{"email_invalid"}, createdBy)

		assertions.Equal("Email is not valid", err.Error())

	})
	t.Run("should create a campaign with status PENDING", func(t *testing.T) {
		assertions := assert.New(t)

		data, _ := campaign.NewCampaign(name, content, []string{"mail@mail.com"}, createdBy)

		assertions.Equal(enums.Pending, data.Status)

	})
	t.Run("should throw error if createby is empty", func(t *testing.T) {
		assertions := assert.New(t)

		_, err := campaign.NewCampaign(name, content, []string{"mail@mail.com"}, "")

		assertions.Equal("CreatedBy is not valid", err.Error())

	})

}
