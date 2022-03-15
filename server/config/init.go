package config

import (
	"github.com/magiconair/properties"
	"log"
	"os"
)

var Server *properties.Properties

func Init() {
	// If the file doesn't exist
	if _, err := os.Stat("data/server.properties"); os.IsNotExist(err) {
		//create it
		file, err := os.Create("data/server.properties")
		if err != nil {
			log.Fatalln("Failed to create a server.properties file: ", err)
			return
		}

		file.Chmod(0666)
		log.Println("Successfully created a server.properties file")
	}

	// init from a file
	Server = properties.MustLoadFile("data/server.properties", properties.UTF8)
}
