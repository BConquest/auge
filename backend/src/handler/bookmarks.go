package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo/v4"

	"paxavis.dev/paxavis/auge/src/lib"
	"paxavis.dev/paxavis/auge/src/models"
)

func usernameFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}

func CreateBookmark(c echo.Context) (err error) {
	u := &models.User{
		ID: usernameFromToken(c),
	}

	b := &models.Bookmark{}
	if err = c.Bind(b); err != nil {
		return
	}

	if b.Link == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid url"}
	}

	b.User = u.ID
	b.DateCreated = time.Now()

	check, er := lib.CheckIfBookmarked(b.User, b.Link)
	if er != nil || check == false {
		log.Printf("Bookmark Already Exists")
		return c.JSON(http.StatusAlreadyReported, "Bookmark already added")
	}

	err = lib.InsertBookmark(*b)
	if err != nil {
		return
	}

	return c.JSON(http.StatusCreated, usernameFromToken(c))
}

func GetBookmarks(c echo.Context) (err error) {
	username := usernameFromToken(c)
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 100
	}
	if page == 0 {
		page = 1
	}

	bookmarks, err := lib.GetUserBookmarks(username)

	return c.JSON(http.StatusOK, bookmarks)
}
