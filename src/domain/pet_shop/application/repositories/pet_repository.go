package repositories

import (
	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/entities"
)

type PetRepository interface {
	Save(pet *entities.Pet) error
}
