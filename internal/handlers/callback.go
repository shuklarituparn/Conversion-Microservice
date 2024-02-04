package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/shuklarituparn/Conversion-Microservice/configs"
	"github.com/shuklarituparn/Conversion-Microservice/user_sessions"
	"io"
	"log"
	"net/http"
)

var (
	conf = configs.Config()
)

func Callback(c *gin.Context) {
	if c.Query("state") != state {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid state parameter"))
		return
	}

	token, err := conf.Exchange(context.Background(), c.Query("code"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error exchanging token: %v", err))
		return
	}
	client := conf.Client(context.Background(), token)
	resp, err := client.Get("https://api.vk.com/method/users.get?fields=bdate,photo_max_orig&access_token=" + token.AccessToken + "&v=5.131")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error getting user info: %v", err))
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	//body, err := io.ReadAll(resp.Body)
	//userdata := string(body)
	//c.String(http.StatusFound, userdata)

	result := struct {
		Response []struct {
			UserId    int    `json:"id"`
			UserPhoto string `json:"photo_max_orig"`
			UserName  string `json:"first_name"`
		} `json:"response"`
	}{}

	if errGettingVKresponse := json.NewDecoder(resp.Body).Decode(&result); errGettingVKresponse != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error decoding user info: %v", err))
		return
	}

	session, err := user_sessions.Store.Get(c.Request, "Logged_Session")
	if err != nil {

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if len(result.Response) > 0 {
		userId := result.Response[0].UserId
		userPic := result.Response[0].UserPhoto
		userName := result.Response[0].UserName
		session.Values["authenticated"] = true
		session.Values["UserId"] = userId
		session.Values["userPhoto"] = userPic
		session.Values["userName"] = userName

	} else {
		log.Println("Error getting user data")

	}
	err = sessions.Save(c.Request, c.Writer)
	if err != nil {

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusFound, "/dashboard")
}
