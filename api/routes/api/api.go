package api

import (
	"github.com/Kalitsune/Haddock/api/commands"
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	/*
		Define the routes
	*/
	api := app.Group("/api")
	container := api.Group("/container")

	/*
		Containers
	*/
	container.Get("/:app?", commands.GetApp)
	container.Post("/:app", commands.PostApp)
}
