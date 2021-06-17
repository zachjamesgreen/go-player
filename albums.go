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

func getAlbum(db *sql.DB, id int) Album {
	var album Album
	sqlStatment := `SELECT * FROM albums WHERE id = $1`
	row := db.QueryRow(sqlStatment, id)
	err := row.Scan(&album.Id, &album.Title, &album.ArtistId)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows")
		} else {
			panic(err)
		}
	}
	return album
}

func getAlbumSongs(db *sql.DB, id int) []Song {
	var songs []Song
	var song Song
	sqlStatment := `SELECT * FROM songs WHERE album_id = $1`
	rows, err := db.Query(sqlStatment, id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero Rows")
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
		songs = append(songs, song)
	}
	return songs
}
