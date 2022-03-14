package commands

import (
	"context"
	"encoding/json"
	"github.com/Kalitsune/Haddock/api/docker"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/gofiber/fiber/v2"
	"io"
	"strings"
)

type image struct {
	Name []string `json:"name"`
	Id   string   `json:"id"`
}

func GetApp(ctx *fiber.Ctx) error {
	var (
		q      = ctx.Query("image")
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
			//TODO: include state of owned containers
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

func PostApp(ctx *fiber.Ctx) error {
	var image = ctx.Query("image")

	if image == "" {
		/*
			There is a missing argument
		*/
		return fiber.ErrBadRequest
	}

	//say that the server will process the command (the download time may raise a timed out error)
	ctx.SendString("OK")

	//pull the image
	events, err := docker.Client.ImagePull(context.Background(), image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	//start a json decoder
	d := json.NewDecoder(events)
	type Event struct {
		Status         string `json:"status"`
		Error          string `json:"error"`
		ProgressDetail struct {
			Current int `json:"current"`
			Total   int `json:"total"`
		} `json:"progressDetail"`
	}

	//read the flux
	var event *Event
	for {
		//decode / handle decoding error
		if err := d.Decode(&event); err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		//check if the download is done
		done := strings.HasPrefix(event.Status, "Image is up to date for ") || strings.HasPrefix(event.Status, "Downloaded newer image for ")

		//send an adapted event
		if event.Error != "" {
			//send an error event
		} else if event.Status == "Downloading" {
			// send a download event
		} else if event.Status == "Extracting" {
			//send an extract event
		} else if done {
			//send a download complete
		}
		//TODO send event
	}

	return nil
}
