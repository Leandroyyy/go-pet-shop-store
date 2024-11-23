package use_cases

import (
	"errors"
	"testing"

	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/entities"
	enterprise_errors "github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/errors"
	test_factories "github.com/leandroyyy/poc-golang/tests/factories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPetRepository struct {
	mock.Mock
}

func (m *MockPetRepository) Save(pet *entities.Pet) error {
	args := m.Called(pet)
	return args.Error(0)
}

func TestRegisterPetUseCase_Execute(t *testing.T) {
	mockOwnerRepo := new(MockOwnerRepository)
	mockPetRepo := new(MockPetRepository)

	useCase := &RegisterPetUseCase{
		ownerRepository: mockOwnerRepo,
		petRepository:   mockPetRepo,
	}

	t.Run("should return not found error when owner doesn't exists", func(t *testing.T) {

		mockOwnerRepo.ExpectedCalls = nil
		mockOwnerRepo.Calls = nil

		mockOwnerRepo.On("FindById", "123123").Return(nil)

		_, err := useCase.Execute(RegisterPetUseCaseRequest{
			OwnerId:  "123123",
			Name:     "Thor",
			Birthday: "2023-10-10",
			Breed:    "lhasa",
			Gender:   "male",
			Kind:     "dog",
		})

		assert.Error(t, err)

		var notFoundErr *enterprise_errors.NotFoundError

		assert.True(t, errors.As(err, &notFoundErr), "Expected error of type NotFoundError")
		assert.Equal(t, "Owner doesn't exists", notFoundErr.Error())

		mockOwnerRepo.AssertCalled(t, "FindById", "123123")
		mockOwnerRepo.AssertNotCalled(t, "Edit")
		mockPetRepo.AssertNotCalled(t, "Save")
	})

	t.Run("should be able to register a pet", func(t *testing.T) {
		mockOwnerRepo.ExpectedCalls = nil
		mockOwnerRepo.Calls = nil

		ownerId := "123123"
		owner := *test_factories.MakeOwner()
		owner.Id = ownerId

		mockOwnerRepo.On("FindById", ownerId).Return(&owner)
		mockOwnerRepo.On("Edit", mock.Anything).Return(nil)
		mockPetRepo.On("Save", mock.Anything).Return(nil)

		input := RegisterPetUseCaseRequest{
			OwnerId:  "123123",
			Name:     "Thor",
			Birthday: "2023-10-10",
			Breed:    "lhasa",
			Gender:   "male",
			Kind:     "dog",
		}

		result, err := useCase.Execute(input)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.Id, "Expected owner ID to be generated")
		assert.Equal(t, result.Name, input.Name)
		assert.Equal(t, result.Breed, input.Breed)
		assert.Equal(t, result.Gender, input.Gender)
		assert.Equal(t, result.Kind, input.Kind)

		mockOwnerRepo.AssertCalled(t, "FindById", "123123")
		mockOwnerRepo.AssertCalled(t, "Edit", mock.Anything)
		mockPetRepo.AssertCalled(t, "Save", mock.Anything)
	})

}
