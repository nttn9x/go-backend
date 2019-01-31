package main

import (
	"fmt"
	"net/http"

	common "peakvise/common"
	"peakvise/routers"

	log "github.com/sirupsen/logrus"
)

func main() {
	common.StartUp()

	// Get the mux router object
	router := routers.InitRoutes()

	server := &http.Server{
		Addr:    common.AppConfig.Port,
		Handler: router,
	}

	log.Info("Starting the development server...")
	log.Info(fmt.Sprintf("Use Port: %s", common.AppConfig.Port))

	server.ListenAndServe()
}
