package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/dgrijalva/jwt-go"

	"paxavis.dev/paxavis/auge/src/lib"
	"paxavis.dev/paxavis/auge/src/models"
)

func Signup(c echo.Context) (err error) {
	u := &models.User{}

	err = c.Bind(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Malformed Input")
	}

	check, response := lib.ValidateUsername(u.Username)
	if check == false {
		return c.JSON(http.StatusBadRequest, response)
	}

	check, response = lib.ValidatePassword(u.Password)
	if check == false {
		return c.JSON(http.StatusBadRequest, response)
	}

	check, err = lib.CheckUsernameExists(u.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if check == true {
		return c.JSON(http.StatusBadRequest, "Username exists")
	}

	u.Password = lib.HashAndSalt([]byte(u.Password))
	u.DateCreated = time.Now()
	u.Email = ""
	u.Type = "user"

	err = lib.InsertUser(*u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, "success")
}

func Login(c echo.Context) (err error) {
	u := new(models.User)
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userP, err := lib.GetUser(u.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	check, res := lib.ComparePassword(userP.Password, u.Password)
	if res != nil || check == false {
		return c.JSON(http.StatusBadRequest, res)
	}

	key := []byte(lib.GetJWTSecret("./config.toml"))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userP.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["audience"] = userP.Type

	userP.Password = ""
	userP.Token, err = token.SignedString([]byte(key))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, userP)
}

func RequestUser(c echo.Context) (err error) {
	username := usernameFromToken(c)

	user := models.User{}
	user, err = lib.GetUser(username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user.Password = ""

	return c.JSON(http.StatusOK, user)
}
