package main

import (
	"fmt"

	"paxavis.dev/paxavis/auge/src/models"
)

func main() {
	var f models.User

	f.Email = "email"

	fmt.Printf("%s\n", f.Email)
}
