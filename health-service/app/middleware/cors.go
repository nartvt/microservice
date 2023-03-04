package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsFilter() fiber.Handler {
	return cors.New(*corsConfig())
}

func corsConfig() *cors.Config {
	corsConfig := cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}
	return &corsConfig
}
