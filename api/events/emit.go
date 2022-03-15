package events

import (
	"errors"
	"github.com/Kalitsune/Haddock/server/ws"
)

func (event *Event) Send(token string) error {

	/*
		Find the targets
	*/
	var targets []*ws.Client
	if token == "" {
		//if no token was provided ping every connected clients
		targets = ws.Websocket.Clients
	} else {
		return errors.New("unsupported yet")
	}

	for _, target := range targets {
		/*
			Check connection
		*/
		conn := target.Conn
		if conn == nil {
			//client is disconnected but cache hasn't been updated
			target.Remove()
		}

		/*
			Send the event
		*/
		conn.WriteJSON(event)
	}

	return nil
}
