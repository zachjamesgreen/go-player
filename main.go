package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	db "music/database"
)

// var db *sql.DB

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db.Start()
	r := mux.NewRouter()
	mount(r)
	http.Handle("/", r)
	r.Use(loggingMiddleware)
	r.Use(setCookies)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
