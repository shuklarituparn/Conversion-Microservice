package Authentcation_Handles

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shuklarituparn/Conversion-Microservice/Authentication/Models"
	"net/http"
	"time"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	var creds Models.Credentials

	err := json.NewDecoder(r.Body).Decode(&creds) //Decode the body and decode it into the cred
	if err != nil {

		//if the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := Models.User[creds.Username] //getting the expected password from our in Memory Map

	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized) //if the password is not the same or it doesn't exist in the database
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute) //token will expire after 5 minutes

	claims := &Models.Claims{

		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(Models.JwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) //There is some unknown error
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
