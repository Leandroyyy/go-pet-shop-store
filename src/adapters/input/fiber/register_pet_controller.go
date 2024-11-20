package input_fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/application/use_cases"
	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/enterprise/entities"
)

type registerPetDto struct {
	OwnerId  string `json:"ownerId"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Breed    string `json:"breed"`
	Gender   string `json:"gender"`
	Kind     string `json:"kind"`
}

type RegisterPetController struct {
	registerPetUseCase use_cases.RegisterPetUseCase
}

func NewRegisterPetController(registerPetUseCase use_cases.RegisterPetUseCase) RegisterPetController {
	return RegisterPetController{
		registerPetUseCase: registerPetUseCase,
	}
}

func (r RegisterPetController) Execute(c *fiber.Ctx) error {

	var petDto registerPetDto

	if err := c.BodyParser(&petDto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Payload")
	}

	pet, err := r.registerPetUseCase.Execute(use_cases.RegisterPetUseCaseRequest{
		Name:     petDto.Name,
		OwnerId:  petDto.OwnerId,
		Birthday: petDto.Birthday,
		Breed:    petDto.Breed,
		Gender:   entities.PetGender(petDto.Gender),
		Kind:     entities.PetKind(petDto.Kind),
	})

	if err != nil {
		return HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"id":       pet.Id,
		"name":     pet.Name,
		"birthday": pet.Birthday.Format("2006-10-20"),
		"breed":    pet.Breed,
		"gender":   pet.Gender,
		"kind":     pet.Kind,
	})
}
