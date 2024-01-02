package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"os"

	"go.mongodb.org/mongo-driver/bson"
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
	file, err := os.Open("./text.txt")
	uploadOpts := options.GridFSUpload().SetMetadata(bson.D{{"metadata tag", "first"}})

	objectID, err := bucket.UploadFromStream("file.txt", io.Reader(file),
		uploadOpts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New file uploaded with ID %s", objectID)

}

//TODO: This is the code to upload, we need one to downlaod
