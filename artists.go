package main

import (
	"database/sql"
	"fmt"
	"log"
)

type Artist struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getArtists(db *sql.DB) []Artist {
	var artist Artist
	var artists []Artist
	sqlStatment := `SELECT * FROM artists`
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
		err := rows.Scan(&artist.Id, &artist.Name)
		if err != nil {
			log.Fatal(err)
		}
		// Artist{Id: id, Name: name}
		artists = append(artists, artist)
	}
	return artists
}

func getArtistSongs(db *sql.DB, artist_id int) []Song {
	var song Song
	var songs []Song
	sqlStatment := `SELECT * FROM songs WHERE artist_id = $1`
	rows, err := db.Query(sqlStatment, artist_id)
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
