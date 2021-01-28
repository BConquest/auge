package handler

import (
	"net/http"
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
		return c.JSON(http.StatusBadRequest, err)
	}

	if lib.CheckLinkResponse(b.Link) == false {
		return c.JSON(http.StatusBadRequest, "invalid url")
	}

	b.User = u.ID
	b.DateCreated = time.Now()

	check, er := lib.CheckIfBookmarked(b.User, b.Link)
	if er != nil || check == false {
		return c.JSON(http.StatusAlreadyReported, "Bookmark already added")
	}

	b.Tags = append(b.Tags, "")

	err = lib.InsertBookmark(*b)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bookmark not added")
	}

	return c.JSON(http.StatusCreated, usernameFromToken(c))
}

func GetBookmarks(c echo.Context) (err error) {
	username := usernameFromToken(c)

	bookmarks, err := lib.GetUserBookmarks(username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, bookmarks)
}

func GetBookmark(c echo.Context) (err error) {
	username := usernameFromToken(c)
	id := c.Param("id")

	bookmark, err := lib.GetUserBookmark(username, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, bookmark)
}

func RemoveBookmark(c echo.Context) (err error) {
	username := usernameFromToken(c)
	id := c.Param("id")

	err = lib.RemoveBookmark(username, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, "bookmark removed")
}

func AddTag(c echo.Context) (err error) {
	username := usernameFromToken(c)
	id := c.Param("id")
	tag := c.Param("tag")

	err = lib.AddTag(username, id, tag)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, "tag added")
}

func RemoveTag(c echo.Context) (err error) {
	username := usernameFromToken(c)
	id := c.Param("id")
	tag := c.Param("tag")

	err = lib.RemoveTag(username, id, tag)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, "tag removed")
}
