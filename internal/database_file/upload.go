package database_file

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shuklarituparn/Conversion-Microservice/internal/consumer"
	"github.com/shuklarituparn/Conversion-Microservice/internal/models"
	"github.com/shuklarituparn/Conversion-Microservice/internal/producer"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_database"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func MongoUpload(filePath string, UserId int, UserName string, FileName string, VideoKey string) {
	errLoadingEnv := godotenv.Load("../../.env")
	if errLoadingEnv != nil {
		log.Print("Error opening env file to get the MONGO settings", errLoadingEnv)
	}
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
	uploadOpts := options.GridFSUpload().SetMetadata(bson.D{
		{Key: "UserName", Value: UserName},
		{Key: "FileName", Value: FileName},
		{Key: "VideoKey", Value: VideoKey},
		{Key: "VideoKey", Value: VideoKey},
	})

	objectID, err := bucket.UploadFromStream(FileName, io.Reader(file),
		uploadOpts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("New file uploaded with ID %s", objectID)

	p, err := producer.NewProducer("localhost:9092")
	if err != nil {
		log.Println("Error creating a producer after upload message", err)
	}

	//Ok now the video is upload now time to send the user the email and the download link

	//The Email template needs userName string, mode string, userID int, fileId string

	localDB := user_database.ReturnDbInstance() //storing the objID in the localDB

	result, err := user_database.GetVideoByID(localDB, VideoKey) //getting the video with the key
	if err != nil {
		log.Println("Error pulling the Video from the DB by key", err)
	}

	result.MongoDBOID = objectID.String()

	var FileDownloadMailMsg models.FiledownloadMailMessage

	FileDownloadMailMsg.FileId = objectID.String()
	FileDownloadMailMsg.Mode = result.Mode
	FileDownloadMailMsg.UserName = UserName
	FileDownloadMailMsg.UserID = UserId
	serialize, err := json.Marshal(FileDownloadMailMsg)

	localDB.Save(result)

	producer.NewProducer("localhost:9092")
	producer.ProduceNewMessage(p, "download_mail", string(serialize))

	//NEED to add a download mail consumer

	//FileDownload done and the mail producing done
	//At this point we have the successful upload then we need to insert the video in the DB and send user an Email

}

func MongoUploadConsumer() {
	//This consumer will listen for the upload topic after the conversion is done successfully
	c, _ := consumer.NewConsumer("localhost:9092", "upload_service")
	_ = c.Subscribe("upload", nil)

	defer consumer.Close(c)

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		log.Printf("Received message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		var uploadConvertedFile models.AfterConvertUpload
		err = json.Unmarshal(msg.Value, &uploadConvertedFile)
		if err != nil {
			fmt.Println(err)
		}
		MongoUpload(uploadConvertedFile.FilePath, uploadConvertedFile.UserId, uploadConvertedFile.UserName, uploadConvertedFile.FileName, uploadConvertedFile.VideoKey)

		_, commitErr := c.CommitMessage(msg)
		if commitErr != nil {
			log.Printf("Failed to commit offset: %v", commitErr)
		}
	}

	//Now the consumer listens for the upload topic and triggers the mongo

	//Mongo Uploads and produces a message

}
