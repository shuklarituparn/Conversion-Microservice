package database_file

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"os"
)

func MongoDownload(fileID string, filePath string) error {
	errLoadingEnv := godotenv.Load("../../.env")
	if errLoadingEnv != nil {
		log.Println("Error opening env file to get the MONGO settings", errLoadingEnv)
		return errLoadingEnv
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoUrl := os.Getenv("MONGO_URL")
	opts := options.Client().ApplyURI(mongoUrl).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Println("Error connecting to MongoDB:", err)
		return err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	db := client.Database("myDB")
	bucket, err := gridfs.NewBucket(db)
	if err != nil {
		log.Println("Error creating GridFS bucket:", err)
		return err
	}

	objectID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		log.Println("Error converting file ID to ObjectID:", err)
		return err
	}

	downloadStream, err := bucket.OpenDownloadStream(objectID)
	if err != nil {
		log.Println("Error opening download stream:", err)
		return err
	}
	defer downloadStream.Close()

	outFile, err := os.Create(filePath)
	if err != nil {
		log.Println("Error creating output file:", err)
		return err
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, downloadStream); err != nil {
		log.Println("Error copying data to output file:", err)
		return err
	}

	log.Println("File downloaded successfully")
	return nil
}

func MongoImageDownload(fileID string, filePath string) error {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Error opening the env file for image download:", err)
		return err
	}

	mongoUrl := os.Getenv("MONGO_URL")
	opts := options.Client().ApplyURI(mongoUrl)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Println("Error connecting to MongoDB for image download:", err)
		return err
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Println("Error disconnecting from MongoDB:", err)
		}
	}(client, context.TODO())

	db := client.Database("myDB")
	collection := db.Collection("your_collection_name") // Update with your collection name

	objectID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		log.Println("Error converting file ID to ObjectID:", err)
		return err
	}

	var result bson.M
	if err := collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&result); err != nil {
		log.Println("Error finding image in MongoDB:", err)
		return err
	}

	imageData := result["imageData"].(string) // Assuming the imageData field is stored as a string

	if err := os.WriteFile(filePath, []byte(imageData), 0644); err != nil {
		log.Println("Error writing image data to file:", err)
		return err
	}

	log.Println("Image downloaded successfully")
	return nil
}
