package models

import (
	"database/sql"
	"fmt"
	"log"
	db "music/database"
)

type Album struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	ArtistId int    `json:"artist_id"`
	Image    bool   `json:"image"`
	Artist   string `json:"artist"`
}

func (album Album) Create(artist_id string) (album_id string, err error) {
	err = nil
	album.Image = false
	sql := `
	INSERT INTO albums (title, artist_id, image) 
	VALUES ($1, $2, $3) ON CONFLICT (title, artist_id) DO UPDATE SET title=EXCLUDED.title 
	returning id`
	db.DB.QueryRow(sql, album.Title, artist_id, album.Image).Scan(&album_id)
	return
}

func GetAlbums() (albums []Album) {
	log.Println("Getting Albums")
	var album Album
	sqlStatment := `
	SELECT albums.*, artists.name FROM albums
	JOIN artists on albums.artist_id = artists.id
	ORDER BY id`
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
		err := rows.Scan(&album.Id, &album.Title, &album.ArtistId, &album.Image, &album.Artist)
		if err != nil {
			log.Fatal(err)
		}
		// Artist{Id: id, Name: name}
		albums = append(albums, album)
	}
	return albums
}

func GetAlbum(id int) (album Album) {
	sqlStatment := `
	SELECT albums.*, artists.name FROM albums
	JOIN artists on albums.artist_id = artists.id 
	WHERE albums.id = $1`
	row := db.DB.QueryRow(sqlStatment, id)
	err := row.Scan(&album.Id, &album.Title, &album.ArtistId, &album.Image, &album.Artist)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows")
		} else {
			panic(err)
		}
	}
	return
}

func GetAlbumSongs(id int) (songs []Song) {
	var song Song
	sqlStatment := `
	SELECT s.id, s.title, s.track, s.comment, s.year, s.last_played, s.path, s.genre, s.album_id, s.artist_id, s.created_at, al.title as album_title, ar.name
	FROM songs AS s 
	JOIN albums AS al ON s.album_id = al.id 
	JOIN artists AS ar ON s.artist_id = ar.id 
	WHERE al.id = $1
	ORDER BY s.id`
	rows, err := db.DB.Query(sqlStatment, id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero Rows")
		} else {
			panic(err)
		}
	}
	defer rows.Close()
	for rows.Next() {
		// err := rows.Scan(&song.Id, &song.Title, &song.Track, &song.Comment, &song.AlbumId, &song.ArtistId, &song.Genre.Name, &song.Path)
		err := rows.Scan(&song.Id, &song.Title, &song.Track, &song.Comment, &song.Year, &song.LastPlayed, &song.Path, &song.Genre.Name, &song.AlbumId, &song.ArtistId, &song.CreatedAt, &song.Album, &song.Artist)
		if err != nil {
			log.Fatal(err)
		}
		songs = append(songs, song)
	}
	return songs
}
