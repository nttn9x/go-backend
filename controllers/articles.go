package controllers

import (
	"net/http"

	"github.com/nttn9x/go-backend/common"
	"github.com/nttn9x/go-backend/elastic"
	"github.com/nttn9x/go-backend/models"
)

//GetArticle Get all GetArticle
func GetArticle(w http.ResponseWriter, r *http.Request) {
	//call api, db

	data := elastic.SearchEs()
	common.RespondWithJSON(w, models.Response{Data: data})
}
