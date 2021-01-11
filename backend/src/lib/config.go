package lib

import (
	"io/ioutil"
	"log"

	"github.com/pelletier/go-toml"
)

func GetMongoDBURI(path string) string {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	text := string(content)

	config, _ := toml.Load(text)

	uri := config.Get("mongodb.uri").(string)

	return string(uri)
}

func GetJWTSecret(path string) string {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	text := string(content)

	config, _ := toml.Load(text)

	secret := config.Get("jwt.key").(string)
	return secret
}

func GetJWTLifeTime(path string) int {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	text := string(content)

	config, _ := toml.Load(text)
	lifetime := config.Get("jwt.lifetime").(int)
	return lifetime
}
