package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	routes_owner "github.com/leandroyyy/poc-golang/src/infra/api-rest/pet-shop/owner"
	routes_pet "github.com/leandroyyy/poc-golang/src/infra/api-rest/pet-shop/pet"
)

func routes(app *fiber.App) {
	routes_owner.OwnerRoutes(app)
	routes_pet.PetRoutes(app)
}

func main() {
	app := fiber.New()

	routes(app)

	fmt.Println("Servidor iniciado em http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}

}
