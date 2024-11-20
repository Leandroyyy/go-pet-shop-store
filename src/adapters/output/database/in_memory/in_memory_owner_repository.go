package output_database

import (
	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/entities"
)

type InMemoryOwnerRepository struct {
}

var owners []entities.Owner

func (o InMemoryOwnerRepository) Save(owner *entities.Owner) error {

	owners = append(owners, *owner)

	return nil
}

func (o InMemoryOwnerRepository) FindByDocument(document string) *entities.Owner {

	for _, owner := range owners {
		if owner.Document == document {
			return &owner
		}
	}

	return nil
}

func (o InMemoryOwnerRepository) FindById(id string) *entities.Owner {

	for _, owner := range owners {
		if owner.Id == id {
			return &owner
		}
	}

	return nil
}

func (o InMemoryOwnerRepository) Edit(owner *entities.Owner) error {

	var itemIndex int

	for index, ownerToEdit := range owners {
		if owner.Id == ownerToEdit.Id {
			itemIndex = index
			break
		}
	}

	owners[itemIndex] = *owner

	return nil
}
