package mocks

import (
	"emailn/internal/domain/campaign"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Save(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *RepositoryMock) Find() ([]campaign.Campaign, error) {
	//args := r.Called(campaign)
	return nil, nil
}

func (r *RepositoryMock) FindById(id string) (*campaign.Campaign, error) {
	args := r.Called(id)
	return args.Get(0).(*campaign.Campaign), args.Error(1)
}
