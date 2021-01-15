package lib

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"paxavis.dev/paxavis/auge/src/models"
)

func getConnection() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(GetMongoDBURI("./config.toml")))
	if err != nil {
		return client, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return client, err
	}

	return client, nil
}

func InsertUser(user models.User) error {
	client, err := getConnection()

	if err != nil {
		return err
	}

	userCollection := client.Database("development").Collection("users")

	res, err := userCollection.InsertOne(context.Background(), user)
	log.Printf("%v\n", res)
	if err != nil {
		return err
	}
	return nil
}

func InsertBookmark(bm models.Bookmark) error {
	client, err := getConnection()

	if err != nil {
		return err
	}

	userCollection := client.Database("development").Collection("bookmarks")

	res, err := userCollection.InsertOne(context.Background(), bm)
	log.Printf("%v\n", res)
	if err != nil {
		return err
	}
	return nil
}

func CheckIfBookmarked(username string, link string) (bool, error) {
	client, err := getConnection()
	if err != nil {
		return true, err
	}

	userCollection := client.Database("development").Collection("bookmarks")

	log.Printf("User: %v\tLink: %v\n", username, link)
	filter := bson.D{{"user", username}, {"link", link}}
	var result bson.M
	res := userCollection.FindOne(context.Background(), filter).Decode(&result)
	if res != nil {
		if res == mongo.ErrNoDocuments {
			return true, nil
		}
	}
	return false, nil
}

func CheckUsernameExists(username string) (bool, error) {
	client, err := getConnection()
	if err != nil {
		return true, err
	}

	userCollection := client.Database("development").Collection("users")

	filter := bson.D{{"username", username}}
	var result bson.M
	res := userCollection.FindOne(context.Background(), filter).Decode(&result)
	if res != nil {
		if res == mongo.ErrNoDocuments {
			return false, nil
		}
	}
	return true, nil
}

func GetUser(username string) (models.User, error) {
	var user models.User

	client, err := getConnection()
	if err != nil {
		log.Printf("(EE) GetUser: error getting connection >>> %v\n", err)
		return user, err
	}

	userCollection := client.Database("development").Collection("users")

	filter := bson.D{{"username", username}}
	res := userCollection.FindOne(context.Background(), filter).Decode(&user)
	if res != nil {
		log.Printf("(WW) GetUser: error finding user: %v\n", res)
		return user, err
	}

	return user, nil
}
