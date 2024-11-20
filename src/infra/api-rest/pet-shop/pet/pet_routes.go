package routes_pet

import (
	"github.com/gofiber/fiber/v2"
)

func PetRoutes(router fiber.Router) {

	petRoutes := router.Group("/pets")

	registerPetRoute(petRoutes)

}
