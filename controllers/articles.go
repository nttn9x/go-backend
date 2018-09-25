package controllers

import (
	"net/http"

	"github.com/nttn9x/go-backend/common"
	"github.com/nttn9x/go-backend/models"
)

//GetArticle Get all GetArticle
func GetArticle(w http.ResponseWriter, r *http.Request) {
	common.RespondWithJSON(w, models.Response{Data: "Hello"})
}
