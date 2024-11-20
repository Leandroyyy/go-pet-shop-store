package use_cases

import (
	"time"

	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/application/repositories"
	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/entities"
	enterprise_errors "github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/errors"
)

type RegisterOwnerUseCaseRequest struct {
	Name     string
	Document string
	Birthday string
	Email    string
}

func NewRegisterOwnerUseCase(ownerRepository repositories.OwnerRepository) RegisterOwnerUseCase {
	return RegisterOwnerUseCase{
		ownerRepository: ownerRepository,
	}
}

type RegisterOwnerUseCase struct {
	ownerRepository repositories.OwnerRepository
}

func (r *RegisterOwnerUseCase) Execute(input RegisterOwnerUseCaseRequest) (*entities.Owner, error) {
	documentAlreadyExists := r.ownerRepository.FindByDocument(input.Document)

	if documentAlreadyExists != nil {
		return nil, enterprise_errors.NewConflictError("Document already exists")
	}

	layout := "2006-01-02 15:04:05"

	parsedBirthday, _ := time.Parse(layout, input.Birthday)

	owner := entities.NewOwner(entities.OwnerProps{
		Name:     input.Name,
		Document: input.Document,
		Birthday: parsedBirthday,
		Email:    input.Email,
	}, nil)

	r.ownerRepository.Save(&owner)

	return &owner, nil
}
