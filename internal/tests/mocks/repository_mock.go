package mocks

import (
	"emailn/internal/domain/campaign"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Update(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *RepositoryMock) Create(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}
func (r *RepositoryMock) Find() ([]campaign.Campaign, error) {
	//args := r.Called(campaign)
	return nil, nil
}

func (r *RepositoryMock) FindById(id string) (*campaign.Campaign, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*campaign.Campaign), nil
}

func (r *RepositoryMock) Delete(campaign *campaign.Campaign) error {
	//args := r.Called(campaign)
	return nil
}
