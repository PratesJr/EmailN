package campaign

import (
	"emailn/internal/contract"
	"emailn/internal/domain/exceptions"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(payload contract.NewCampaign) (string, error) {

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
func (s *Service) Find() ([]Campaign, error) {

	result, err := s.Repository.Find()

	if err != nil {
		return nil, exceptions.UnkownErrror
	}
	return result, nil

}
