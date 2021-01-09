package lib

import (
	"errors"
	"log"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"

	"paxavis.dev/paxavis/auge/src/models"
)

func checkUsername(username string) (bool, string) {
	// Must be in between 3 and 15 characters
	// Must be alphanumeric

	var alphaNumeric = regexp.MustCompile(`^[a-zA-Z0-9_]*$`).MatchString

	runeUsername := []rune(username)
	if len(runeUsername) < 3 || len(runeUsername) > 15 {
		return false, "Username is not the correct length"
	}

	if !alphaNumeric(username) {
		return false, "Username is not alphanumeric"
	}

	if !CheckUsernameExists(username) {
		return false, "Username exists"
	}

	return true, ""
}

func checkPassword(password string) (bool, string) {
	// Must be at least 8 characters long
	runePassword := []rune(password)
	if len(runePassword) < 8 {
		return false, "Password not long enough"
	}
	return true, ""
}

func hashAndSalt(password []byte) string {
	// Use GenerateFromPassword from bcrypt

	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func CreateUser(u models.User) (models.User, error) {
	var check, statement = checkUsername(u.Username)

	if !check {
		return u, errors.New(statement)
	}
	check, statement = checkPassword(u.Password)
	if !check {
		return u, errors.New(statement)
	}

	u.Password = hashAndSalt([]byte(u.Password))
	u.DateCreated = time.Now()

	return u, nil
}

func LoginUser(hashedPassword string, password []byte) (bool, error) {
	byteHash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, password)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}
