package main

import (
	"log"
	"github.com/zachjamesgreen/go-player/models"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func setHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET POST DELETE OPTIONS")
		next.ServeHTTP(w, r)
	})
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/upload" {
			u, p, ok := r.BasicAuth()
			user := models.GetUserByUsername(u)
			if !ok {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			if u != user.Username {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p))
			if err != nil {
				log.Print(err)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}
		log.Print("Authorized")
		next.ServeHTTP(w, r)
	})
}
