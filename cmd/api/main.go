package main

import (
	"email-n/internal/domain/campaign"
	"email-n/internal/endpoints"
	"email-n/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	handler := endpoints.Handler{
		CampaignService: campaignService,
	}

	r.Post("/campaigns", handler.CampaignPost)
	r.Get("/campaigns", handler.CampaignGet)

	http.ListenAndServe(":3000", r)
}
