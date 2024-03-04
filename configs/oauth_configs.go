package configs

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
	"os"
)

var conf *oauth2.Config

func init() {
	conf = &oauth2.Config{
		ClientID:     os.Getenv("VK_CLIENT_ID"),
		ClientSecret: os.Getenv("VK_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes:       []string{"email"},
		Endpoint:     vk.Endpoint,
	}
}

func OauthConfig() *oauth2.Config {
	return conf
}
