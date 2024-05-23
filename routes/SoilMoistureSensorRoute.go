package routes

import (
	"backend-smart-garden-golang/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupSoilMoistureSensorRouter(router fiber.Router) {

	soilMoistureSensorRoutes := router.Group("/soilmoisturesensor")

	soilMoistureSensorRoutes.Post("/create", controllers.CreateDataSoilMoistureSensor)
}
