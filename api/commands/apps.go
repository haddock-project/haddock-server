package commands

import (
	"context"
	"errors"
	"github.com/Kalitsune/Haddock/api/docker"
	"github.com/docker/docker/api/types"
	"github.com/gofiber/fiber/v2"
	"log"
)

type image struct {
	Name []string `json:"name"`
	Id   string   `json:"id"`
}

//GetApp returns processed images
func GetApp(ctx *fiber.Ctx) error {
	var (
		q      = ctx.Query("app")
		err    error
		images []types.ImageSummary
	)

	images, err = docker.GetContainers(q)
	if err != nil {
		return fiber.ErrInternalServerError
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
		Check if the image is valid
	*/
	search, err := docker.SearchImage(img)
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
		log.Println("Pulling a new image: ", img)
		docker.PullImage(img)
	}()

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
