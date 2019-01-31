package common

import (
	colorable "github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
)

// InitializeLogging ...
func InitLogging() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{ForceColors: true, FullTimestamp: true})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(colorable.NewColorableStdout())

	// Only log the info severity or above.
	log.SetLevel(log.InfoLevel)

}
