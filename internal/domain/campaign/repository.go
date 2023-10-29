package campaign

type Repository interface {
	Save(campaign *Campaign) error
	Find() ([]Campaign, error)
}
