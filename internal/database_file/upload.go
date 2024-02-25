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

func MongoUpload(filePath string, UserId int, UserName string, FileName string, VideoKey string) {
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
	file, err := os.Open(filePath)
	uploadOpts := options.GridFSUpload().SetMetadata(bson.D{{UserName, FileName}})

	objectID, err := bucket.UploadFromStream(FileName, io.Reader(file),
		uploadOpts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("New file uploaded with ID %s", objectID)

	//At this point we have the successful upload then we need to insert the video in teh DB and send user an Email

}
