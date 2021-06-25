package models

import (
	"database/sql"
	"fmt"
	"log"
	db "music/database"
)

type Song struct {
	Id         int          `json:"id"`
	Title      string       `json:"title"`
	Track      int          `json:"track"`
	Comment    string       `json:"comment"`
	Year       int          `json:"year"`
	LastPlayed sql.NullTime `json:"last_played"`
	ArtistId   int          `json:"artist_id"`
	AlbumId    int          `json:"album_id"`
	Artist     string       `json:"artist"`
	Album      string       `json:"album"`
	Path       string       `json:"path"`
	Genre      Genre        `json:"genre"`
}

type Genre struct {
	Name string `json:"name"`
}

func (s Song) String() string {
	return fmt.Sprintf("Song<%+v %+v %+v %+v %+v %+v %+v %+v>", s.Id, s.Title, s.Track, s.Comment, s.AlbumId, s.ArtistId, s.Path, s.Genre)
}

func (g Genre) String() string {
	return fmt.Sprintf("Genre<%+v>", g.Name)
}

func GetSongs() []Song {
	var song Song
	var songs []Song
	sqlStatment := `
	SELECT s.id, s.title, s.track, s.comment, s.year, s.last_played, s.path, s.genre, s.album_id, s.artist_id, al.title as album_title, ar.name
	FROM songs AS s 
	JOIN albums AS al ON s.album_id = al.id 
	JOIN artists AS ar ON s.artist_id = ar.id ORDER BY s.id`
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
		err := rows.Scan(&song.Id, &song.Title, &song.Track, &song.Comment, &song.Year, &song.LastPlayed, &song.Path, &song.Genre.Name, &song.AlbumId, &song.ArtistId, &song.Album, &song.Artist)
		if err != nil {
			log.Fatal(err)
		}
		songs = append(songs, song)
	}
	return songs
}
