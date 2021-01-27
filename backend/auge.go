package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	//"golang.org/x/crypto/acme/autocert"

	"paxavis.dev/paxavis/auge/src/handler"
	"paxavis.dev/paxavis/auge/src/lib"
)

func allowOrigin(origin string) (bool, error) {
	return true, nil
}

func main() {
	e := echo.New()

	secret := []byte(lib.GetJWTSecret("./config.toml"))

	//e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

	//e.Pre(middleware.HTTPSRedirect())

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
	e.POST("/bookmark/:id/:tag", handler.AddTag)

	e.GET("/user/:id", handler.RequestUser)
	e.GET("/bookmark", handler.GetBookmarks)
	e.GET("/bookmark/:id", handler.GetBookmark)

	e.DELETE("/bookmark/:id", handler.RemoveBookmark)

	//e.Logger.Fatal(e.StartAutoTLS(":1235"))
	e.Logger.Fatal(e.Start(":1235"))
}
