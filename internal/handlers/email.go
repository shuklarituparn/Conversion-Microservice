package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/internal/ID"
	"github.com/shuklarituparn/Conversion-Microservice/internal/models"
	"github.com/shuklarituparn/Conversion-Microservice/internal/producer"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_database"
	"github.com/shuklarituparn/Conversion-Microservice/internal/user_sessions"
	"log"
	"net/http"
)

func EmailHandler(c *gin.Context) {

	session, err := user_sessions.Store.Get(c.Request, "Logged_Session")
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
	var user models.User

	user, errorGettingUser := user_database.GetUserWithID(db, userId)
	if errorGettingUser != nil {
		log.Println("Error getting user from DB")
	}

	userEmail := user.UserEmail

	NoEmail := "Вы еще не добаили Email("
	if userEmail != "" {
		c.HTML(http.StatusFound, "email.html", gin.H{
			"userName":    userName,
			"userpicture": userPicture,
			"userEMail":   userEmail,
		})
	} else {
		c.HTML(http.StatusFound, "email.html", gin.H{
			"userName":    userName,
			"userpicture": userPicture,
			"userEMail":   NoEmail,
		})
	}

}

func EmailUpdateHandler(c *gin.Context) {
	session, err := user_sessions.Store.Get(c.Request, "Logged_Session")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	userEmail := c.PostForm("email") //getting the email from the user

	userId, ok := session.Values["UserId"].(int) //userid of the user, able to send the email
	if !ok {
		log.Println("Error resolving the userId from the sessions")
	}
	db := user_database.ReturnDbInstance() //db instance
	var user models.User

	user, err = user_database.GetUserWithID(db, userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	VerificationToken := ID.ReturnID()
	user.VerificationToken = VerificationToken
	var message models.EmailVerificationMessage

	message.UserName = user.Username
	message.UserEmail = userEmail
	message.VerificationCode = VerificationToken
	message.UserID = userId

	serializedMessage, err := json.Marshal(message)
	if err != nil {
		log.Println("Failed to Serialize the message")
	}

	p, err := producer.NewProducer("localhost:9092")
	err = producer.ProduceNewMessage(p, "verification_mail", string(serializedMessage))
	if err != nil {
		return
	}

	c.HTML(http.StatusFound, "verification_mail.html", gin.H{})

}

func EmailConfirm(c *gin.Context) {
	userEmail := c.PostForm("email") //getting user email
	c.HTML(http.StatusFound, "email_confirm.html", gin.H{
		"userEmail": userEmail,
	})
}

//TODO: SO GET REQUEST TO VERIFIED TO ENTER THE CODE

//TODO: THEN POST TO THE EMAIL IF IT MATCHES AND SHOW CONF?

//TODO: IN fact can just send the user a mail with the code in URL and click to verify
