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
	opts := options.Client().ApplyURI("mongodb+srv://ritu222:KC2vA897vg@telegrambot.xcbltlt.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
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
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

//package main
//
//import (
//	"context"
//	"fmt"
//	"os"
//	"time"
//
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/gridfs"
//	"go.mongodb.org/mongo-driver/mongo/options"
//)
//
//func main() {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
//	opts := options.Client().ApplyURI("mongodb+srv://ritu222:KC2vA897vg@telegrambot.xcbltlt.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
//
//	client, err := mongo.Connect(ctx, opts)
//	if err != nil {
//		panic(err)
//	}
//
//	defer func() {
//		if client != nil {
//			if err = client.Disconnect(ctx); err != nil {
//				panic(err)
//			}
//		}
//	}()
//
//	db := client.Database("myDB")
//	bucket, err := gridfs.NewBucket(db)
//	if err != nil {
//		panic(err)
//	}
//
//	file, err := os.Open("text.txt")
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//
//	uploadOpts := options.GridFSUpload().SetMetadata(bson.D{{"key", "value"}})
//
//	objectID, err := bucket.UploadFromStream("file.txt", file, uploadOpts)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Printf("New file uploaded with ID %s\n", objectID)
//
//	if err := client.Database("admin").RunCommand(ctx, bson.D{{"ping", 1}}).Err(); err != nil {
//		panic(err)
//	}
//	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
//}
