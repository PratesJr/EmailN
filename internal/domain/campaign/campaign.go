package campaign

import (
	"emailn/internal/enums"
	"emailn/internal/validators"
	"errors"
	"github.com/rs/xid"
	"time"
)

type Contact struct {
	ID         string `gorm:"size=50"`
	Email      string `validate:"email" gorm:"size=100"`
	CampaignId string `gorm:"size=50"`
}
type Campaign struct {
	ID        string    `validate:"required"  gorm:"size=50"`
	Name      string    `validate:"min=5,max=24" gorm:"size=24"`
	CreatedOn time.Time `validate:"required" `
	Content   string    `validate:"min=5,max=255"  gorm:"size=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `gorm:"size=20"`
}

func NewCampaign(name string, content string, email []string) (*Campaign, error) {

	contacts := make([]Contact, len(email))

	for index, value := range email {
		contacts[index].ID = xid.New().String()
		contacts[index].Email = value
	}
	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
		Status:    enums.Pending,
	}

	err := validators.Validate(campaign)

	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (c *Campaign) Cancel() error {

	if c.Status != enums.Pending {
		return errors.New("Unable to Update campaign in status:" + c.Status)
	}

	c.Status = enums.Canceled

	return nil
}
