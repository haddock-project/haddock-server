package database

import (
	"encoding/base64"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
)

type App struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"` //base64 encoded image
	Description string    `json:"description"`
	AppUrl      string    `json:"app_url"`
	Version     string    `json:"version"`
	RepoUrl     string    `json:"repo_url"`
	RepoName    string    `json:"repo_name"`
}

//Set update an existing app or create a new one
func (app *App) Set() error {
	//check if the icon field is an image
	app.CheckIcon()

	// insert the app into the database / update
	_, err := db.Exec(
		"INSERT INTO apps (app_uuid, app_name, app_icon, app_description, app_version, app_url, repo_url, repo_name) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"+
			"ON CONFLICT(app_uuid) DO UPDATE SET (app_name, app_icon, app_description, app_version, app_url, repo_url, repo_name) = ($2, $3, $4, $5, $6, $7, $8)",
		app.UUID, app.Name, app.Icon, app.Description, app.Version, app.AppUrl, app.RepoUrl, app.RepoName)

	return err
}

// Get takes a UUID and returns the app
func (app *App) Get() error {
	err := db.QueryRow("SELECT app_name, app_icon, app_description, app_url, repo_url, repo_name FROM apps WHERE app_uuid = ?", app.UUID).Scan(&app.Name, &app.Icon, &app.Description, &app.AppUrl, &app.RepoUrl, &app.RepoName)
	return err
}

//Delete an app from the db
func (app *App) Delete() error {
	_, err := db.Exec("DELETE FROM apps WHERE app_uuid = ?", app.UUID)
	return err
}

//CheckIcon clear the icon if it's not an image
func (app *App) CheckIcon() {
	//Get the image from the byte stream
	decodeString, err := base64.StdEncoding.DecodeString(app.Icon)
	if err != nil {
		log.Println("Warning: invalid icon for \"", app.Name, "\". Icon cleared.")
		app.Icon = ""
	}

	//detect content type, if the mime doesn't start with "image/" then the image is considered as invalid
	if mime := http.DetectContentType(decodeString); !strings.HasPrefix(mime, "image/") {
		log.Println("Warning: invalid icon for \"", app.Name, "\". Icon cleared.")
		app.Icon = ""
	}
}
