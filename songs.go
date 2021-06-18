package main

import (
	"database/sql"
	"fmt"
	"log"
)

type Song struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Track    int    `json:"track"`
	Comment  string `json:"comment"`
	ArtistId int    `json:"artist_id"`
	AlbumId  int    `json:"album_id"`
	Path     string `json:"path"`
	Genre    Genre  `json:"genre"`
}

type Genre struct {
	Name string `json:"name"`
}

func getSongs() []Song {
	var song Song
	var songs []Song
	sqlStatment := `SELECT * FROM songs`
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
		err := rows.Scan(&song.Id, &song.Title, &song.Track, &song.Comment, &song.AlbumId, &song.ArtistId, &song.Genre)
		if err != nil {
			log.Fatal(err)
		}
		// Artist{Id: id, Name: name}
		songs = append(songs, song)
	}
	return songs
}
