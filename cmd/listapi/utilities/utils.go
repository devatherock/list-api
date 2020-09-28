package utilities

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func LogHandler(function http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		function(writer, request)
		log.Printf("Request: %s %s. Time taken: %s", request.Method, request.URL.Path, time.Since(start))
	}
}
