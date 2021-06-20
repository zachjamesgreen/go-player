package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Start() {
	db, err := sql.Open("postgres", "user=zach dbname=musicplayer")
	// defer db.Close()
	if err != nil {
		panic(err)
	}
	DB = db
	return
}
