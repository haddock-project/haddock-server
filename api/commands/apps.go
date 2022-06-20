package commands

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Kalitsune/Haddock/api/database"
	"github.com/Kalitsune/Haddock/api/docker"
	"github.com/Kalitsune/Haddock/utils"
	"github.com/docker/docker/api/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
)

//GetApp returns a list of owned apps
func GetApp(ctx *fiber.Ctx) error {
	var (
		arg = ctx.Query("app")
	)

	if arg == "" {
		//get all apps
		//TODO: get user's apps

		//send an empty json
		return ctx.JSON(fiber.Map{})
	} else {
		//if there is an argument parse the uuid
		id, err := uuid.Parse(arg)
		if err != nil {
			return fiber.ErrBadRequest
		}

		//get the app associated to this uuid
		app := database.App{UUID: id}
		err = app.Get()
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fiber.ErrNotFound
			}
			return fiber.ErrInternalServerError
		}

		return ctx.JSON(app)
	}
}

//PostApp download a new image
func PostApp(ctx *fiber.Ctx) error {
	var app database.App

	err := ctx.BodyParser(&app)
	if err != nil {
		return fiber.ErrBadRequest
	}

	if app.Name == "" {
		/*
			There is a missing argument
		*/
		return fiber.ErrBadRequest
	}

	app.UUID = uuid.New()

	/*
		Check if the image is valid
	*/
	search, err := docker.SearchImage(app.Name)
	if err != nil {
		//handle timeout error
		if errors.Is(err, context.DeadlineExceeded) {
			return fiber.ErrRequestTimeout
		}
		return fiber.ErrInternalServerError
	}

	// no image has been found
	if len(search) == 0 {
		return fiber.ErrBadRequest
	}

	/*
		Download the image and handle the decoding/event delivery
	*/
	go func() {
		//pull the image
		log.Println("Pulling a new image: ", app.Name)
		err := docker.PullImage(app.Name)
		utils.HandleError("[ERROR] unable to pull the new image: "+app.Name, err)

		//add the app into the db
		log.Println("Registering a new app in the db: ", app.Name)
		if err = app.Set(); err != nil {
			utils.HandleError("[ERROR] unable to pull the new image: "+app.Name+". Reverting...", err)
			docker.Client.ImageRemove(context.Background(), app.Name, types.ImageRemoveOptions{Force: true})
		}

		log.Println("New app successfully set up: ", app.Name)
	}()

	//tell that the server will process the command (the download time may raise a timed out error)
	return ctx.JSON(fiber.Map{
		"status": "ok",
	})
}

func PatchApp(ctx *fiber.Ctx) error {
	var newApp database.App

	err := ctx.BodyParser(&newApp)
	if err != nil {
		return fiber.ErrBadRequest
	}

	if newApp.UUID == uuid.Nil {
		/*
			There is a missing argument
		*/
		return fiber.ErrBadRequest
	}

	//get the old app to compare
	oldApp := database.App{UUID: newApp.UUID}
	err = oldApp.Get()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fiber.ErrNotFound
		}
		return fiber.ErrInternalServerError
	}

	log.Println("App configuration is being updated: ", oldApp.Name)

	//Check if the image has changed
	if newApp.Name != oldApp.Name || newApp.RepoUrl != oldApp.RepoUrl {
		log.Println("Pulling a new image: ", newApp.Name)
		if err := docker.PullImage(newApp.Name); err != nil {
			utils.HandleError("[ERROR] unable to pull the new image: "+newApp.Name, err)
			return fiber.ErrInternalServerError
		}
	}

	//Update the db
	if err = newApp.Set(); err != nil {
		return fiber.ErrInternalServerError
	}

	log.Println("App configuration done: ", oldApp.Name)

	return nil
}

//DeleteApp removes an image from the docker daemon
func DeleteApp(ctx *fiber.Ctx) error {
	/*
		Remove the image
	*/

	//parse the uuid
	var img = ctx.Query("app")
	id, err := uuid.Parse(img)
	if err != nil {
		return fiber.ErrBadRequest
	}

	//get the app associated to this uuid
	app := database.App{UUID: id}
	err = app.Get()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fiber.ErrNotFound
		}
		return fiber.ErrInternalServerError
	}

	/*
		Remove the image
	*/
	log.Println("Removing the app: ", app.Name)
	_, err = docker.Client.ImageRemove(context.Background(), app.Name, types.ImageRemoveOptions{Force: true})
	if err != nil {
		utils.HandleError("[ERROR] unable to remove the image from docker: "+app.Name, err)
		return fiber.ErrInternalServerError
	}

	//remove the app from the db
	if err = app.Delete(); err != nil {
		utils.HandleError("[ERROR] unable to remove the image from the db: "+app.Name, err)
		return fiber.ErrInternalServerError
	}

	log.Println("App removed: ", app.Name)

	return ctx.JSON(fiber.Map{
		"message": "Image removed",
	})
}
