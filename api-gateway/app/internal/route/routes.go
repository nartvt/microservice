package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"api-gateway/app/middleware"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Use(logger.New())
	v1.Use(middleware.CorsFilter(), middleware.RateLimit())

	setupUserRoute(v1)
	setupNewsfeedRoute(v1.Group("/user/me", middleware.Auth.RequireLogin()))
	setupAboutRoute(v1)
}
