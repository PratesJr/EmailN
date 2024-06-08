package campaign

import (
	"emailn/internal/contract"
	"emailn/internal/domain/exceptions"
	"emailn/internal/infra/mail"
	"errors"
	"gorm.io/gorm"
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

	campaign, err := NewCampaign(payload.Name, payload.Content, payload.Email, payload.CreatedBy)
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
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.UnkownErrror
		}
		return nil, err
	}
	return result, nil

}

func (s *ServiceImpl) Cancel(id string) (*Campaign, error) {

	result, err := s.Repository.FindById(id)

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.UnkownErrror
		}
		return nil, err
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
		if !errors.Is(errFind, gorm.ErrRecordNotFound) {
			return exceptions.UnkownErrror
		}
		return errFind
	}

	err := s.Repository.Delete(result)

	return err
}

func (s *ServiceImpl) Start(id string) error {
	result, errFind := s.Repository.FindById(id)

	if errFind != nil {
		if !errors.Is(errFind, gorm.ErrRecordNotFound) {
			return exceptions.UnkownErrror
		}
		return errFind
	}

	return mail.SendMail(*result)
}
