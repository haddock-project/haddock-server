package routes

import (
	"encoding/json"
	"github.com/Kalitsune/Kontainerized/api/ws"
	"github.com/gofiber/websocket/v2"
)

func wsHandler(c *websocket.Conn) {
	var (
		request struct {
			Name string          `json:"name"`
			Args json.RawMessage `json:"args"`
		}
	)

	//save the connection to be able to send events
	ws.Register(c)

	//main listening loop
	for {
		err := c.ReadJSON(&request)
		if err != nil {
			return
		}

		/*
			Look for a matching command and gives it a client reference and the request arguments
		*/
		switch request.Name {
		default:
			/*
				No matching event has been found
			*/
			err = errors.New("unknown event")
		}

		if err != nil {
			//inform the client that an error occurred
			events.InternalError(&client, fmt.Sprintf("An error occured while handling your event: %s", err), events.InternalErrorArgs{})
		}
	}
}