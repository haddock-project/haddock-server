package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Kalitsune/Haddock/api/events"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/registry"
	"io"
	"strings"
	"time"
)

//GetContainers takes a query string as input and returns a list of containers filtered by the query
func GetContainers(query string) ([]types.ImageSummary, error) {
	var opt types.ImageListOptions

	/*
		Get containers
	*/

	//if a repo name has been provided, filter it
	if query != "" {
		//filter for a specific repository
		filter := filters.NewArgs(filters.Arg("reference", query+"*"))
		opt = types.ImageListOptions{Filters: filter}
	} else {
		// returns a blank filter
		opt = types.ImageListOptions{}
	}

	// ask the docker daemon
	return Client.ImageList(context.Background(), opt)
}

func SearchImage(query string) ([]registry.SearchResult, error) {
	//create the timeout and a cancel func
	timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//search the image on docker hub
	return Client.ImageSearch(timeout, query, types.ImageSearchOptions{Limit: 1})
}

func PullImage(image string) {
	stream, err := Client.ImagePull(context.Background(), image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	//decode and send events
	decodePullStream(stream, image)
}

func decodePullStream(stream io.ReadCloser, image string) {
	//start a json decoder
	d := json.NewDecoder(stream)
	type Event struct {
		Status         string `json:"status"`
		Error          string `json:"error"`
		ProgressDetail struct {
			Current int `json:"current"`
			Total   int `json:"total"`
		} `json:"progressDetail"`
	}

	//read the flux
	var pullEvent *Event
	for {
		//decode / handle decoding error
		if err := d.Decode(&pullEvent); err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		/*
			Send an adapted event
		*/
		var (
			name string
			args = "{}"
		)
		//check if the download is done
		done := strings.HasPrefix(pullEvent.Status, "Status: Image is up to date for ") || strings.HasPrefix(pullEvent.Status, "Status: Downloaded newer image for ")

		//look for the event to send
		if pullEvent.Error != "" {
			/*
				Send an error event
			*/
			name = "APP_DOWNLOAD_ERROR"
			args = fmt.Sprintf(`{"name":"%s","error":"%s"}`, image, pullEvent.Error)

		} else if done {
			/*
				Send a download complete event
			*/
			name = "APP_DOWNLOAD_COMPLETE"
			args = fmt.Sprintf(`{"name":"%s"}`, image)

		} else {
			//these events are handling an installation progress
			//check if the total is 0 (because the image isn't being downloaded yet)
			if pullEvent.ProgressDetail.Total == 0 {
				continue
			}

			progress, _ := json.Marshal(&pullEvent.ProgressDetail)

			if pullEvent.Status == "Downloading" {
				/*
					Send a download event
				*/
				name = "APP_DOWNLOAD_PROGRESS"
				args = fmt.Sprintf(`{"name":"%s","progress": %s}`, image, progress)
			} else if pullEvent.Status == "Extracting" {
				/*
					Send an extract event
				*/
				name = "APP_EXTRACT_PROGRESS"
				args = fmt.Sprintf(`{"name":"%s","progress": %s}`, image, progress)

			} else {
				// this event is unhandled so skip
				continue
			}
		}

		event := events.Event{Name: name, Args: []byte(args)}
		event.Send("") //TODO: (accounts) find the targets
	}
}
