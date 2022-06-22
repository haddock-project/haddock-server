package routes

import (
	"github.com/Kalitsune/Haddock/api/commands"
	"github.com/Kalitsune/Haddock/server/tokens"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Init(app *fiber.App) {
	/*
		Define the routes
	*/
	api := app.Group("/api")
	api.Use(tokens.Middleware)

	ws := api.Group("/ws")
	applications := api.Group("/app")
	users := api.Group("/user")

	/*
		Define the websocket routes
	*/
	ws.Use("/", wsAuthenticator)
	ws.Get("/", websocket.New(wsHandler))

	/*
		Define the App routes
	*/
	applications.Get("/", commands.GetApp)
	applications.Post("/", commands.PostApp)
	applications.Patch("/", commands.PatchApp)
	applications.Delete("/", commands.DeleteApp)

	/*
		Define the User routes
	*/
	users.Get("/", commands.GetUser)
	users.Post("/auth", commands.AuthUser)
	users.Post("/", commands.PostUser)
	users.Patch("/", commands.PatchUser)
	users.Delete("/", commands.DeleteUser)
}
