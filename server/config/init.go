package config

import (
	"github.com/magiconair/properties"
	"os"
)

var props *properties.Properties

func Init() {
	// If the file doesn't exist
	if _, err := os.Stat("data/server.properties"); os.IsNotExist(err) {
		// Create the file
		out, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer out.Close()
	}

	// init from a file
	props = properties.MustLoadFile("data/server.properties", properties.UTF8)
}
