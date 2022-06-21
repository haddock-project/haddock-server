package tokens

import (
	"github.com/Kalitsune/Haddock/api/database"
	"github.com/Kalitsune/Haddock/server/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// Middleware is used by go-fiber to create a User object in the Locals.
// Note: If the user is not authenticated, the User object is empty
func Middleware(ctx *fiber.Ctx) error {

	//load jwt from the cookie
	token := ctx.Cookies("token")
	if token == "" {
		//return a blank user because it is not authenticated
		ctx.Locals("user", database.User{})
	} else {
		// Initialize a new instance of `Claims`
		claims := jwt.MapClaims{}

		//read more here: https://pkg.go.dev/github.com/golang-jwt/jwt#Parser.ParseWithClaims
		//We do not use the Keyfunc argument here, so we just return the private key
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return config.GetPrivateKey(), nil
		})
		if err != nil || !tkn.Valid {
			//the provided token is invalid/has an invalid signature, clear it and add a blank user
			ctx.ClearCookie("token")
			ctx.Locals("user", database.User{})
		}
	}

	return ctx.Next()
}
