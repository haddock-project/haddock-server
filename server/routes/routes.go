package routes

import (
	"github.com/Kalitsune/Haddock/server/routes/api"
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	//api.go routing
	api.Init(app)
}
