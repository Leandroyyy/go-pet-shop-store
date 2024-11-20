package routes_owner

import (
	"github.com/gofiber/fiber/v2"
)

func OwnerRoutes(router fiber.Router) {

	ownerRoutes := router.Group("/owners")

	registerOwnerRoute(ownerRoutes)

}
