package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_database"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_sessions"
	"log"
	"net/http"
	"os"
)

func Cut(c *gin.Context) {
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

} //This is the get request for the cut page, that let's user upload the video

func CutEditResult(c *gin.Context) {
	startTime := c.PostForm("start_time") //The startTime to cut the video
	endTime := c.PostForm("end_time")
	fileName := c.PostForm("filename")
	fmt.Println(startTime, endTime, fileName)
	fmt.Println(os.Getwd())

	//At this step we get the start and the end time and the filename
	//Need to pull up the filePath

	//Why not include the VideoKey?  In this way the filePath can be pulled from DB?

}

/*
Convert: User uploads the file, the file gets uploaded to the upload folder

Now to get the file somehow to the convert edit page


*/
