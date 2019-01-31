package controllers

import (
	"net/http"
	"peakvise/elastic"
)

//GetArticle Get all GetArticle
func GetArticle(w http.ResponseWriter, r *http.Request) {
	elastic.SearchEs(w)
}
