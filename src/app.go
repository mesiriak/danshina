package src

import (
	"github.com/gofiber/fiber/v2"
	"main/cfg"
	"main/pkg/mongo"
	"main/src/api/versions/v1/routes"
)

func RegisterRoutes(app *fiber.App) {
	routes.RegisterBabyRoutes(app)
}

func CreateApp(config *cfg.Config) *fiber.App {
	app := fiber.New()

	mongo.InitDB(config.MongoUrl, config.DBName)
	RegisterRoutes(app)

	return app
}
