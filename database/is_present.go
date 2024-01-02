package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoUrl := os.Getenv("MONGO_URL")
	opts := options.Client().ApplyURI(mongoUrl).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	db := client.Database("myDB")
	bucket, err := gridfs.NewBucket(db)
	if err != nil {
		panic(err)
	}
	filter := bson.D{{"length", bson.D{{"$gt", 1500}}}} //$gt means greater than in Mongo
	cursor, err := bucket.Find(filter)
	if err != nil {
		panic(err)
	}
	type gridfsFile struct {
		Name   string `bson:"filename"`
		Length int    `bson:"length"`
	}
	var foundFiles []gridfsFile
	if err = cursor.All(context.TODO(), &foundFiles); err != nil {
		panic(err)
	}

	for _, file := range foundFiles {
		fmt.Printf("Filename: %s, length: %d", file.Name, file.Length)
	}
}

//TODO: This is the code to upload, we need one to downlaod
