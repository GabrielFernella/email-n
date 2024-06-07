package endpoints

import "email-n/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
