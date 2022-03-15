package routes

import (
	"github.com/Kalitsune/Haddock/server/routes/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Init(app *fiber.App) {
	//ws routing
	ws := app.Group("/ws")
	ws.Use("/", wsAuthenticator)
	ws.Get("/", websocket.New(wsHandler))

	//api.go routing
	api.Init(app)
}
