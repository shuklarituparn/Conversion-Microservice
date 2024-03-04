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
	"strconv"
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
	var user *models.User

	user, errorGettingUser := user_database.GetUserWithID(db, userId)
	if errorGettingUser != nil {
		log.Println("Error getting user from DB")
	}

	userEmail := user.UserEmail

	NoEmail := "Вы еще не добаили Email("
	if userEmail != "" && user.Verified == true {
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
	var user *models.User

	user, err = user_database.GetUserWithID(db, userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	VerificationToken := ID.ReturnID()
	user.VerificationToken = VerificationToken
	user.UserEmail = userEmail
	user.Verified = false

	if errorSaving := db.Save(user).Error; errorSaving != nil {
		c.AbortWithError(http.StatusInternalServerError, errorSaving)
	}
	var message models.EmailVerificationMessage

	message.UserName = user.Username
	message.UserEmail = userEmail
	message.VerificationCode = VerificationToken
	message.UserID = userId

	serializedMessage, err := json.Marshal(message)
	if err != nil {
		log.Println("Failed to Serialize the message")
	}

	p, err := producer.NewProducer("broker:9092")
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

func VerificationEmail(c *gin.Context) {
	code := c.Query("code")
	userIdStr := c.Query("userId")
	userId, _ := strconv.Atoi(userIdStr)

	if code == "" || userIdStr == "" { // OR on strings ||
		c.HTML(http.StatusBadRequest, "error.html", gin.H{})
		return
	}
	db := user_database.ReturnDbInstance()

	var user *models.User

	user, _ = user_database.GetUserWithID(db, userId)

	if user.VerificationToken == code {
		user.Verified = true
		if errorSaving := db.Save(user).Error; errorSaving != nil {
			c.AbortWithError(http.StatusInternalServerError, errorSaving)
		}
		c.Redirect(http.StatusFound, "/profile/email")

	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{})
	}
}
