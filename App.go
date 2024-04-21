package main

import (
	"backend-smart-garden/config"
	"backend-smart-garden/entity"
	"backend-smart-garden/routes"
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
	err0 := config.DB.AutoMigrate(&entity.Garden{})
	if err0 != nil {
		log.Fatal("Error Migrate Garden")
	}
	//migration.SeedGarden()
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

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Server Error!")
	}
}
