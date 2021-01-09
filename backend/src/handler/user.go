package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	//"go.mongodb.org/mongo-driver/bson"

	"paxavis.dev/paxavis/auge/src/lib"
	"paxavis.dev/paxavis/auge/src/models"
)

func Signup(c echo.Context) (err error) {
	u := &models.User{}

	log.Printf(">> %v\n", c)
	err = c.Bind(u)
	if err != nil {
		log.Printf("(EE) Signup: %s\n", err)
		return
	}

	nu, err := lib.CreateUser(*u)
	if err != nil {
		log.Printf("%v\n", err)
		return &echo.HTTPError{Code: http.StatusBadRequest,
			Message: "invalid email or password"}
	}

	lib.InsertUser(nu)

	return c.JSON(http.StatusCreated, nu)
}
