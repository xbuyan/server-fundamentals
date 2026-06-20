package middleware

import (
	"fmt"
	"log"
	"net/http"
)

func sendError(w http.ResponseWriter, code int, message string) {

	w.WriteHeader(code)
	fmt.Fprintf(w, message)
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Printf("%s %s", r.Method, r.URL.Path)
		next(w, r)

	}
}

func MethodMiddleware(next http.HandlerFunc, allowedMethod string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != allowedMethod {

			sendError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}
		next(w, r)

	}
}
