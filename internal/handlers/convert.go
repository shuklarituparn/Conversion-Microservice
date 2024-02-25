package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/internal/ID"
	"github.com/shuklarituparn/Conversion-Microservice/internal/models"
	"github.com/shuklarituparn/Conversion-Microservice/internal/producer"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_database"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_sessions"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Convert(c *gin.Context) {
	session, err := user_sessions.Store.Get(c.Request, "Logged_Session") //getting the session from the session store
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	userName, ok := session.Values["userName"].(string)
	if !ok {
		log.Println("Error finding userID from the sessions")
	}
	userPicture, ok := session.Values["userPhoto"].(string)
	if !ok {
		log.Println("Error finding userID from the sessions")
	}

	userId, ok := session.Values["UserId"].(int)
	if !ok {
		log.Println("Error resolving the userId from the sessions")
	}
	db := user_database.ReturnDbInstance()
	VerfiedCheck, _ := user_database.IsVerified(db, userId)
	if !VerfiedCheck {
		c.HTML(http.StatusFound, "email_redirect.html", gin.H{})

	} else {
		c.HTML(http.StatusOK, "convert.html", gin.H{
			"userName":    userName,
			"userpicture": userPicture,
		})
	}
}

func ConvertUpload(c *gin.Context) {

	err := c.Request.ParseMultipartForm(20 << 20)
	if err != nil {
		fmt.Println("Error parsing multipart form:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse multipart form"})
		return
	}

	uploadDir := "../../uploads"
	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating upload directory:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create upload directory"})
		return
	}

	form := c.Request.MultipartForm

	var filename string
	var newFilePath string

	fileHeaders := form.File["file"]
	if len(fileHeaders) == 0 {

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "No file uploaded"})
		return
	}

	for _, fileHeader := range fileHeaders { //Why are we looping over it?

		file, err := fileHeader.Open()
		if err != nil {
			fmt.Println("Error opening file:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to open file"})
			return
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)
		session, err := user_sessions.Store.Get(c.Request, "Logged_Session") //getting the session from the session store

		filename = fmt.Sprintf("%s_%s", session.Values["userName"].(string), SanitizeFilename(fileHeader.Filename))

		newFilePath = filepath.Join(uploadDir, filename)
		newFile, err := os.Create(newFilePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create file"})
			return
		}
		defer func(newFile *os.File) {
			err := newFile.Close()
			if err != nil {

			}
		}(newFile)

		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Println("Error copying uploaded file contents:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to copy uploaded file contents"})
			return
		}
	}

	outputFormat := form.Value["output_format"]
	if len(outputFormat) == 0 {
		fmt.Println("No output format provided")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "No output format provided"})
		return
	}
	session, err := user_sessions.Store.Get(c.Request, "Logged_Session")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	userId, ok := session.Values["UserId"].(int)
	if !ok {
		log.Println("Error resolving the userId from the sessions")
	}
	userName, ok := session.Values["userName"].(string)
	if !ok {
		log.Println("Error finding userID from the sessions")
	}
	db := user_database.ReturnDbInstance() //getting db, now will store the video

	//var existingVideo models.Video

	//result := db.Where("title=?", filename).First(&existingVideo)
	//if result.Error == nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "video already exists"})
	//	return
	//} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "error checking for existing video"})
	//	return
	//}
	filePathofVideo := fmt.Sprintf("../uploads/%s", filename)
	encodedFilePath := url.PathEscape(filePathofVideo) //to encode the filepath
	Videokey := ID.ReturnID()                          //In this way can pass the videoKey to the producer
	video := models.Video{
		UserID:     userId,
		Title:      filename,
		FilePath:   encodedFilePath,
		MongoDBOID: "",
		CreatedAt:  time.Now(),
		Mode:       "Конвертировать",
		VideoKey:   Videokey, //In this way all the video will be unique
	}

	createVideo := db.Create(&video)
	if createVideo.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating video"})
		return
	}
	c.HTML(http.StatusOK, "convert_success.html", gin.H{})

	//Now we have the video saved in the db
	//Need to produce a message on the convert topic and then let the convert handler convert it

	//After convert it produces a message on the upload, and uploads the video, returns obj Id and save it in the DB

	fmt.Println("Output format:", outputFormat[0])

	conversionMessage := models.ConversionMessage{
		UserId:       userId,
		UserName:     userName,
		FileName:     filename,
		FilePath:     filePathofVideo,
		OutputFormat: outputFormat[0],
		VideoKey:     Videokey,
	}
	serializedMessage, errorSerializingConvMessage := json.Marshal(conversionMessage)
	if errorSerializingConvMessage != nil {
		log.Println("Error Creating Conversion Message: ", errorSerializingConvMessage)
	}
	p, err := producer.NewProducer("localhost:9092")
	err = producer.ProduceNewMessage(p, "conversion", string(serializedMessage))
	if err != nil {
		return
	}

}

func SanitizeFilename(filename string) string {
	filebase := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
	fileExt := filepath.Ext(filename)

	sanitizedFilename := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r
		}
		return '_'
	}, filebase)
	sanitizedFilename += fileExt
	return sanitizedFilename
}

//Now need to make the logic for the post request // SO here the file comes
