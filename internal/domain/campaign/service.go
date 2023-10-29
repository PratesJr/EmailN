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
