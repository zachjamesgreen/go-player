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
	sqlStatment := `
	SELECT s.id, s.title, s.track, s.comment, s.year, s.last_played, s.path, s.genre, s.album_id, s.artist_id, al.title as album_title, ar.name
	FROM songs AS s 
	JOIN albums AS al ON s.album_id = al.id 
	JOIN artists AS ar ON s.artist_id = ar.id 
	WHERE ar.id = $1
	ORDER BY s.id`
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
		// err := rows.Scan(&song.Id, &song.Title, &song.Track, &song.Comment, &song.AlbumId, &song.ArtistId, &song.Genre.Name, &song.Path, &song.Year, &song.LastPlayed)
		err := rows.Scan(&song.Id, &song.Title, &song.Track, &song.Comment, &song.Year, &song.LastPlayed, &song.Path, &song.Genre.Name, &song.AlbumId, &song.ArtistId, &song.Album, &song.Artist)
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
		err := rows.Scan(&album.Id, &album.Title, &album.ArtistId, &album.Artist)
		if err != nil {
			log.Fatal(err)
		}
		albums = append(albums, album)
	}
	return albums
}
