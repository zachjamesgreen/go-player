package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var err error
	db, err = sql.Open("postgres", "host=127.0.0.1 user=zach dbname=musicplayer")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	mount(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
