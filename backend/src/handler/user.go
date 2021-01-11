package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"

	"github.com/cristalhq/jwt/v3"

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
	log.Printf("(II) Signup: User >>> %s\n", u)

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

	err = lib.InsertUser(*u)

	return c.JSON(http.StatusCreated, u)
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
	signer, err := jwt.NewSignerHS(jwt.HS256, key)
	if err != nil {
		log.Printf("(WW) Login: error getting new jwt signer >>> %v\n", err)
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Error Logging In"}
	}

	claims := &jwt.StandardClaims{
		Audience:  []string{"user"},
		ID:        u.username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
	}

	builder := jwt.NewBuilder(signer)
	token, err := builder.Build(claims)
	if err != nil {
		log.Printf("(WW) Login: error building claims >>> %v\n", err)
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Error Logging In"}
	}

	u.Token = token.String()
	u.Password = ""

	return c.JSON(http.StatusOK, u)
}

func AddBookMark(c echo.Context) (err error) {

	return
}
