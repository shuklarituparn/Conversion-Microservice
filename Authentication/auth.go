package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
)

var (
	err = godotenv.Load("../.env")
	// Configure your OAuth2 parameters here.
	conf = &oauth2.Config{
		ClientID:     os.Getenv("VK_CLIENT_ID"),     // Replace with your Client ID
		ClientSecret: os.Getenv("VK_CLIENT_SECRET"), // Replace with your Client Secret
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes:       []string{"email, offline"}, // Scopes you require
		Endpoint:     vk.Endpoint,
	}
	// Random string for the state parameter to protect against CSRF attacks
	state = "randomwstafgsd"
)

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)

	fmt.Println("Server started at http://localhost:8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><a href='/login'>VK Login</a></body></html>")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := conf.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusFound)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != state {
		http.Error(w, "Invalid state parameter", http.StatusBadRequest)
		return
	}

	token, err := conf.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		http.Error(w, "Error exchanging token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Use the token to get user data from VK
	client := conf.Client(context.Background(), token)
	resp, err := client.Get("https://api.vk.com/method/users.get?fields=bdate&access_token=" + token.AccessToken + "&v=5.131")
	if err != nil {
		http.Error(w, "Error getting user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read and display the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User Info: %s", body)
}
