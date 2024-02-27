package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/internal/models"
	"github.com/shuklarituparn/Conversion-Microservice/internal/producer"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_database"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_sessions"
	"log"
	"net/http"
)

func Screenshot(c *gin.Context) {

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
		c.HTML(http.StatusOK, "screenshot.html", gin.H{
			"userName":    userName,
			"userpicture": userPicture,
		})
	}
}

func ScreenShotResult(c *gin.Context) {
	screenshotTime := c.PostForm("time") //The startTime to cut the video
	fileName := c.PostForm("filename")
	videoKey := c.PostForm("videoKey")
	db := user_database.ReturnDbInstance()
	result, errorGettingVideo := user_database.GetVideoByID(db, videoKey)
	if errorGettingVideo != nil {
		log.Println("Error getting Video in the cut Page")
	}
	p, err := producer.NewProducer("localhost:9092")

	var ScreenshotMessage models.ScreenshotMessage
	ScreenshotMessage.UserId = result.UserID
	ScreenshotMessage.FileName = fileName
	ScreenshotMessage.FilePath = result.FilePath
	ScreenshotMessage.Time = screenshotTime
	ScreenshotMessage.UserName = result.User.Username
	serialize, err := json.Marshal(ScreenshotMessage)
	err = producer.ProduceNewMessage(p, "screenshot_video", string(serialize)) //C
	if err != nil {
		return
	}

	c.HTML(http.StatusFound, "screenshot_success.html", gin.H{})
}
