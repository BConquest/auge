package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"

	"paxavis.dev/paxavis/auge/src/lib"
	"paxavis.dev/paxavis/auge/src/models"
)

func (h *Handler) Signup(c echo.Context) (err error) {
	u := &models.User{ID: bson.TypeObjectID}

	if err = c.Bind(u); err != nil {
		log.Printf("%s\n", err)
		return
	}

	if !lib.ValidUser(u) {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	lib.InsertUser(u)

	return c.JSON(http.StatusCreated, u)
}
