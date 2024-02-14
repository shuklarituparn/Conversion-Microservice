package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/internal/producer"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_database"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_sessions"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

	userId, ok := session.Values["UserId"].(uint64)
	if !ok {
		log.Println("Error resolving the userId from the sessions")
	}
	db := user_database.ReturnDbInstance()
	result, _ := user_database.UserWithID(db, userId)

	if !result {
		emailCheck, _ := user_database.EmailExits(db, userId)
		if !emailCheck {
			c.HTML(http.StatusFound, "email_redirect.html", gin.H{})
		}
	} else {
		c.HTML(http.StatusOK, "convert.html", gin.H{
			"userName":    userName,
			"userpicture": userPicture,
		})
	}

	//TODO: Now the db and everything is working, need to initialize the DB with userdata also need to add email after verify
	//TODO: After adding email, only thing left will be the conversion logic

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

		filename = fmt.Sprintf("%s_%s.%s", session.Values["userName"].(string), sanitizeFilename(fileHeader.Filename), form.Value["output_format"][0])

		newFilePath := filepath.Join(uploadDir, filename)
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
	fmt.Println("Output format:", outputFormat[0])
	p, err := producer.NewProducer("localhost:9092")
	err = producer.ProduceNewMessage(p, "email", filename)
	if err != nil {
		return
	}

	c.HTML(http.StatusOK, "convert_success.html", gin.H{})
}

func sanitizeFilename(filename string) string {
	sanitizedFilename := filepath.Base(filename)
	sanitizedFilename = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r
		}
		return '_'
	}, sanitizedFilename)
	return sanitizedFilename
}

//TODO:ADD USER_ID TO FILE FOR EASIER HANDLING

/*
Create the table in DB with userID from VK, that will make everyuser different since they are using VK

UserId
UserName
UserPic (the URL string for the image we get from VK)
...


Basically the first table will contain the field we get from the VK

Second table will store the File details for the given users
*/
