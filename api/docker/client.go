package docker

import (
	"context"
	"github.com/docker/docker/client"
	"log"
)

var Client *client.Client

func Init() {
	var err error
	Client, err = client.NewClientWithOpts()
	if err != nil {
		log.Fatalln("Failed to connect to the docker daemon : \n", err)
	}

	//test if the docker daemon is running
	_, err = Client.Info(context.Background())
	if err != nil {
		log.Fatalln("Failed to connect to the docker daemon : \n", err)
	}
}
