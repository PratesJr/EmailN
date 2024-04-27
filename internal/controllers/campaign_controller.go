package controllers

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"emailn/internal/domain/exceptions"
	"errors"
	"github.com/go-chi/chi/v5"
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
	_contract.CreatedBy = r.Context().Value("email").(string)

	if err != nil {
		if errors.Is(err, exceptions.DbError) {
			return nil, 500, exceptions.DbError
		} else {
			return nil, 400, err
		}

	}

	return map[string]string{"id": id}, 201, nil
}

func (h CampaignHandler) GetCampaignById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	result, err := h.CampaignService.FindById(id)
	return result, 200, err
}
func (h CampaignHandler) CancelCampaign(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	result, err := h.CampaignService.Cancel(id)
	return result, 200, err
}

func (h CampaignHandler) Delete(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	err := h.CampaignService.Delete(id)
	return nil, 200, err
}
