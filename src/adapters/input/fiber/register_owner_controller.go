package input_fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/application/use_cases"
)

type registerOwnerDto struct {
	Name     string `json:"name"`
	Document string `json:"document"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
}

type RegisterOwnerController struct {
	registerOwnerUseCase use_cases.RegisterOwnerUseCase
}

func NewRegisterOwnerController(registerOwnerUseCase use_cases.RegisterOwnerUseCase) RegisterOwnerController {
	return RegisterOwnerController{
		registerOwnerUseCase: registerOwnerUseCase,
	}
}

func (r RegisterOwnerController) Execute(c *fiber.Ctx) error {

	var ownerDto registerOwnerDto

	if err := c.BodyParser(&ownerDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid Payload"})
	}

	owner, err := r.registerOwnerUseCase.Execute(use_cases.RegisterOwnerUseCaseRequest{
		Name:     ownerDto.Name,
		Document: ownerDto.Document,
		Birthday: ownerDto.Birthday,
		Email:    ownerDto.Email,
	})

	if err != nil {
		return HandleError(c, err)
	}

	return c.JSON(fiber.Map{
		"id":       owner.Id,
		"name":     owner.Name,
		"email":    owner.Email,
		"document": owner.Document,
		"birthday": owner.Birthday.Format("2006-10-20"),
	})
}
