package docker

import "github.com/docker/docker/client"

var Client *client.Client

func Init() {
	var err error
	Client, err = client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}
}
