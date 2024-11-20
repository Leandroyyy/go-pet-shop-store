package output_database

import (
	"testing"
	"time"

	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/entities"
	"github.com/stretchr/testify/assert"
)

var petRepository = InMemoryPetRepository{}

func TestInMemoryPetRepository_Save(t *testing.T) {
	pets = []entities.Pet{}

	pet := entities.NewPet(entities.PetProps{
		Name:     "thor",
		Birthday: time.Now(),
		Breed:    "lhasa",
		Gender:   "male",
		Kind:     "dog",
	}, nil)

	err := petRepository.Save(&pet)

	assert.NoError(t, err)

	assert.Len(t, pets, 1)
	assert.Equal(t, pet, pets[0])
}
