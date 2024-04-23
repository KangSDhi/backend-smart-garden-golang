package routes

import (
	"backend-smart-garden-golang/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupGardenRouter(router fiber.Router) {

	gardenRoutes := router.Group("/garden")

	gardenRoutes.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    fiber.StatusOK,
			"message": "pong",
		})
	})

	gardenRoutes.Post("/create", controllers.CreateDataGarden)
	gardenRoutes.Get("/last", controllers.GetLastDataGarden)
}
