package repositories

import (
	"emailn/internal/domain/campaign"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (c *CampaignRepository) Update(campaign *campaign.Campaign) error {
	tx := c.Db.Save(campaign)

	return tx.Error
}

func (c *CampaignRepository) Create(campaign *campaign.Campaign) error {
	tx := c.Db.Create(campaign)

	return tx.Error
}

func (c *CampaignRepository) Find() ([]campaign.Campaign, error) {
	var result []campaign.Campaign
	tx := c.Db.Find(&result)
	return result, tx.Error
}
func (c *CampaignRepository) FindById(id string) (*campaign.Campaign, error) {
	var result campaign.Campaign
	tx := c.Db.First(&result, "id=?", id)
	return &result, tx.Error
}
func (c *CampaignRepository) Delete(campaign *campaign.Campaign) error {
	tx := c.Db.Delete(&campaign)
	return tx.Error
}
