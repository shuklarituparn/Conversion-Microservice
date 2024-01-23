package user_sessions

import (
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Store *sessions.CookieStore
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	sessionSecret := os.Getenv("SESSION_SECRET")
	if sessionSecret == "" {
		log.Fatal("SESSION_SECRET environment variable not set")
	}
	Store = sessions.NewCookieStore([]byte(sessionSecret))
}
