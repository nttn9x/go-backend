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
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./build/static/"))))

	// Serve index page on all unhandled routes
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./build/index.html")
	})
	log.Println("Listening...")
	server.ListenAndServe()
}
