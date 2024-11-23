package output_database

import (
	"testing"

	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/entities"
	test_factories "github.com/leandroyyy/poc-golang/tests/factories"
	"github.com/stretchr/testify/assert"
)

var ownerRepository = InMemoryOwnerRepository{}

func TestInMemoryOwnerRepository_Save(t *testing.T) {
	owners = []entities.Owner{}

	owner := test_factories.MakeOwner()

	err := ownerRepository.Save(owner)

	assert.NoError(t, err)

	assert.Len(t, owners, 1)
	assert.Equal(t, owner, &owners[0])
}

func TestInMemoryOwnerRepository_FindByDocument(t *testing.T) {

	t.Run("should be able to find an owner", func(t *testing.T) {
		owners = nil

		ownerMocked := test_factories.MakeOwner()

		owners = append(owners, *ownerMocked)

		owner := ownerRepository.FindByDocument("123132213")

		assert.Equal(t, &ownerMocked.Name, &owner.Name)
		assert.Equal(t, &ownerMocked.Document, &owner.Document)
		assert.Equal(t, &ownerMocked.Email, &owner.Email)

		assert.NotEmpty(t, owner.Id)
	})

}

func TestInMemoryOwnerRepository_FindById(t *testing.T) {

	t.Run("should be able to find an owner", func(t *testing.T) {
		owners = nil

		ownerMocked := test_factories.MakeOwner()

		owners = append(owners, *ownerMocked)

		owner := ownerRepository.FindById(ownerMocked.Id)

		assert.Equal(t, ownerMocked.Name, owner.Name)
		assert.Equal(t, ownerMocked.Document, owner.Document)
		assert.Equal(t, ownerMocked.Email, owner.Email)
	})

}

func TestInMemoryOwnerRepository_Edit(t *testing.T) {

	t.Run("should edit an owner", func(t *testing.T) {
		ownerId := "1"

		owners = nil

		owners = append(owners, *test_factories.MakeOwner())

		updatedOwner := test_factories.MakeOwner()
		updatedOwner.Id = ownerId

		err := ownerRepository.Edit(updatedOwner)
		assert.NoError(t, err, "Expected no error when editing an existing owner")

		assert.Equal(t, "john due", owners[0].Name)
		assert.Equal(t, "john@due.com", owners[0].Email)
	})

}
