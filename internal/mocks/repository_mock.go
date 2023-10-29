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

func (r *RepositoryMock) Find() []campaign.Campaign {
	//args := r.Called(campaign)
	return nil
}
