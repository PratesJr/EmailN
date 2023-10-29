package controllers

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"emailn/internal/domain/exceptions"
	"errors"
	"github.com/go-chi/render"
	"net/http"
)

type CampaignHandler struct {
	CampaignService campaign.Service
}

func (h CampaignHandler) PostCampaign(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var _contract contract.NewCampaign

	errDecode := render.DecodeJSON(r.Body, &_contract)
	if errDecode != nil {
		return nil, 500, errDecode
	}
	id, err := h.CampaignService.Create(_contract)

	if err != nil {
		if errors.Is(err, exceptions.DbError) {
			return nil, 500, exceptions.DbError
		} else {
			return nil, 400, err
		}

	}

	return map[string]string{"id": id}, 201, nil
}

func (h CampaignHandler) GetCampaign(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	result, err := h.CampaignService.Repository.Find()
	return result, 200, err
}
