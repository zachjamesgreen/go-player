package main

import (
	"log"
	"net/http"
	"os"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func setCookies(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func auth(next http.Handler) http.Handler {
	user := os.Getenv("GO_USER")
	password := os.Getenv("GO_PASSWORD")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		if u != user || p != password {
			http.Error(w, "Forbidden", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
