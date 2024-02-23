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

//Video name is required, UserName is required and Filepath is required.....

//TODO: This is the code to upload, we need one to downlaod
/*
package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

var mongoClient *mongo.Client

func main() {
	// Connect to MongoDB
	mongoURL := os.Getenv("MONGO_URL")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	mongoClient = client

	// Create GridFS bucket
	bucket, err := gridfs.NewBucket(client.Database("myDB"))
	if err != nil {
		log.Fatalf("Failed to create GridFS bucket: %v", err)
	}

	// Configure Kafka consumer
	brokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	topic := os.Getenv("KAFKA_TOPIC")

	config := kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
	}

	reader := kafka.NewReader(config)
	defer reader.Close()

	// Trap SIGINT and SIGTERM to gracefully stop consumer
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-sigchan:
				log.Println("Received termination signal. Shutting down consumer...")
				return
			default:
				msg, err := reader.ReadMessage(context.Background())
				if err != nil {
					log.Printf("Error reading message from Kafka: %v", err)
					continue
				}

				log.Printf("Received message from Kafka: %s", msg.Value)

				// Upload file to MongoDB GridFS
				filePath := string(msg.Value)
				objectID, err := uploadFileToGridFS(bucket, filePath)
				if err != nil {
					log.Printf("Failed to upload file to GridFS: %v", err)
					continue
				}

				log.Printf("File uploaded to GridFS with ObjectID: %s", objectID)

				// Produce message with ObjectID
				produceMessage(objectID)
			}
		}
	}()

	// Wait for termination signal
	<-sigchan
}

func uploadFileToGridFS(bucket *gridfs.Bucket, filePath string) (primitive.ObjectID, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	opts := options.GridFSUpload().SetMetadata(bson.D{{"metadata tag", "first"}})
	objectID, err := bucket.UploadFromStream(filepath.Base(filePath), file, opts)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to upload file to GridFS: %w", err)
	}

	return objectID, nil
}

func produceMessage(objectID primitive.ObjectID) {
	// Configure Kafka producer
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  strings.Split(os.Getenv("KAFKA_BROKERS"), ","),
		Topic:    os.Getenv("KAFKA_PRODUCER_TOPIC"),
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close()

	// Produce message with ObjectID
	message := fmt.Sprintf("ObjectID: %s", objectID)
	err := writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("key"),
		Value: []byte(message),
	})
	if err != nil {
		log.Printf("Failed to produce message with ObjectID to Kafka: %v", err)
	}
}

*/
