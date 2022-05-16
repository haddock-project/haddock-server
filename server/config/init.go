package config

import (
	"github.com/magiconair/properties"
	"log"
	"os"
)

var props *properties.Properties

func Init() {
	// If the file doesn't exist
	file, err := os.OpenFile("data/server.properties", os.O_CREATE, 660)
	if err != nil {
		log.Fatalln("Failed to open server.properties: ", err)
	}
	file.Close()

	// init from a file
	props = properties.MustLoadFile("data/server.properties", properties.UTF8)
}
