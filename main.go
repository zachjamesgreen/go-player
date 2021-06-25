package main

import (
	"log"
	"os"
	// "net"
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

	f, err := os.OpenFile("files/logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(f)

	db.Start()
	defer db.DB.Close()
	Router = mux.NewRouter()
	Router.PathPrefix("/song/").Handler(http.StripPrefix("/song/", http.FileServer(http.Dir("./files/"))))
	mount(Router)
	http.Handle("/", Router)
	Router.Use(auth)
	Router.Use(setHeaders)
	Router.Use(loggingMiddleware)
	// l, _ := net.Listen("tcp4", "localhost:8081")
	// s := &http.Server{
	// 	Handler: Router,
	// }
	// log.Fatal(s.Serve(l))
	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", Router))
}
