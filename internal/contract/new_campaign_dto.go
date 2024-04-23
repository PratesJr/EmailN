package contract

type NewCampaign struct {
	ID        string
	Name      string
	Content   string
	Email     []string
	CreatedBy string
}
