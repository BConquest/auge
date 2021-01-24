package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"paxavis.dev/paxavis/auge/src/handler"
	"paxavis.dev/paxavis/auge/src/lib"
)

func allowOrigin(origin string) (bool, error) {
	return true, nil
}

func main() {
	e := echo.New()

	secret := []byte(lib.GetJWTSecret("./config.toml"))

	e.Use(middleware.CORS())

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		ContextKey:    "user",
		SigningKey:    secret,
		SigningMethod: middleware.AlgorithmHS256,
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)
	e.POST("/addbookmark", handler.CreateBookmark)

	e.GET("/user/:id", handler.RequestUser)
	e.GET("/bookmark", handler.GetBookmarks)

	e.Logger.Fatal(e.Start(":1234"))
}
