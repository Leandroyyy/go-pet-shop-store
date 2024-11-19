package repositories

import (
	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/entities"
)

type OwnerRepository interface {
	Save(owner *entities.Owner) error
	FindByDocument(document string) *entities.Owner
	// FindyById(id string) (entities.Owner, error)
	// Edit(owner entities.Owner) error
}