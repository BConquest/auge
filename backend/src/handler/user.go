package handler

import (
	"log"
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
		log.Printf("(WW) Signup: Invalid Username >>> %s\n", u.Username)
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: response}
	}

	check, response = lib.ValidatePassword(u.Password)
	if check == false {
		log.Printf("(WW) Signup: Invalid Password >>> %s\n", u.Password)
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: response}
	}

	/*
		Check to make shure that the username does not already exists in the database
	*/

	check, err = lib.CheckUsernameExists(u.Username)
	if err != nil {
		log.Printf("(WW) Signup: CheckUsernameExists >>> %s\n", err)
	}

	if check == true {
		log.Printf("(WW) Signup: Username Already Exists >>> %s\n", u.Username)
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Username exists"}
	}

	/*
		Hash the password now that it is being inserted
	*/
	u.Password = lib.HashAndSalt([]byte(u.Password))

	/*
		Set the date created to be the time that it is inserted into the database
	*/
	u.DateCreated = time.Now()

	/*
		Leave email blank for now
		Update in different API call
	*/
	u.Email = ""

	u.Type = "user"

	err = lib.InsertUser(*u)
	if err != nil {
		log.Printf("(WW) Error Inserting User >>> %v\n", err)
		return err
	}

	return c.JSON(http.StatusCreated, "success")
}

func Login(c echo.Context) (err error) {
	u := new(models.User)
	if err = c.Bind(u); err != nil {
		return
	}

	userP, err := lib.GetUser(u.Username)
	if err != nil {
		log.Printf("(WW) Login: Error Getting User >>> %s\n", u.Username)
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Error Logging In"}
	}

	check, res := lib.ComparePassword(userP.Password, u.Password)
	if res != nil || check == false {
		log.Printf("(WW) Login: Wrong Password\n")
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Error Logging In"}
	}

	/*
		JWT
	*/
	key := []byte(lib.GetJWTSecret("./config.toml"))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userP.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["audience"] = userP.Type

	/*
		Set Required Fields
	*/
	userP.Password = ""
	userP.Token, err = token.SignedString([]byte(key))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userP)
}

func RequestUser(c echo.Context) (err error) {
	username := usernameFromToken(c)

	user := models.User{}
	user, err = lib.GetUser(username)
	if err != nil {
		return err
	}

	user.Password = ""

	return c.JSON(http.StatusOK, user)
}
