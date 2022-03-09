package routes

import "github.com/gofiber/fiber/v2"

func Init(app *fiber.App) {
	//home page
	home := app.Group("/home")
	home.Get("/")

	//ws routing
	ws := app.Group("/ws")
	ws.Use("/", wsMiddleware)
	ws.Get("/", wsMiddleware)
}
