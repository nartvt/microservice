package router

import "github.com/gofiber/fiber/v2"

func GET(app fiber.Router, relativePath string, f fiber.Handler) {
	route(app, "GET", relativePath, f)
}

func POST(app fiber.Router, relativePath string, f fiber.Handler) {
	route(app, "POST", relativePath, f)
}
func PUT(app fiber.Router, relativePath string, f fiber.Handler) {
	route(app, "PUT", relativePath, f)
}

func DELETE(app fiber.Router, relativePath string, f fiber.Handler) {
	route(app, "DELETE", relativePath, f)
}
func route(app fiber.Router, method string, relativePath string, f fiber.Handler) {
	switch method {
	case "POST":
		app.Post(relativePath, f)
	case "GET":
		app.Get(relativePath, f)
	case "PUT":
		app.Put(relativePath, f)
	case "DELETE":
		app.Delete(relativePath, f)
	}
}
