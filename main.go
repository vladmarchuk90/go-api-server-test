package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/vladmarchuk90/go-api-server-test/docs"

	"github.com/gofiber/fiber/v2"
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

func setupSwaggerRoutes(app *fiber.App) {

	app.Get("/swagger/*", swagger.Handler)
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:3000/swagger/oauth2-redirect.html",
	}))
}

// @title Ethereum smart contract API
// @version 1.0
// @description This is a server to get information from smart contract 0x4f7f1380239450AAD5af611DB3c3c1bb51049c29 and general block information.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:3000
// @BasePath /

// @securityDefinitions.basic not used in current API's version

func main() {

	// Initialize model
	model.ConnectToEthereumSmartContract()
	defer model.CloseConnection()

	// Create and run web-server
	app := fiber.New()
	setupRoutes(app)

	// adding swagger documentation
	setupSwaggerRoutes(app)

	app.Listen("localhost:3000")

}
