package main

import (
	"github.com/Kalitsune/Kontainerized/web/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	routes.Init(app)

	log.Fatal(app.Listen(":8080"))
}
