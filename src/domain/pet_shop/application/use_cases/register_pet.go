package use_cases

import (
	"time"

	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/application/repositories"
	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/entities"
	enterprise_errors "github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/errors"
)

type RegisterPetUseCaseRequest struct {
	OwnerId  string
	Name     string
	Birthday string
	Breed    string
	Gender   entities.PetGender
	Kind     entities.PetKind
}

type RegisterPetUseCase struct {
	ownerRepository repositories.OwnerRepository
	petRepository   repositories.PetRepository
}

func NewRegisterPetUseCase(ownerRepository repositories.OwnerRepository, petRepository repositories.PetRepository) RegisterPetUseCase {
	return RegisterPetUseCase{
		ownerRepository: ownerRepository,
		petRepository:   petRepository,
	}
}

func (r RegisterPetUseCase) Execute(input RegisterPetUseCaseRequest) (*entities.Pet, error) {

	owner := r.ownerRepository.FindById(input.OwnerId)

	if owner == nil {
		return nil, enterprise_errors.NewNotFoundError("Owner doesn't exists")
	}

	layout := "2006-01-02 15:04:05"

	parsedBirthday, _ := time.Parse(layout, input.Birthday)

	pet := entities.NewPet(entities.PetProps{
		Name:     input.Name,
		Birthday: parsedBirthday,
		Breed:    input.Breed,
		Gender:   input.Gender,
		Kind:     input.Kind,
	}, nil)

	r.petRepository.Save(&pet)

	owner.RegisterPet(pet)

	r.ownerRepository.Edit(owner)

	return &pet, nil
}
