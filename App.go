package main

import (
	"backend-smart-garden-golang/config"
	"backend-smart-garden-golang/entity"
	"backend-smart-garden-golang/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

func loadDatabase() {
	config.InitDB()
	err := config.DB.AutoMigrate(
		&entity.Garden{},
		&entity.SoilMoitureSensor{})
	if err != nil {
		log.Fatal("Error Migrate Garden")
	}
}

func serveApplication() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins: "*",
	}))

	apiRoutes := app.Group("/api")

	apiRoutes.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    fiber.StatusOK,
			"message": "pong",
		})
	})

	routes.SetupGardenRouter(apiRoutes)
	routes.SetupSoilMoistureSensorRouter(apiRoutes)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Server Error!")
	}
}
