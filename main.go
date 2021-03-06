package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	db "github.com/zachjamesgreen/go-player/database"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var Router *mux.Router

func main() {

	f, err := os.OpenFile("files/logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(f)

	db.Start()
	Router = mux.NewRouter()
	Router.PathPrefix("/song/").Handler(http.StripPrefix("/song/", http.FileServer(http.Dir("./files/"))))
	mount(Router)
	http.Handle("/", Router)
	Router.Use(auth)
	// Router.Use(setHeaders)
	methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	Router.Use(loggingMiddleware)
	log.Println("Starting Server")
	fmt.Println("Starting Server")
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", handlers.CORS(methods, origins)(Router)))
}
