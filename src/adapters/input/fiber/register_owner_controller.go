package input_fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/application/use_cases"
)

type dto struct {
	name     string `json:"name"`
	document string `json:"document"`
	birthday string `json:"birthday"`
	email    string `json:"email"`
}

type RegisterOwnerController struct {
	RegisterOwnerUseCase use_cases.RegisterOwnerUseCase
}

func (r RegisterOwnerController) Execute(c fiber.Ctx) error {

	var ownerDto dto

	if err := c.BodyParser(&ownerDto); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}

	owner, err := r.RegisterOwnerUseCase.Execute(use_cases.RegisterOwnerUseCaseRequest{
		Name:     ownerDto.name,
		Document: ownerDto.document,
		Birthday: ownerDto.birthday,
		Email:    ownerDto.email,
	})

	if err != nil {
		return HandleError(&c, err)
	}

	return c.JSON(fiber.Map{
		"id":       owner.Id,
		"name":     owner.Name,
		"email":    owner.Email,
		"document": owner.Document,
		"birthday": owner.Birthday,
	})
}
