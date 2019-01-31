package routers

import (
	"peakvise/middlewares"

	"github.com/gorilla/mux"
)

// InitRoutes ...
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	// Routes for the Articles entity
	SetArticleRoutes(router)
	SetUserRoutes(router)

	router.Use(middlewares.LoggingMiddleware)
	return router
}
