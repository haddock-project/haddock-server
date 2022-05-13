package commands

import (
	"context"
	"errors"
	"github.com/Kalitsune/Haddock/api/docker"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

type image struct {
	Name []string `json:"name"`
	Id   string   `json:"id"`
}

//GetApp returns processed images
func GetApp(ctx *fiber.Ctx) error {
	var (
		q      = ctx.Query("app")
		images []types.ImageSummary
		err    error
		opt    types.ImageListOptions
	)

	/*
		Get containers
	*/

	//if a repo name has been provided, filter it
	if q != "" {
		//filter for a specific repository
		filter := filters.NewArgs(filters.Arg("reference", q+"*"))
		opt = types.ImageListOptions{Filters: filter}
	} else {
		// returns a blank filter
		opt = types.ImageListOptions{}
	}

	// ask the docker daemon
	images, err = docker.Client.ImageList(context.Background(), opt)
	if err != nil {
		return err
	}

	if len(images) > 0 {
		/*
			Create a simplified array of containers to be sent to the client
		*/
		var body []image
		for _, c := range images {
			// for each iteration, append a new image in the array
			body = append(body, image{Name: c.RepoTags, Id: c.ID})
		}

		/*
			Respond to the request
		*/
		return ctx.JSON(body)
	} else {
		/*
			There is no containers, send an empty list to the client
		*/
		ctx.Set("Content-Type", "application/json")
		return ctx.SendString("[]")
	}
}

//PostApp download a new image
func PostApp(ctx *fiber.Ctx) error {
	var img = ctx.Query("app")

	if img == "" {
		/*
			There is a missing argument
		*/
		return fiber.ErrBadRequest
	}

	/*
		Check if the image is valid with a 10s timeout
	*/
	//create the timeout and a cancel func
	timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//search the image on docker hub
	search, err := docker.Client.ImageSearch(timeout, img, types.ImageSearchOptions{Limit: 1})
	//handle timeout error
	if errors.Is(err, context.DeadlineExceeded) {
		cancel()
		return fiber.ErrRequestTimeout
	}

	// no image has been found
	if len(search) == 0 {
		cancel()
		return fiber.ErrBadRequest
	}

	/*
		Download the image and handle the decoding/event delivery
	*/
	go func(cancel context.CancelFunc) {
		log.Println("Pulling a new image: ", img)
		docker.PullImage(img)
		cancel()
	}(cancel)

	//say that the server will process the command (the download time may raise a timed out error)
	return ctx.JSON(fiber.Map{
		"status": "ok",
	})
}

//DeleteApp removes an image from the docker daemon
func DeleteApp(ctx *fiber.Ctx) error {
	var img = ctx.Query("app")

	if img == "" {
		/*
			There is a missing argument
		*/
		return fiber.ErrBadRequest
	}

	/*
		Remove the image
	*/
	_, err := docker.Client.ImageRemove(context.Background(), img, types.ImageRemoveOptions{Force: true})
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "Image removed",
	})
}
