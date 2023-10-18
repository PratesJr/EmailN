package tests

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"emailn/internal/exceptins"
	"emailn/internal/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaign = contract.NewCampaign{
		Name:    "Marcelo",
		Content: "content",
		Email:   []string{"teste@mail.com", "testeee@mail.com"},
	}
	service = campaign.Service{}
)

func TestService(t *testing.T) {
	t.Run("should create a campaign swith no errors", func(t *testing.T) {
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

		repoMock.On("Save", mock.Anything).Return(exceptins.DbError)

		service.Repository = repoMock
		_, err := service.Create(newCampaign)

		assertions.NotNil(err)
		assertions.True(errors.Is(exceptins.DbError, err))
	})
}
