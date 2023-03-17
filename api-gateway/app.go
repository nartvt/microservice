package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"api-gateway/app/infra/db"
	"api-gateway/app/router"
)

func main() {
	setupInfra()
	defer closeInfra()
	app := fiber.New()
	initRouter(app)
}

func setupInfra() {
	setupDatabase()
}

func closeInfra() {
	db.ClosePostgres()
}

func setupDatabase() {
	db.InitPostgres()
}
func initRouter(app *fiber.App) {
	router.SetupRoutes(app)
	if err := app.Listen("0.0.0.0:3000"); err != nil {
		log.Println("Application startup interrupt")
		panic(err.Error())
	}

}
