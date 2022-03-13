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
	containers := api.Group("/container")

	/*
		Containers
	*/
	containers.Get("/", commands.GetContainers)
}
