package campaign

import (
	"emailn/internal/contract"
	"emailn/internal/domain/exceptions"
)

type Service interface {
	Create(campaign contract.NewCampaign) (string, error)
	Find() ([]Campaign, error)
	FindById(id string) (*Campaign, error)
}
type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(payload contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(payload.Name, payload.Content, payload.Email)
	if err != nil {
		return "", err

	}
	err = s.Repository.Save(campaign)
	if err != nil {
		return "", exceptions.DbError
	}

	return campaign.ID, err
}
func (s *ServiceImpl) Find() ([]Campaign, error) {

	result, err := s.Repository.Find()

	if err != nil {
		return nil, exceptions.UnkownErrror
	}
	return result, nil

}
func (s *ServiceImpl) FindById(id string) (*Campaign, error) {

	result, err := s.Repository.FindById(id)

	if err != nil {
		return nil, exceptions.UnkownErrror
	}
	return result, nil

}
