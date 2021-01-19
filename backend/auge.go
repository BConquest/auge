package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"

	//"github.com/dgrijalva/jwt-go"

	"paxavis.dev/paxavis/auge/src/handler"
	"paxavis.dev/paxavis/auge/src/lib"
)

func allowOrigin(origin string) (bool, error) {
	return true, nil
}

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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodPost, http.MethodDelete, http.MethodGet},
	}))

	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)
	e.POST("/bookmark", handler.CreateBookmark)

	e.GET("/user/:id", handler.RequestUser)

	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

	e.Logger.Fatal(e.StartAutoTLS(":1234"))
}
