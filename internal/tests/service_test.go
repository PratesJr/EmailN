package tests

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"emailn/internal/domain/exceptions"
	"emailn/internal/enums"
	"emailn/internal/tests/mocks"
	"github.com/jaswdr/faker"
	"gorm.io/gorm"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService(t *testing.T) {
	var (
		fake        = faker.New()
		newCampaign = contract.NewCampaign{
			Name:      fake.Lorem().Text(13),
			Content:   "content",
			Email:     []string{"teste@mail.com", "testeee@mail.com"},
			CreatedBy: "mail@email.com.br",
		}
		service = campaign.ServiceImpl{}
	)
	t.Run("should create a campaign with no errors", func(t *testing.T) {
		assertions := assert.New(t)
		repoMock := new(mocks.RepositoryMock)

		repoMock.On("Create", mock.MatchedBy(func(c *campaign.Campaign) bool {
			return c.Name == newCampaign.Name && c.Content == newCampaign.Content && len(c.Contacts) == len(newCampaign.Email)
		})).Return(nil)

		service.Repository = repoMock
		id, err := service.Create(newCampaign)

		assertions.NotNil(id)
		assertions.Nil(err)
	})

	t.Run("should save campaign", func(t *testing.T) {
		repoMock := new(mocks.RepositoryMock)

		repoMock.On("Create", mock.MatchedBy(func(c *campaign.Campaign) bool {
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
			Name:      fake.Lorem().Text(13),
			Content:   "content",
			Email:     []string{"teste@mail.com", "testeee@mail.com"},
			CreatedBy: "mail@email.com",
		}
		repoMock.On("Create", mock.Anything).Return(exceptions.DbError)

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
			[]string{}, "")

		assertions.NotNil(err)
		assertions.Equal("Contacts is required with min 1", err.Error())
	})
	t.Run("should return the find by id result", func(t *testing.T) {
		assertions := assert.New(t)
		repoMock := new(mocks.RepositoryMock)
		service.Repository = repoMock
		nCampaign, _ := campaign.NewCampaign(
			fake.Lorem().Text(13),
			"content",
			[]string{"teste@mail.com", "testeee@mail.com"},
			"meail@gmail.com",
		)
		repoMock.On("FindById", mock.Anything).Return(nCampaign, nil)

		campaignReturned, _ := service.FindById(nCampaign.ID)
		assertions.Equal(
			nCampaign.ID,
			campaignReturned.ID,
		)
	})

	t.Run("should delete with no error", func(t *testing.T) {
		assertions := assert.New(t)
		repoMock := new(mocks.RepositoryMock)
		service.Repository = repoMock
		nCampaign, _ := campaign.NewCampaign(
			fake.Lorem().Text(13),
			"content",
			[]string{"teste@mail.com", "testeee@mail.com"},
			"mail@bol.com",
		)
		repoMock.On("FindById", mock.Anything).Return(nCampaign, nil)

		err := service.Delete(nCampaign.ID)
		assertions.Nil(err)

	})

	t.Run("should throw error before delete if record not found", func(t *testing.T) {
		assertions := assert.New(t)
		repoMock := new(mocks.RepositoryMock)
		service.Repository = repoMock

		repoMock.On("FindById", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

		err := service.Delete("34234345asr234")
		assertions.Equal(err.Error(), gorm.ErrRecordNotFound.Error())

	})

	t.Run("should cancel campaign with no errors", func(t *testing.T) {
		assertions := assert.New(t)
		repoMock := new(mocks.RepositoryMock)
		service.Repository = repoMock
		nCampaign, _ := campaign.NewCampaign(
			fake.Lorem().Text(13),
			"content",
			[]string{"teste@mail.com", "testeee@mail.com"},
			"mail@yahoo.com",
		)

		repoMock.On("FindById", mock.Anything).Return(nCampaign, nil)
		repoMock.On("Update", mock.Anything).Return(nil)

		canceled, err := service.Cancel("34234345asr234")

		assertions.Nil(err)
		assertions.Equal(canceled.Status, enums.Canceled)

	})

	t.Run("should cancel campaign with no errors", func(t *testing.T) {
		assertions := assert.New(t)
		repoMock := new(mocks.RepositoryMock)
		service.Repository = repoMock
		nCampaign, _ := campaign.NewCampaign(
			fake.Lorem().Text(13),
			"content",
			[]string{"teste@mail.com", "testeee@mail.com"},
			"mail@dock.com",
		)

		nCampaign.Status = enums.Started
		repoMock.On("FindById", mock.Anything).Return(nCampaign, nil)

		_, err := service.Cancel("34234345asr234")

		assertions.Equal("Unable to Update campaign in status:"+enums.Started, err.Error())

	})
}
