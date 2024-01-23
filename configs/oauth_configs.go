package configs

import (
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
	"log"
	"os"
)

var conf *oauth2.Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	conf = &oauth2.Config{
		ClientID:     os.Getenv("VK_CLIENT_ID"),
		ClientSecret: os.Getenv("VK_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes:       []string{"email", "offline"},
		Endpoint:     vk.Endpoint,
	}
}

func Config() *oauth2.Config {
	return conf
}
