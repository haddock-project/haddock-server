package api

import "github.com/gofiber/fiber/v2"

func Init(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/")
}
