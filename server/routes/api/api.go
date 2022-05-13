package api

import (
	"github.com/Kalitsune/Haddock/api/commands"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Init(app *fiber.App) {
	/*
		Define the routes
	*/
	api := app.Group("/api")

	ws := api.Group("/ws")
	container := api.Group("/app")

	/*
		Define the websocket routes
	*/
	ws.Use("/", wsAuthenticator)
	ws.Get("/", websocket.New(wsHandler))

	/*
		Define the App routes
	*/
	container.Get("/", commands.GetApp)
	container.Post("/", commands.PostApp)
	container.Delete("/", commands.DeleteApp) //TODO: test
}
