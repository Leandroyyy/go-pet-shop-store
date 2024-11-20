package routes_owner

import (
	"github.com/gofiber/fiber/v2"
	input_fiber "github.com/leandroyyy/poc-golang/src/adapters/input/fiber"
	output_database "github.com/leandroyyy/poc-golang/src/adapters/output/database/in_memory"
	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/application/use_cases"
)

var inMemoryOwnerRepository = output_database.InMemoryOwnerRepository{}
var registerOwnerUseCase = use_cases.NewRegisterOwnerUseCase(inMemoryOwnerRepository)
var registerOwnerController = input_fiber.NewRegisterOwnerController(registerOwnerUseCase)

func registerOwnerRoute(router fiber.Router) {

	router.Post("/", func(c *fiber.Ctx) error {
		return registerOwnerController.Execute(c)
	})

}
