package main

import (
	"kodegiri/app/services"
)

// @title Kodegiri App
// @version 1.0
// @description for test purpose
// @host localhost:8000
// @BasePath /api
func main() {
	services.InitDB()

	// log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
