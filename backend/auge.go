package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	//"github.com/dgrijalva/jwt-go"

	"paxavis.dev/paxavis/auge/src/handler"
	"paxavis.dev/paxavis/auge/src/lib"
)

func main() {
	e := echo.New()

	secret := []byte(lib.GetJWTSecret("./config.toml"))

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
	e.POST("/bookmark", handler.CreateBookmark)
    
    e.GET("/user/:id", handler.RequestUser)

	e.Logger.Fatal(e.Start(":1234"))
}
