package mocks

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (r *ServiceMock) Find() ([]campaign.Campaign, error) {
	return nil, nil
}

func (r *ServiceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := r.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (r *ServiceMock) FindById(id string) (*campaign.Campaign, error) {
	args := r.Called(id)
	return args.Get(0).(*campaign.Campaign), args.Error(1)
}
