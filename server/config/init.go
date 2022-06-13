package config

import (
	"github.com/magiconair/properties"
	"os"
)

var props *properties.Properties

func Init() {
	var err error

	// init from a file
	props, err = properties.LoadFile("data/server.properties", properties.UTF8)
	if err != nil {
		props = properties.NewProperties()
	}
}

func Save(props *properties.Properties) error {
	//load file
	file, err := os.OpenFile("data/server.properties", os.O_CREATE+os.O_RDWR, 0660)
	if err != nil {
		return err
	}

	//write new props
	_, err = props.Write(file, properties.UTF8)
	return err
}
