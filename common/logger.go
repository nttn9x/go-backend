package common

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// InitializeLogging ...
func InitializeLogging() {
	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)

		// if Environment == "production" {
		log.SetFormatter(&log.JSONFormatter{})
		// } else {
		// The TextFormatter is default, you don't actually have to do this.
		// log.SetFormatter(&log.TextFormatter{})
		// }
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

}
