package main

import (
	"github.com/Kalitsune/Kontainerized/api/docker"
	"github.com/Kalitsune/Kontainerized/web/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	//init docker api.go
	docker.Init()

	//create webserver using https://github.com/gofiber/fiber
	app := fiber.New()

	// init webserver routes
	routes.Init(app)

	//listen on port 8080
	log.Fatal(app.Listen(":8080"))
}
