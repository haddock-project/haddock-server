package storage

import (
	"github.com/Kalitsune/Haddock/api/database"
	"log"
	"os"
)

func Init() {
	//check for the Icons' directory
	//if not exists, create it
	if _, err := os.Stat("data/icons"); os.IsNotExist(err) {
		err := os.Mkdir("data/icons", 0744)
		if err != nil {
			log.Fatalln("Failed to create icons' folder: ", err)
		}
	}
}

//AddApp generate the app's files and edit the icon
func AddApp(app database.App) (database.App, error) {
	//download app icon

	return app, nil
}
