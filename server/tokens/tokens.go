package tokens

import (
	"github.com/Kalitsune/Haddock/api/database"
	"github.com/Kalitsune/Haddock/server/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

// Middleware is used by go-fiber to create a User object in the Locals.
// Note: If the user is not authenticated, the User object is empty
func Middleware(ctx *fiber.Ctx) error {

	//load jwt from the header
	token := ctx.GetReqHeaders()["Authorization"]
	if token == "" {
		//return a blank user because it is not authenticated
		ctx.Locals("user", &database.User{})
		return ctx.Next()

	} else {
		//the header should be `Bearer <token>` so we split on space
		token = strings.Split(token, " ")[1]

		// Initialize a new instance of `Claims`
		claims := &jwt.StandardClaims{}

		//read more here: https://pkg.go.dev/github.com/golang-jwt/jwt#Parser.ParseWithClaims
		//We do not use the Keyfunc argument here, so we just return the public key
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return config.GetPrivateKey().Public(), nil
		})

		//Handle validation errors
		v, _ := err.(*jwt.ValidationError)
		tokenNoExpireErr := claims.ExpiresAt < claims.IssuedAt && v.Errors == jwt.ValidationErrorExpired
		if err != nil && !tokenNoExpireErr {
			//the provided token is invalid, clear it and add a blank user
			ctx.ClearCookie("token")
			ctx.Locals("user", &database.User{})
			return ctx.Next()
		}

		//get the user from the claims
		user := &database.User{Name: claims.Audience}
		if err = user.Get(); err != nil {
			println(err.Error())
			//the user does not exist/an error has occurred, clear the cookie and return a blank user
			ctx.ClearCookie("token")
			ctx.Locals("user", &database.User{})

		} else {
			//set the user in the locals
			ctx.Locals("user", user)
		}

		return ctx.Next()
	}
}

func MakeToken(user *database.User, rememberMe bool) (string, error) {
	var (
		expiration int64
	)

	//check which duration should be applied
	if rememberMe {
		expiration = config.GetRememberMeTokenExpiration().Unix()
	} else {
		expiration = config.GetTokenExpiration().Unix()
	}

	//Set the token claims
	claims := jwt.StandardClaims{
		Audience:  user.Name,
		Issuer:    "github.com/haddock-project/haddock-server",
		ExpiresAt: expiration,
		IssuedAt:  time.Now().Unix(),
	}

	//create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	//sign the token with the private key
	return token.SignedString(config.GetPrivateKey())
}
