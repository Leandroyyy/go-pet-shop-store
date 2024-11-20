package output_database

import (
	"testing"
	"time"

	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/entities"
	"github.com/stretchr/testify/assert"
)

var ownerRepository = InMemoryOwnerRepository{}

func TestInMemoryOwnerRepository_Save(t *testing.T) {
	owners = []entities.Owner{}

	owner := entities.NewOwner(entities.OwnerProps{
		Name:     "Leandro",
		Document: "1231232131",
		Birthday: time.Now(),
		Email:    "leandro@email.com",
	}, nil)

	err := ownerRepository.Save(&owner)

	assert.NoError(t, err)

	assert.Len(t, owners, 1)
	assert.Equal(t, owner, owners[0])
}

func TestInMemoryOwnerRepository_FindByDocument(t *testing.T) {

	t.Run("should be able to find an owner", func(t *testing.T) {

		ownerMocked := entities.NewOwner(entities.OwnerProps{
			Name:     "Leandro",
			Document: "1231232131",
			Birthday: time.Now(),
			Email:    "leandro@email.com",
		}, nil)

		owners = append(owners, ownerMocked)

		owner := ownerRepository.FindByDocument("1231232131")

		assert.Equal(t, ownerMocked.Name, owner.Name)
		assert.Equal(t, ownerMocked.Document, owner.Document)
		assert.Equal(t, ownerMocked.Email, owner.Email)

		assert.NotEmpty(t, owner.Id)
	})

}

func TestInMemoryOwnerRepository_FindById(t *testing.T) {

	t.Run("should be able to find an owner", func(t *testing.T) {

		ownerId := "1"
		ownerMocked := entities.NewOwner(entities.OwnerProps{
			Name:     "Leandro",
			Document: "1231232131",
			Birthday: time.Now(),
			Email:    "leandro@email.com",
		}, &ownerId)

		owners = append(owners, ownerMocked)

		owner := ownerRepository.FindById(ownerId)

		assert.Equal(t, ownerMocked.Name, owner.Name)
		assert.Equal(t, ownerMocked.Document, owner.Document)
		assert.Equal(t, ownerMocked.Email, owner.Email)
		assert.Equal(t, owner.Id, ownerId)
	})

}

func TestInMemoryOwnerRepository_Edit(t *testing.T) {

	t.Run("should edit an owner", func(t *testing.T) {
		ownerId := "1"

		owners = nil

		owners = append(owners, entities.NewOwner(entities.OwnerProps{
			Name:     "John Doe",
			Document: "123456789",
			Email:    "john.doe@example.com",
			Birthday: time.Now(),
		}, &ownerId))

		updatedOwner := entities.NewOwner(entities.OwnerProps{
			Name:     "John Smith",
			Document: "123456789",
			Email:    "john.smith@example.com",
			Birthday: time.Now(),
		}, &ownerId)

		err := ownerRepository.Edit(&updatedOwner)
		assert.NoError(t, err, "Expected no error when editing an existing owner")

		assert.Equal(t, "John Smith", owners[0].Name)
		assert.Equal(t, "john.smith@example.com", owners[0].Email)
	})

}
