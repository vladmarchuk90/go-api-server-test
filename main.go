package main

import (
	"github.com/gofiber/fiber"
	"github.com/vladmarchuk90/go-api-server-test/controller"
	"github.com/vladmarchuk90/go-api-server-test/model"
)

func setupRoutes(app *fiber.App) {

	app.Get("/", controller.Default)
	app.Get("/groups", controller.GetGroups)
	app.Get("/groups/:groupId", controller.GetGroup)
	app.Get("/indexes/:indexId", controller.GetIndex)
	app.Get("/blocks/:blockParameter", controller.GetBlock)
}

func main() {

	// Initialize model
	model.ConnectToEthereumSmartContract()
	defer model.CloseConnection()

	// Create and run web-server
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
}
