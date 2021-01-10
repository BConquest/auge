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

	err = c.Bind(u)
	if err != nil {
		log.Printf("(EE) Signup: Binding Error >>> %s\n", err)
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Malformed Input"}
	}

	/*
		Checks the "form" of the username and password. Done so different frontends
		can be made and still have the same requirements. They should also be checked
		on the frontend.
	*/
	check, response := lib.ValidateUsername(u.Username)
	if check == false {
		log.Printf("(WW) Signup: Invalid Username >>> %s\n", err)
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: response}
	}

	check, response = lib.ValidatePassword(u.Username)
	if check == false {
		log.Printf("(WW) Signup: Invalid Password >>> %s\n", err)
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: response}
	}

	return c.JSON(http.StatusCreated, u)
}
