package output_database

import (
	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/entities"
)

type InMemoryPetRepository struct {
}

var pets []entities.Pet

func (o InMemoryPetRepository) Save(pet *entities.Pet) error {

	pets = append(pets, *pet)

	return nil
}
