package rest

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// Error sends a json containing the error and the reason
func Error(ctx *fiber.Ctx, reason string) error {
	body := fmt.Sprintf(`{"name":"ERROR","args":{"reason":"%s"}}`, reason)
	return ctx.Send([]byte(body))
}
