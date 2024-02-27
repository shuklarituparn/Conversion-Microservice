package handlers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_sessions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

func Download(c *gin.Context) {
	useridQ := c.Query("userid") //can match from the cookie if it is same, if not not allow
	fileId := c.Query("fileid")  //the fileId to pull the file from the MongoDB database
	mode := c.Query("mode")
	session, err := user_sessions.Store.Get(c.Request, "Logged_Session") //getting the session from the session store
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	userName, ok := session.Values["userName"].(string)
	if !ok {
		log.Println("Error finding userID from the sessions")
	}

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

	if useridQ == "" && mode == "" && fileId == "" {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{})
		return
	}
	if mode == "Скриншоты" {
		err := godotenv.Load("/home/rituparn/Documents/Dev/languages/Conversion-Microservice/.env")
		if err != nil {
			log.Println("Error opening the env file in Screenshot upload")
		}

		mongoUrl := os.Getenv("MONGO_URL")
		opts := options.Client().ApplyURI(mongoUrl)
		client, err := mongo.Connect(context.TODO(), opts)
		if err != nil {
			log.Println("Error connecting to mongo for screenshot")
		}
		defer func(client *mongo.Client, ctx context.Context) {
			err := client.Disconnect(ctx)
			if err != nil {

			}
		}(client, context.TODO())

		db := client.Database("myDB")
		collection := db.Collection(userName) //This is correct

		objId, _ := primitive.ObjectIDFromHex(fileId)

		filter := bson.M{"_id": objId}

		var UserFile struct {
			ImageData primitive.Binary `bson:"imageData"`
			FileName  string           `bson:"fileName"`
		} //so if I did it as a type struct I will have to fill it, but var can take the value

		if err := collection.FindOne(context.TODO(), filter).Decode(&UserFile); err != nil {
			log.Println("Error fetching file from MongoDB:", err)
			return
		}
		fileName := UserFile.FileName
		completeFilePath := fmt.Sprintf("../../internal/userfiles/downloaded_files/%s", UserFile.FileName)
		if fileName == "" {
			log.Fatal("Filename not found or is empty")
		}

		err = os.WriteFile(completeFilePath, UserFile.ImageData.Data, os.ModePerm)
		log.Println("File downloaded and saved to:", fileName)
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "Binary")
		c.Header("Content-Disposition", "attachment; filename="+UserFile.FileName)
		log.Println("File downloaded and saved to:", completeFilePath)

	} else {

		// Access GridFS
		db := client.Database("myDB")
		bucket, err := gridfs.NewBucket(db)
		if err != nil {
			log.Fatal(err)
		}

		// Convert string to ObjectID
		objID, err := primitive.ObjectIDFromHex(fileId)
		if err != nil {
			log.Fatal(err)
		}

		// Get Metadata
		var result bson.M
		err = db.Collection("fs.files").FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		// Example: Use metadata to determine the file name or extension
		// This assumes 'FileName' is a key in your metadata
		fileName, ok := result["metadata"].(bson.M)["FileName"].(string)
		filenameWithPath := fmt.Sprintf("../../internal/userfiles/downloaded_files/%s", fileName)
		if !ok {
			log.Fatal("FileName in metadata not found or not a string")
		}

		// Download the file
		var buf bytes.Buffer
		_, err = bucket.DownloadToStream(objID, &buf)
		if err != nil {
			log.Fatal(err)
		}

		// Save the file
		// You might want to use the fileName from the metadata to determine the correct file extension
		err = os.WriteFile(filenameWithPath, buf.Bytes(), 0644)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("File downloaded and saved as %s\n", fileName)

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+fileName)
		c.File(filenameWithPath)

	}

}

//Just need to implement the download handler and its done
