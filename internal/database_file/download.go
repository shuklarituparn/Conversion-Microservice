package database_file

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"os"
)

func downloadFile(objId string) {

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
	id, err := primitive.ObjectIDFromHex(objId)
	if err != nil {
		panic(err)
	}

	downloadStream, err := bucket.OpenDownloadStream(id)
	if err != nil {
		panic(err)
	}
	defer func(downloadStream *gridfs.DownloadStream) {
		err := downloadStream.Close()
		if err != nil {

		}
	}(downloadStream)
	file, err := os.Create("output_file.txt")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(file)

	_, err = io.Copy(file, downloadStream)
	if err != nil {
		panic(err)
	}

	println("File downloaded successfully")
}
