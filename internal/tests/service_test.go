package tests

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"emailn/internal/domain/exceptions"
	"emailn/internal/mocks"
	"github.com/jaswdr/faker"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService(t *testing.T) {
	var (
		fake        = faker.New()
		newCampaign = contract.NewCampaign{
			Name:    fake.Lorem().Text(13),
			Content: "content",
			Email:   []string{"teste@mail.com", "testeee@mail.com"},
		}
		service = campaign.Service{}
	)
	t.Run("should create a campaign with no errors", func(t *testing.T) {
		assertions := assert.New(t)
		repoMock := new(mocks.RepositoryMock)

		repoMock.On("Save", mock.MatchedBy(func(c *campaign.Campaign) bool {
			return c.Name == newCampaign.Name && c.Content == newCampaign.Content && len(c.Contacts) == len(newCampaign.Email)
		})).Return(nil)

		service.Repository = repoMock
		id, err := service.Create(newCampaign)

		assertions.NotNil(id)
		assertions.Nil(err)
	})

	t.Run("should save campaign", func(t *testing.T) {
		repoMock := new(mocks.RepositoryMock)

		repoMock.On("Save", mock.MatchedBy(func(c *campaign.Campaign) bool {
			return c.Name == newCampaign.Name && c.Content == newCampaign.Content && len(c.Contacts) == len(newCampaign.Email)
		})).Return(nil)

		service.Repository = repoMock

		service.Create(newCampaign)

		repoMock.AssertExpectations(t)
	})
	t.Run("should return the 'min' exception", func(t *testing.T) {
		newCampaign.Name = ""
		assertions := assert.New(t)
		repoMock := new(mocks.RepositoryMock)

		repoMock.On("Save", mock.MatchedBy(func(c *campaign.Campaign) bool {
			return c.Name == newCampaign.Name && c.Content == newCampaign.Content && len(c.Contacts) == len(newCampaign.Email)
		})).Return(nil)

		service.Repository = repoMock
		_, err := service.Create(newCampaign)

		assertions.NotNil(err)
		assertions.Equal("Name is required with min 5", err.Error())
	})
	t.Run("should return the correct db exception", func(t *testing.T) {
		assertions := assert.New(t)
		repoMock := new(mocks.RepositoryMock)
		nCampaign := contract.NewCampaign{
			Name:    fake.Lorem().Text(13),
			Content: "content",
			Email:   []string{"teste@mail.com", "testeee@mail.com"},
		}
		repoMock.On("Save", mock.Anything).Return(exceptions.DbError)

		service.Repository = repoMock
		_, err := service.Create(nCampaign)

		assertions.NotNil(err)
		assertions.Equal(exceptions.DbError.Error(), err.Error())
	})
	t.Run("should return the 'max' exception", func(t *testing.T) {
		newCampaign.Name = fake.Lorem().Text(25)
		assertions := assert.New(t)
		repoMock := new(mocks.RepositoryMock)

		repoMock.On("Save", mock.MatchedBy(func(c *campaign.Campaign) bool {
			return c.Name == newCampaign.Name && c.Content == newCampaign.Content && len(c.Contacts) == len(newCampaign.Email)
		})).Return(nil)

		service.Repository = repoMock
		_, err := service.Create(newCampaign)

		assertions.NotNil(err)
		assertions.Equal("Name is required with max 24", err.Error())
	})
	t.Run("should validate contacts required", func(t *testing.T) {
		assertions := assert.New(t)

		_, err := campaign.NewCampaign(
			fake.Lorem().Text(13),
			fake.Lorem().Text(100),
			[]string{})

		assertions.NotNil(err)
		assertions.Equal("Contacts is required with min 1", err.Error())
	})
	t.Run("should validate contacts format", func(t *testing.T) {
		assertions := assert.New(t)

		_, err := campaign.NewCampaign(
			fake.Lorem().Text(13),
			fake.Lorem().Text(100),
			[]string{"email"})

		assertions.NotNil(err)
		assertions.Equal("Email is not valid", err.Error())
	})
}
