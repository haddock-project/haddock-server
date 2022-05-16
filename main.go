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

	//get config
	config.Init()
	log.Println("Config loaded!")

	//init docker api.go
	docker.Init()
	log.Println("Docker connexion established!")

	//init database
	database.Init()
	log.Println("Sqlite ready!")

	// ensure server's private key is generated
	if key := config.GetPrivateKey(); key == "" {
		log.Println("Generating server's private key...")
		config.GeneratePrivateKey()
		log.Println("Server's private key generated!")
	}

	//create webserver using https://github.com/gofiber/fiber
	app := fiber.New()

	// init webserver routes
	routes.Init(app)
	log.Println("Fiber ready!")

	//listen on port 8080
	log.Fatal(app.Listen(":8080"))
}
