package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Start() {
	// Turn on ssl mode on macos
	connStr := fmt.Sprintf("host=%s user=%s dbname=musicplayer", os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"))
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}
