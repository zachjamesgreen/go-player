package main

import (
	"log"
	"net/http"

	db "music/database"

	"github.com/gorilla/mux"
)

// var db *sql.DB

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var Router *mux.Router

func main() {
	db.Start()
	defer db.DB.Close()
	Router = mux.NewRouter()
	mount(Router)
	http.Handle("/", Router)
	Router.Use(loggingMiddleware)
	Router.Use(setCookies)
	log.Fatal(http.ListenAndServe(":8081", Router))

}
