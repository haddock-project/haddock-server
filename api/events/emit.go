package events

import (
	"github.com/Kalitsune/Haddock/server/ws"
)

func (event *Event) Send(target *ws.Client) error {
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
	return conn.WriteJSON(event)
}
