package main

import (
	"database/sql"
	"fmt"
	"log"
)

type Album struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	ArtistId int    `json:"artist_id"`
}

func getAlbums(db *sql.DB) []Album {
	var album Album
	var albums []Album
	sqlStatment := `SELECT * FROM albums`
	rows, err := db.Query(sqlStatment)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows")
		} else {
			panic(err)
		}
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&album.Id, &album.Title, &album.ArtistId)
		if err != nil {
			log.Fatal(err)
		}
		// Artist{Id: id, Name: name}
		albums = append(albums, album)
	}
	return albums
}
