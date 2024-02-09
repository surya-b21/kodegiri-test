package main

import (
	"kodegiri/app/router"
	"kodegiri/app/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

// @title Kodegiri App
// @version 1.0
// @description for test purpose
// @host localhost:8000
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	services.InitDB()

	app := fiber.New()
	router.Handle(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
