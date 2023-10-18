package campaign

import (
	"emailn/internal/validators"
	"github.com/rs/xid"
	"time"
)

type Contact struct {
	Email string `validate:"email"`
}
type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=255"`
	Contacts  []Contact `validate:"min=1,dive"`
}

func NewCampaign(name string, content string, email []string) (*Campaign, error) {

	contacts := make([]Contact, len(email))

	for index, value := range email {
		contacts[index].Email = value
	}
	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}

	err := validators.Validate(campaign)

	if err != nil {
		return nil, err
	}
	return campaign, nil
}
