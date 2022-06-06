package config

import (
	"github.com/magiconair/properties"
	"log"
	"os"
)

var props *properties.Properties

func Init() {
	// If the file doesn't exist
	if _, err := os.Stat("data/server.properties"); os.IsNotExist(err) {
		// Create the file
		out, err := os.Create("data/server.properties")
		if err != nil {
			log.Fatalln("Failed to create server.properties: ", err)
		}
		defer out.Close()
	}

	// init from a file
	props = properties.MustLoadFile("data/server.properties", properties.UTF8)
}
