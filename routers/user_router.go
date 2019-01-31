package routers

import (
	"github.com/gorilla/mux"

	"peakvise/controllers"
)

//SetUserRoutes ...
func SetUserRoutes(router *mux.Router) {
	router.HandleFunc("/api/user/login", controllers.DoLogin).Methods("POST")
}
