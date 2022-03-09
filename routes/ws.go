package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func wsMiddleware(c *fiber.Ctx) error {
	// IsWebSocketUpgrade returns true if the client
	// requested upgrade to the WebSocket protocol.
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func wsHandler(c *websocket.Conn) {
	//main listening loop
	//for {
	//	c.ReadJSON()
	//}
}
