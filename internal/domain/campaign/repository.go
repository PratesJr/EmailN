package campaign

type Repository interface {
	Update(campaign *Campaign) error
	Find() ([]Campaign, error)
	FindById(id string) (*Campaign, error)
	Delete(campaign *Campaign) error
	Create(campaign *Campaign) error
}
