package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func Init() {
	var err error
	db, err = sql.Open("sqlite3", "data/haddock.db")
	if err != nil {
		log.Fatalln("Unable to open a new database in the persistent data folder : \n", err)
	}

	// Create the tables if they don't exist
	err = initTables()
	if err != nil {
		log.Fatalln("Unable to init the tables: \n", err)
	}
}

func initTables() error {
	var err error

	//init the user table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `users` (`user_name` varchar(255) NOT NULL UNIQUE PRIMARY KEY,`user_password` varchar(255) NOT NULL, `user_icon` BLOB, `user_permissions` INT DEFAULT 0, `password_reset` BOOLEAN NOT NULL CHECK (password_reset IN (0, 1)));")
	if err != nil {
		return err
	}

	//init the app table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `apps` ( `app_uuid` blob(16) NOT NULL UNIQUE PRIMARY KEY, `app_name` varchar(255) NOT NULL, `app_icon` blob NOT NULL, `app_description` varchar(500) NOT NULL,`app_version` varchar(16) NOT NULL, `app_url` varchar(255) NOT NULL, `repo_url` varchar(255) NOT NULL, `repo_name` varchar(255) NOT NULL);")
	if err != nil {
		return err
	}

	//init the apps_user table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `apps_users` ( `app_uuid` blob(16) NOT NULL, `user_name` varchar(255) NOT NULL, FOREIGN KEY (`app_uuid`) REFERENCES `apps` (`app_uuid`), FOREIGN KEY (`user_name`) REFERENCES `users` (`user_name`) );")
	if err != nil {
		return err
	}

	return nil
}
