package lib

import (
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func ValidateUsername(username string) (bool, string) {
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

	return true, ""
}

func ValidatePassword(password string) (bool, string) {
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

func ComparePassword(hashedPassword string, password []byte) (bool, error) {
	byteHash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, password)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}
