package database

type App struct {
	Id          uint16 `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	AppUrl      string `json:"app_url"`
	RepoUrl     string
	RepoName    string
}

//AddApp bundle a container with infos about the app (it also generates a database id)
func AddApp(app App) (App, error) {

	// insert the app into the database
	_, err := db.Exec("INSERT INTO apps (app_name, app_icon, app_description, app_url, repo_url, repo_name) VALUES (?, ?, ?, ?, ?, ?)", app.Name, app.Icon, app.Description, app.AppUrl, app.RepoUrl, app.RepoName)
	if err != nil {
		return app, err
	}

	//get the app id
	var id uint16
	err = db.QueryRow("SELECT app_id FROM apps WHERE app_name = ?", app.Name).Scan(&id)
	if err != nil {
		return app, err
	}

	//set the app id
	app.Id = id

	return app, nil
}

// GetAppFromID takes an ID and returns the app
func GetAppFromID(id uint16) (App, error) {
	var app App

	err := db.QueryRow("SELECT app_name, app_icon, app_description, app_url, repo_url, repo_name FROM apps WHERE app_id = ?", id).Scan(&app.Name, &app.Icon, &app.Description, &app.AppUrl, &app.RepoUrl, &app.RepoName)
	if err != nil {
		return app, err
	}

	app.Id = id

	return app, nil
}

//DeleteApp delete an app from the db as well as all its files
func DeleteApp(id uint16) error {
	_, err := db.Exec("DELETE FROM apps WHERE app_id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
