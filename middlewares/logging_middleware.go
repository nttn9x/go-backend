package middlewares

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type customeResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *customeResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

//LoggingMiddleware ...
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		sw := customeResponseWriter{ResponseWriter: w}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(&sw, r)

		log.Printf("Request finished %s %s %s in %s %d", r.Proto, r.Method, r.RequestURI, time.Now().Sub(start).String(), sw.status)
	})
}
