package commands

import (
	"context"
	"github.com/Kalitsune/Kontainerized/api/docker"
	"github.com/docker/docker/api/types"
	"github.com/gofiber/fiber/v2"
)

type container struct {
	Name  []string `json:"name"`
	Id    string   `json:"id"`
	State string   `json:"state"`
}

func GetContainers(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json") // "application/json"

	/*
		Get containers
	*/
	containers, err := docker.Client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return err
	}

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
}
