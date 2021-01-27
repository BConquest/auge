package lib

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	client.Disconnect(context.Background())

	if err != nil {
		return err
	}
	_ = res
	return nil
}

func InsertBookmark(bm models.Bookmark) error {
	client, err := getConnection()

	if err != nil {
		return err
	}

	userCollection := client.Database("development").Collection("bookmarks")

	res, err := userCollection.InsertOne(context.Background(), bm)
	client.Disconnect(context.Background())

	if err != nil {
		return err
	}
	_ = res
	return nil
}

func CheckIfBookmarked(username string, link string) (bool, error) {
	client, err := getConnection()
	if err != nil {
		return true, err
	}

	userCollection := client.Database("development").Collection("bookmarks")

	filter := bson.M{"user": username, "link": link}
	var result bson.M
	res := userCollection.FindOne(context.Background(), filter).Decode(&result)
	client.Disconnect(context.Background())

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

	filter := bson.M{"username": username}
	var result bson.M
	res := userCollection.FindOne(context.Background(), filter).Decode(&result)
	client.Disconnect(context.Background())

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
		return user, err
	}

	userCollection := client.Database("development").Collection("users")

	filter := bson.M{"username": username}
	res := userCollection.FindOne(context.Background(), filter).Decode(&user)
	client.Disconnect(context.Background())

	if res != nil {
		return user, err
	}

	return user, nil
}

func GetUserBookmarks(username string) ([]models.Bookmark, error) {
	var bookmarks []models.Bookmark

	client, err := getConnection()
	if err != nil {
		return bookmarks, err
	}

	collection := client.Database("development").Collection("bookmarks")

	filter := bson.M{"user": username}
	cur, err := collection.Find(context.Background(), filter)

	if err != nil {
		return bookmarks, err
	}
	for cur.Next(context.Background()) {
		var b models.Bookmark
		err := cur.Decode(&b)
		if err != nil {
			client.Disconnect(context.Background())
			return bookmarks, err
		}

		bookmarks = append(bookmarks, b)
	}
	cur.Close(context.Background())

	client.Disconnect(context.Background())
	return bookmarks, nil
}

func GetUserBookmark(username string, id string) (models.Bookmark, error) {
	var bookmark models.Bookmark

	client, err := getConnection()
	if err != nil {
		return bookmark, err
	}

	collection := client.Database("development").Collection("bookmarks")

	t, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"user": username, "_id": t}
	res := collection.FindOne(context.Background(), filter).Decode(&bookmark)

	client.Disconnect(context.Background())
	if res != nil {
		return bookmark, res
	}

	return bookmark, res
}

func RemoveBookmark(username string, id string) error {
	client, err := getConnection()
	if err != nil {
		return err
	}

	collection := client.Database("development").Collection("bookmarks")

	t, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": t, "username": username}
	res, err := collection.DeleteOne(context.Background(), filter)
	client.Disconnect(context.Background())

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no bookmark deleted")
	}

	return nil
}

func AddTag(username string, id string, tag string) error {
	client, err := getConnection()
	if err != nil {
		return err
	}

	collection := client.Database("development").Collection("bookmarks")

	t, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": t, "user": username}
	update := bson.D{
		{"$addToSet", bson.D{
			{"tags", tag},
		}},
	}

	updateResult, err := collection.UpdateOne(context.Background(), filter, update)

	client.Disconnect(context.Background())
	if err != nil {
		return err
	}

	log.Printf("%v\n", updateResult)
	if updateResult.MatchedCount == 1 && updateResult.ModifiedCount == 0 {
		return errors.New("tag already added")
	}
	return nil
}

func RemoveTag(username string, id string, tag string) error {
	client, err := getConnection()
	if err != nil {
		return err
	}

	collection := client.Database("development").Collection("bookmarks")

	t, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": t, "user": username}
	update := bson.D{
		{"$pull", bson.M{
			"tags": tag,
		}},
	}

	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	client.Disconnect(context.Background())

	if err != nil {
		return err
	}

	log.Printf("%v\n", updateResult)
	if updateResult.MatchedCount == 1 && updateResult.ModifiedCount == 0 {
		return errors.New("tag already added")
	}
	return nil
}
