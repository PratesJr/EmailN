package main

import (
	"emailn/internal/controllers"
	"emailn/internal/domain/campaign"
	"emailn/internal/infra/database/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignHandler := controllers.CampaignHandler{
		CampaignService: &campaign.ServiceImpl{
			Repository: &repositories.CampaignRepository{},
		},
	}
	r.Post("/campaign", controllers.HandlerError(campaignHandler.PostCampaign))
	r.Get("/campaign/{id}", controllers.HandlerError(campaignHandler.GetCampaignById))

	err := http.ListenAndServe(":3000", r)

	if err != nil {
		return
	}

}
