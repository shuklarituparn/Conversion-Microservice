package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/internal/ID"
	"github.com/shuklarituparn/Conversion-Microservice/internal/models"
	"github.com/shuklarituparn/Conversion-Microservice/internal/producer"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_database"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_sessions"
	"gorm.io/gorm"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

func CutEditPage(c *gin.Context) {
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

	err = c.Request.ParseMultipartForm(20 << 20) //getting the file from user
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

		filename = fmt.Sprintf("%s_cut_%s.mp4", session.Values["userName"].(string), SanitizeFilename(fileHeader.Filename))

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

	p, err := producer.NewProducer("localhost:9092")
	err = producer.ProduceNewMessage(p, "cut_file", filename)
	if err != nil {
		return
	}

	db := user_database.ReturnDbInstance() //getting db, now will store the video
	filePathofVideo := fmt.Sprintf("../uploads/%s", filename)
	encodedFilePath := url.PathEscape(filePathofVideo) //to encode the filepath
	var existingVideo models.Video

	result := db.Where("title=?", filename).First(&existingVideo)
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "video already exists"})
		return
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error checking for existing video"})
		return
	}

	videoKey := ID.ReturnID()
	video := models.Video{
		UserID:     userId,
		Title:      filename,
		FilePath:   encodedFilePath,
		MongoDBOID: "",
		CreatedAt:  time.Now(),
		Mode:       "cut",
		VideoKey:   videoKey,
	}

	createVideo := db.Create(&video)
	if createVideo.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating video"})
		return
	}

	//correctFilePath, _ := url.PathUnescape(video.FilePath)
	c.HTML(http.StatusFound, "cut_edit.html", gin.H{
		"userVideo":   filename,
		"userpicture": userPicture,
		"userName":    userName,
	})

	//Now we have the video saved in the db
}
