package repositories

import "emailn/internal/domain/campaign"

type CampaignRepository struct {
	campaigns []campaign.Campaign
	campaign  campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.campaigns = append(c.campaigns, *campaign)

	return nil
}

func (c *CampaignRepository) Find() ([]campaign.Campaign, error) {
	return c.campaigns, nil
}
func (c *CampaignRepository) FindById(id string) (*campaign.Campaign, error) {
	return nil, nil
}
