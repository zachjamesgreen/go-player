package models

import (
	"database/sql"
	"fmt"
	"log"
	db "music/database"
)


type Artist struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetArtists() []Artist {
	var artist Artist
	var artists []Artist
	sqlStatment := `SELECT * FROM artists ORDER BY id`
	rows, err := db.DB.Query(sqlStatment)
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

func GetArtist(id int) Artist {
	var artist Artist
	sqlStatment := `SELECT * FROM artists WHERE id = $1`
	row := db.DB.QueryRow(sqlStatment, id)
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

func GetArtistSongs(artist_id int) []Song {
	var song Song
	var songs []Song
	sqlStatment := `SELECT * FROM songs WHERE artist_id = $1 ORDER BY id`
	rows, err := db.DB.Query(sqlStatment, artist_id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows")
		} else {
			panic(err)
		}
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&song.Id, &song.Title, &song.Track, &song.Comment, &song.AlbumId, &song.ArtistId, &song.Genre.Name, &song.Path)
		if err != nil {
			log.Fatal(err)
		}
		// Artist{Id: id, Name: name}
		songs = append(songs, song)
	}
	return songs
}

func GetArtistAlbums(artist_id int) []Album {
	var album Album
	var albums []Album
	sqlStatment := `SELECT * FROM albums WHERE artist_id = $1 ORDER BY id`
	rows, err := db.DB.Query(sqlStatment, artist_id)
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
