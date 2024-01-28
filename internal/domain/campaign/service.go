package campaign

import (
	"emailn/internal/contract"
	"emailn/internal/domain/exceptions"
)

type Service interface {
	Create(campaign contract.NewCampaign) (string, error)
	Find() ([]Campaign, error)
	FindById(id string) (*Campaign, error)
	Cancel(id string) (*Campaign, error)
	Delete(id string) error
}
type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(payload contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(payload.Name, payload.Content, payload.Email)
	if err != nil {
		return "", err

	}
	err = s.Repository.Create(campaign)
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

func (s *ServiceImpl) Cancel(id string) (*Campaign, error) {

	result, err := s.Repository.FindById(id)

	if err != nil {
		return nil, exceptions.UnkownErrror
	}

	ex := result.Cancel()

	if ex != nil {
		return nil, ex
	}

	e := s.Repository.Update(result)

	if e != nil {
		return nil, e
	}
	return result, nil

}

func (s *ServiceImpl) Delete(id string) error {
	result, errFind := s.Repository.FindById(id)

	if errFind != nil {
		return errFind
	}

	err := s.Repository.Delete(result)

	return err
}
