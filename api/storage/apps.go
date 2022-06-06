package storage

import (
	"errors"
	"github.com/Kalitsune/Haddock/api/database"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func InitApps() {
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
func AddApp(app database.App) error {
	//get the icon
	stream, err := getIcon(app.Icon)
	if err != nil {
		return err
	}

	//read the stream
	bytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return err
	}
	stream.Close()

	mimeType, err := getIconType(bytes)
	if err != nil {
		return err
	}

	//process the mime type
	identity := strings.Split(mimeType, "/")
	fileType := identity[0]
	fileExtension := identity[1]

	//check if the mime type is an image
	if fileType != "image" {
		return errors.New("The icon is not an image")
	}

	//save the icon
	err = saveIcon(app, bytes, fileExtension)
	if err != nil {
		return err
	}

	return nil
}

//getIcon do a get request to download the icon
func getIcon(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	return resp.Body, err
}

func saveIcon(app database.App, bytes []byte, extension string) error {
	//Write the bytes to the file
	return ioutil.WriteFile("data/icons/"+app.Name+"."+extension, bytes, 0644)
}

//getIconType get the icon type from the mime
func getIconType(bytes []byte) (string, error) {
	// Get the mime type of the byte stream
	mimeType := http.DetectContentType(bytes)

	return mimeType, nil
}
