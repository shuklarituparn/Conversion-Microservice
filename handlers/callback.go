package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/shuklarituparn/Conversion-Microservice/configs"
	"github.com/shuklarituparn/Conversion-Microservice/user_sessions"
	"io"
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
	resp, err := client.Get("https://api.vk.com/method/users.get?fields=bdate&access_token=" + token.AccessToken + "&v=5.131")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error getting user info: %v", err))
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error reading response body: %v", err))
		return
	}

	//c.String(http.StatusOK, "User Info: %s", body)

	session, err := user_sessions.Store.Get(c.Request, "Logged_Session")
	if err != nil {

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	session.Values["authenticated"] = true

	err = sessions.Save(c.Request, c.Writer)
	if err != nil {

		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Redirect(http.StatusFound, "/dashboard")
}
