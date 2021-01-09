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

func getClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(GetMongoDBURI("./config.toml")))

	if err != nil {
		log.Fatal("getClient() Failed")
		log.Fatal(err)
	}

	return client
}

func CheckUsernameExists(name string) bool {
	client := getClient()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	er := client.Connect(ctx)
	if er != nil {
		log.Fatal(er)
	}
	defer client.Disconnect(ctx)

	userCollection := client.Database("development").Collection("users")

	log.Printf("> %s\n", name)
	cursor, err := userCollection.Find(
		ctx,
		bson.M{"username": name},
	)
	if err != nil {
		log.Fatal(err)
	}

	var usernamesFound []bson.M
	if err = cursor.All(ctx, &usernamesFound); err != nil {
		log.Fatal(err)
	}

	if len(usernamesFound) == 0 {
		return true
	}

	return false
}

func InsertUser(newUser models.User) {
	client := getClient()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	er := client.Connect(ctx)
	if er != nil {
		log.Fatal(er)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("development").Collection("users")
	insertResult, err := collection.InsertOne(context.TODO(), newUser)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted post with ID:", insertResult.InsertedID)
}
