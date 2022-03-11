package routes

import (
	"encoding/json"
	"github.com/gofiber/websocket/v2"
)

func wsHandler(c *websocket.Conn) {
	var (
		request struct {
			Name string          `json:"name"`
			Args json.RawMessage `json:"args"`
		}
	)

	//main listening loop
	for {
		err := c.ReadJSON(&request)
		if err != nil {
			return 
		}
	}
}
