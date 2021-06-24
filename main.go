package main

import (
	"log"
	"net/http"

	db "music/database"

	"github.com/gorilla/mux"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var Router *mux.Router

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	db.Start()
	defer db.DB.Close()
	Router = mux.NewRouter()
	Router.PathPrefix("/song/").Handler(http.StripPrefix("/song/", http.FileServer(http.Dir("./files/"))))
	mount(Router)
	http.Handle("/", Router)
	Router.Use(loggingMiddleware)
	Router.Use(setCookies)
	Router.Use(auth)
	log.Fatal(http.ListenAndServe(":8081", Router))
}
