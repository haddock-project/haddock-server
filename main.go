package main

import (
	"github.com/Kalitsune/Haddock/api/database"
	"github.com/Kalitsune/Haddock/api/docker"
	"github.com/Kalitsune/Haddock/api/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	log.Println("Starting Haddock server!")

	//init docker api.go
	docker.Init()
	log.Println("Docker connexion established!")

	//init database
	database.Init()
	log.Println("Sqlite ready!")

	//create webserver using https://github.com/gofiber/fiber
	app := fiber.New()

	// init webserver routes
	routes.Init(app)
	log.Println("Fiber ready!")

	//listen on port 8080
	log.Fatal(app.Listen(":8080"))
}
