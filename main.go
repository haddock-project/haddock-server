package main

import (
	"github.com/Kalitsune/Haddock/api/database"
	"github.com/Kalitsune/Haddock/api/docker"
	"github.com/Kalitsune/Haddock/server/config"
	"github.com/Kalitsune/Haddock/server/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	log.Println("Starting Haddock server!")

	//init config
	config.Init()
	log.Println("Config:  READY")

	//init docker routes.go
	docker.Init()
	log.Println("Docker:  READY ")

	//init database
	database.Init()
	log.Println("Sqlite:  READY ")

	// ensure core properties are valid
	config.GetPrivateKey()
	config.GetTokenExpiration()
	config.GetRememberMeTokenExpiration()

	//create webserver using https://github.com/gofiber/fiber
	app := fiber.New()

	// init webserver routes
	routes.Init(app)
	log.Println("Fiber:   READY")

	//listen on port 8080
	log.Fatal(app.Listen(config.GetHost()))
}
