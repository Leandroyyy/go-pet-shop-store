package routes_pet

import (
	"github.com/gofiber/fiber/v2"
	input_fiber "github.com/leandroyyy/poc-golang/src/adapters/input/fiber"
	output_database "github.com/leandroyyy/poc-golang/src/adapters/output/database/in_memory"
	"github.com/leandroyyy/poc-golang/src/domain/pet_shop/application/use_cases"
)

var inMemoryPetRepository = output_database.InMemoryPetRepository{}
var inMemoryOwnerRepository = output_database.InMemoryOwnerRepository{}
var registerPetUseCase = use_cases.NewRegisterPetUseCase(inMemoryOwnerRepository, inMemoryPetRepository)
var registerPetController = input_fiber.NewRegisterPetController(registerPetUseCase)

func registerPetRoute(router fiber.Router) {

	router.Post("/", func(c *fiber.Ctx) error {
		return registerPetController.Execute(c)
	})

}
