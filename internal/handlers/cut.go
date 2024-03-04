package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/internal/models"
	"github.com/shuklarituparn/Conversion-Microservice/internal/producer"
	"github.com/shuklarituparn/Conversion-Microservice/internal/prometheus"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_database"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_sessions"
	"log"
	"net/http"
	"net/url"
)

func Cut(c *gin.Context) {
	prometheus.CutApiPingCounter.Inc()
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
		c.HTML(http.StatusOK, "cut.html", gin.H{
			"userName":    userName,
			"userpicture": userPicture,
		})
	}

} //This is the get request for the cut page, that lets user upload the video

func CutEditResult(c *gin.Context) {
	startTime := c.PostForm("start_time") //The startTime to cut the video
	endTime := c.PostForm("end_time")
	fileName := c.PostForm("filename")
	videoKey := c.PostForm("videoKey")
	db := user_database.ReturnDbInstance()
	result, errorGettingVideo := user_database.GetVideoByID(db, videoKey)
	if errorGettingVideo != nil {
		log.Println("Error getting Video in the cut Page")
	}
	p, err := producer.NewProducer("broker:9092")

	var cutMessage models.CutMessage
	filePathforFile, _ := url.QueryUnescape(result.FilePath)
	session, err := user_sessions.Store.Get(c.Request, "Logged_Session") //getting the session from the session store
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	userName, ok := session.Values["userName"].(string)
	if !ok {
		log.Println("Error finding userID from the sessions")
	}
	cutMessage.UserName = userName
	cutMessage.FileName = fileName
	cutMessage.FilePath = filePathforFile
	cutMessage.StartTime = startTime
	cutMessage.EndTime = endTime
	cutMessage.VideoKey = videoKey
	cutMessage.UserId = result.UserID

	serialize, err := json.Marshal(cutMessage)
	err = producer.ProduceNewMessage(p, "cut_video", string(serialize)) //C
	if err != nil {
		return
	}
	c.HTML(http.StatusFound, "cut_success.html", gin.H{})

}
