package main

import (
	"github.com/labstack/echo"

	"paxavis.dev/paxavis/auge/src/handler"
)

func main() {
	e := echo.New()

	e.POST("/signup", handler.Signup)

	e.Logger.Fatal(e.Start(":1234"))
}
