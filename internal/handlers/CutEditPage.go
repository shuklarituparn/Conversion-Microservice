package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/internal/models"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_database"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_sessions"
	"log"
	"net/http"
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
	var video *models.Video
	db := user_database.ReturnDbInstance()

	video, err = user_database.GetLatestVideo(db, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't find the video in the DB"})
	}
	//correctFilePath, _ := url.PathUnescape(video.FilePath)
	c.HTML(http.StatusFound, "cut_edit.html", gin.H{
		"userVideo":   video.Title,
		"userpicture": userPicture,
		"userName":    userName,
	})

	//the problem is that the filename was getting resolved but one folder back

}

func CutEditResult(c *gin.Context) {

	//Function to handle the user cut time and end time and the file that he uploaded
}
