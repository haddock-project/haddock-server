package routes

import (
	"github.com/Kalitsune/Kontainerized/api/routes/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Init(app *fiber.App) {
	//home page
	//home := app.Group("/home")
	//home.Get("/")

	//ws routing
	ws := app.Group("/ws")
	ws.Get("/", websocket.New(wsHandler))

	//api.go routing
	api.Init(app)
}
