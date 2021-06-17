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

func getArtist(db *sql.DB, id int) Artist {
	var artist Artist
	sqlStatment := `SELECT * FROM artists WHERE id = $1`
	row := db.QueryRow(sqlStatment, id)
	err := row.Scan(&artist.Id, &artist.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows")
		} else {
			log.Fatal(err)
		}
	}
	return artist
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

func getArtistAlbums(db *sql.DB, artist_id int) []Album {
	var album Album
	var albums []Album
	sqlStatment := `SELECT * FROM albums WHERE artist_id = $1`
	rows, err := db.Query(sqlStatment, artist_id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero Rows")
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
		albums = append(albums, album)
	}
	return albums
}
