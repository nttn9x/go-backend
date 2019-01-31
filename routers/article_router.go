package routers

import (
	"github.com/gorilla/mux"

	"peakvise/controllers"
)

//SetArticlesRoutes ...
func SetArticleRoutes(router *mux.Router) {
	router.HandleFunc("/api/articles", controllers.GetArticle).Methods("GET")
}
