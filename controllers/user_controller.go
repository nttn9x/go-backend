package controllers

import (
	"net/http"
	"peakvise/common"
	"peakvise/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

//DoLogin ...
func DoLogin(w http.ResponseWriter, r *http.Request) {
	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "nguyenntt",
		"exp":      time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(common.AppConfig.JWTSetting.Key))

	if err != nil {
		log.Error(err)
	}

	common.RespondWithJSON(w, models.Response{Data: tokenString})
}
