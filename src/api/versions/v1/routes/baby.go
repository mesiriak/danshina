package routes

import (
	"github.com/gofiber/fiber/v2"
	"main/src/api/versions/v1/handlers"
)

func RegisterBabyRoutes(app *fiber.App) {
	// FIXME: change groups prefix registration
	babyGroup := app.Group("api/v1/baby")

	babyGroup.Get("/", handlers.GetAllBabies)
	babyGroup.Get("/random/", handlers.GetRandomBabies)
	babyGroup.Post("/", handlers.CreateBaby)
	babyGroup.Patch("/:uuid/", handlers.UpdateBabyCounter)
	// babyGroup.Delete("/:uuid", handlers.DeleteBaby)
}
