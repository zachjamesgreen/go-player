package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db, e := sql.Open("postgres", "user=zach dbname=musicplayer")
	defer db.Close()
	if e != nil {
		panic(e)
	}
	r := mux.NewRouter()
	mount(r, db)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8081", nil))

}
