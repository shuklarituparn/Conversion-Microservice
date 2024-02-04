package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_sessions"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/shuklarituparn/Conversion-Microservice/internal/producer"
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
	c.HTML(http.StatusOK, "convert.html", gin.H{
		"userName":    userName,
		"userpicture": userPicture,
	})
}

func ConvertUpload(c *gin.Context) {

	err := c.Request.ParseMultipartForm(20 << 20)
	if err != nil {
		fmt.Println("Error parsing multipart form:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse multipart form"})
		return
	}

	uploadDir := "./uploads"
	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating upload directory:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create upload directory"})
		return
	}

	form := c.Request.MultipartForm

	fileHeaders := form.File["file"]
	if len(fileHeaders) == 0 {

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "No file uploaded"})
		return
	}

	for _, fileHeader := range fileHeaders {

		file, err := fileHeader.Open()
		if err != nil {
			fmt.Println("Error opening file:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to open file"})
			return
		}
		defer file.Close()

		filename := sanitizeFilename(fileHeader.Filename)

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

	topic := "uploads"
	p, _ := producer.NewKafkaProducer()
	err = p.ProduceMessage(topic, "New Upload done")
	if err != nil {
		log.Println("Error producing message:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to produce message"})
		return
	}
	defer p.Close()
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
