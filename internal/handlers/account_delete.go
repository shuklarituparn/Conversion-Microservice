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

func AccountDelete(c *gin.Context) {
	session, err := user_sessions.Store.Get(c.Request, "Logged_Session") //getting the session from the session store
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	userId, ok := session.Values["UserId"].(int) //this is needed to make sure, user doesn't delete someone else's account
	if !ok {
		log.Println("Error resolving the userId from the sessions")
	}
	userName, ok := session.Values["userName"].(string)
	if !ok {
		log.Println("Error finding userID from the sessions")
	}

	SecureCode := c.Query("code")
	accountToBeDeletedStr := c.Query("userID")

	accountToBeDeleted, err := strconv.Atoi(c.Query("userID")) //if the accounts don't match
	if err != nil && accountToBeDeletedStr != "" {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID"})
		return
	}

	db := user_database.ReturnDbInstance()
	var user *models.User
	userfound, err := user_database.UserWithID(db, userId)

	if accountToBeDeletedStr != "" && SecureCode == "" { //when the user provides the query with the userID
		if userId != accountToBeDeleted { //comparing and redirecting,
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}
		if userfound { //adding the check to not delete untill the user is found
			user, err = user_database.GetUserWithID(db, userId)
			user.RestoreSecureKey = ""
			user.RestoreSecureKey = ID.ReturnID()
			db.Save(user)
			if err = db.Delete(&user).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
				return
			}
			var restoreMessage models.RestoreAccountMessage

			restoreMessage.UserName = userName
			restoreMessage.UserId = userId
			messageToSend, errorserializing := json.Marshal(restoreMessage)
			if errorserializing != nil {
				log.Println("Error serializing the message for the restore main")
			}
			session.Values["authenticated"] = false //deleting the session
			delete(session.Values, "userName")
			delete(session.Values, "userID")
			delete(session.Values, "userPhoto")
			session.Options.MaxAge = -1

			if err = session.Save(c.Request, c.Writer); err != nil {
				log.Println("Error saving session:", err)
				return
			}
			p, errorProducing := producer.NewProducer("broker:9092")
			if errorProducing != nil {
				log.Println("Erorr creating  a producer for the restore mail producer")
			}
			errorProducing = producer.ProduceNewMessage(p, "restore_mail", string(messageToSend))
			if errorProducing != nil {
				log.Println("Error producing message on the mail restore")
			}
			c.HTML(http.StatusOK, "goodbye.html", gin.H{
				"userName": userName,
			})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not in DB to delete"})

	} else if accountToBeDeletedStr == "" && SecureCode != "" {

		result := db.Unscoped().Where("restore_secure_key = ? AND id=?", SecureCode, userId).First(&user)
		if result.Error != nil {
			c.HTML(http.StatusNotFound, "404.html", gin.H{})
			return
		}
		result = db.Unscoped().Model(&user).Update("deleted", nil)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to restore user"})
			return
		}
		c.HTML(http.StatusFound, "restored.html", gin.H{})
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{})
	}

}

//This is wrong as the user wont be able to make post request from the browser
//But can give him link to the restore page that makes the post request
