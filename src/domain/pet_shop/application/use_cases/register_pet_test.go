package use_cases

import (
	"testing"
	"time"

	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/entities"
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

	// t.Run("should return not found error when owner doesn't exists", func(t *testing.T) {

	// 	mockOwnerRepo.On("FindById", "123123").Return(nil)

	// 	_, err := useCase.Execute(RegisterPetUseCaseRequest{
	// 		OwnerId:  "123123",
	// 		Name:     "Thor",
	// 		Birthday: "2023-10-10",
	// 		Breed:    "lhasa",
	// 		Gender:   "male",
	// 		Kind:     "dog",
	// 	})

	// 	assert.Error(t, err)

	// 	var notFoundErr *enterprise_errors.NotFoundError

	// 	assert.True(t, errors.As(err, &notFoundErr), "Expected error of type NotFoundError")
	// 	assert.Equal(t, "Owner doesn't exists", notFoundErr.Error())

	// 	mockOwnerRepo.AssertCalled(t, "FindById", "123123")
	// 	mockOwnerRepo.AssertNotCalled(t, "Edit")
	// 	mockPetRepo.AssertNotCalled(t, "Save")
	// })

	t.Run("should be able to register a pet", func(t *testing.T) {
		ownerId := "123123"
		owner := entities.NewOwner(entities.OwnerProps{
			Name:     "john due",
			Document: "123213123",
			Birthday: time.Now(),
			Email:    "john@email.cm",
		}, &ownerId)

		mockOwnerRepo.On("FindById", "123123").Return(owner)
		// mockOwnerRepo.On("Edit", mock.Anything).Return(nil)
		// mockPetRepo.On("Save", mock.Anything).Return(nil)

		result, err := useCase.Execute(RegisterPetUseCaseRequest{
			OwnerId:  "123123",
			Name:     "Thor",
			Birthday: "2023-10-10",
			Breed:    "lhasa",
			Gender:   "male",
			Kind:     "dog",
		})

		assert.NoError(t, err)
		assert.NotNil(t, result)
		// assert.NotEmpty(t, result.Id, "Expected owner ID to be generated")
		// assert.Equal(t, input.Name, result.Name)
		// assert.Equal(t, input.Document, result.Document)
		// assert.Equal(t, input.Email, result.Email)

		// mockRepo.AssertCalled(t, "FindByDocument", "1122334455")
		// mockRepo.AssertCalled(t, "Save", mock.Anything)
	})

}
