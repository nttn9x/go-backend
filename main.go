package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/nttn9x/go-backend/routers"
)

func main() {
	// Get the mux router object
	router := routers.InitRoutes()

	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    ":8000",
		Handler: n,
	}

	log.Println("Listening...")
	server.ListenAndServe()
}
