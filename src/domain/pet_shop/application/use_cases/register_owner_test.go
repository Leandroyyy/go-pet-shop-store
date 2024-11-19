package use_cases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/entities"
	enterprise_errors "github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/errors"
)

// Mock do OwnerRepository
type MockOwnerRepository struct {
	mock.Mock
}

func (m *MockOwnerRepository) FindByDocument(document string) *entities.Owner {
	args := m.Called(document)
	if owner, ok := args.Get(0).(*entities.Owner); ok {
		return owner
	}
	return nil
}

func (m *MockOwnerRepository) Save(owner *entities.Owner) error {
	args := m.Called(owner)
	return args.Error(0)
}

func TestRegisterOwnerUseCase_Execute(t *testing.T) {
	mockRepo := new(MockOwnerRepository)
	useCase := &RegisterOwnerUseCase{ownerRepository: mockRepo}

	t.Run("should return conflict error if document already exists", func(t *testing.T) {
		mockRepo.On("FindByDocument", "123456789").Return(&entities.Owner{})

		input := RegisterOwnerUseCaseRequest{
			Name:     "John Doe",
			Document: "123456789",
			Birthday: "2000-01-01 00:00:00",
			Email:    "john@example.com",
		}

		result, err := useCase.Execute(input)

		assert.Nil(t, result)
		assert.ErrorIs(t, err, enterprise_errors.NewConflictError("Document already exists"))
		mockRepo.AssertCalled(t, "FindByDocument", "123456789")
	})

	t.Run("should return owner if input is valid", func(t *testing.T) {
		mockRepo.On("FindByDocument", "1122334455").Return(nil)

		input := RegisterOwnerUseCaseRequest{
			Name:     "Alice Doe",
			Document: "1122334455",
			Birthday: "2000-01-01 00:00:00",
			Email:    "alice@example.com",
		}

		result, err := useCase.Execute(input)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, input.Name, result.Name)
		assert.Equal(t, input.Document, result.Document)
		assert.Equal(t, input.Email, result.Email)
		mockRepo.AssertCalled(t, "FindByDocument", "1122334455")
	})
}
