package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Start() {
	// Turn on ssl mode on macos
	connStr := fmt.Sprintf("host=%s user=%s dbname=musicplayer sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	DB = db
	return
}
