package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/internal/ID"
	"github.com/shuklarituparn/Conversion-Microservice/internal/models"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_database"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_sessions"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

func Screenshot_edit(c *gin.Context) {
	//Here we will get the post request with the file

	//Just use the convert handler

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

	fileHeaders := form.File["file"] //getting the file
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
	userPicture, ok := session.Values["userPhoto"].(string)
	if !ok {
		log.Println("Error finding userID from the sessions")
	}
	db := user_database.ReturnDbInstance() //getting db, now will store the video

	filePathofVideo := fmt.Sprintf("../uploads/%s", filename)
	encodedFilePath := url.PathEscape(filePathofVideo) //to encode the filepath
	Videokey := ID.ReturnID()                          //In this way can pass the videoKey to the producer
	video := models.Video{
		UserID:     userId,
		Title:      filename,
		FilePath:   encodedFilePath,
		MongoDBOID: "",
		CreatedAt:  time.Now(),
		Mode:       "Скриншоты",
		VideoKey:   Videokey, //In this way all the video will be unique
	}

	createVideo := db.Create(&video) //here we have the video to screenshot
	if createVideo.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating video"})
		return
	}
	c.HTML(http.StatusOK, "screenshot_edit.html", gin.H{
		"userpicture": userPicture,
		"userName":    userName,
		"userVideo":   filename,
		"videoKey":    Videokey,
	})

}
