package commands

import (
	"context"
	"github.com/Kalitsune/Haddock/api/docker"
	"github.com/docker/docker/api/types"
	"github.com/gofiber/fiber/v2"
)

type container struct {
	Name  []string `json:"name"`
	Id    string   `json:"id"`
	State string   `json:"state"`
}

func GetApp(ctx *fiber.Ctx) error {
	var (
		app        = ctx.Params("app")
		containers []types.Container
		err        error
	)

	if app != "" {
		containers, err = docker.Client.ContainerList(context.Background(), types.ContainerListOptions{})

	} else {
		/*
			Get containers
		*/
		containers, err = docker.Client.ContainerList(context.Background(), types.ContainerListOptions{})
		if err != nil {
			return err
		}

		if len(containers) > 0 {
			/*
				Create a simplified array of containers to be sent to the client
			*/
			var body []container
			for _, c := range containers {
				// for each iteration, append a new container in the array
				body = append(body, container{Name: c.Names, Id: c.ID, State: c.State})
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
	return nil
}

func PostApp(ctx *fiber.Ctx) error {
	//var c = ctx.Params("app")
	//
	//docker.Client.ContainerCreate()
	return nil
}
