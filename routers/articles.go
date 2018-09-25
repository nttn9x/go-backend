package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/nttn9x/go-backend/controllers"
)

func SetArticlesRoutes(router *mux.Router) *mux.Router {
	articlesRouter := mux.NewRouter()
	articlesRouter.HandleFunc("/api/articles", controllers.GetArticle).Methods("GET")
	router.PathPrefix("/api/articles").Handler(negroni.New(
		//negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(articlesRouter),
	))
	return router
}
