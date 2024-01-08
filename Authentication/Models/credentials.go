package Models

import (
	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("HELLO")

var User = map[string]string{
	"user1": "password1",
	"user2": "password2", //if its not caps its not exported
} //Storing the user information

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
} //Reading the username and password from the request body

type Claims struct { //Encoding data to the JWT
	Username string `json:"username"`
	jwt.RegisteredClaims
}
