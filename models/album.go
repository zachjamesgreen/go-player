package models

import (
	"database/sql"
	"fmt"
	"log"
	db "music/database"
)

type Album struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	ArtistId    int    `json:"artist_id"`
	Image       bool   `json:"image"`
	Artist      string `json:"artist"`
	SpotifyId   string `json:"spotify_id"`
	SpotifyLink string `json:"spotify_link"`
	Images      string `json:"images"`
}

func (album Album) Create(artist_id string) (album_id string, err error) {
	err = nil
	album.Image = false
	sqlstr := `
	INSERT INTO albums (title, artist_id, image, spotify_id, spotify_link, images) 
	VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (title, artist_id) DO UPDATE SET title=EXCLUDED.title
	returning id`
	err = db.DB.QueryRow(sqlstr, album.Title, artist_id, album.Image, album.SpotifyId, album.SpotifyLink, album.Images).Scan(&album_id)
	return
}

func (album Album) Update(id string) (album_id string, err error) {
	sqlstr := `
	UPDATE albums SET title=$1, image=$2, spotify_id=$3, spotify_link=$4, images=$5
	WHERE id=$6 RETURNING id`
	_, err = db.DB.Exec(sqlstr, album.Title, album.Image, album.SpotifyId, album.SpotifyLink, album.Images, id) //.Scan(&album_id)
	return
}

func (album Album) Upsert(artist_id string) (album_id string, err error) {
	sqlstr := `SELECT id FROM albums WHERE title = $1`
	err = db.DB.QueryRow(sqlstr, album.Title).Scan(&album_id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Creating Album")
			album_id, err = album.Create(artist_id)
			return
		} else {
			log.Fatal(err)
		}
	}
	log.Println("Updating Album")
	album_id, err = album.Update(artist_id)
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
		err := rows.Scan(&album.Id, &album.Title, &album.ArtistId, &album.Image, &album.SpotifyId, &album.SpotifyLink, &album.Images, &album.Artist)
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
	err := row.Scan(&album.Id, &album.Title, &album.ArtistId, &album.Image, &album.Artist, &album.SpotifyId, &album.SpotifyLink, &album.Images)
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
	SELECT s.id, s.title, s.track,
	s.comment, s.year, s.last_played,
	s.path, s.genre, s.album_id,
	s.artist_id, s.created_at, al.title,
	ar.name, s.liked, s.liked_date
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
		err := rows.Scan(
			&song.Id, &song.Title, &song.Track,
			&song.Comment, &song.Year, &song.LastPlayed,
			&song.Path, &song.Genre.Name, &song.AlbumId,
			&song.ArtistId, &song.CreatedAt, &song.Album,
			&song.Artist, &song.Liked, &song.LikedDate)
		if err != nil {
			log.Fatal(err)
		}
		songs = append(songs, song)
	}
	return songs
}
