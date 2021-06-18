package models

import (
	"database/sql"
	"fmt"
	"log"
	db "music/database"
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

func (s Song) String() string {
	return fmt.Sprintf("Song<%+v %+v %+v %+v %+v %+v %+v %+v>", s.Id, s.Title, s.Track, s.Comment, s.AlbumId, s.ArtistId, s.Path, s.Genre)
}

func (g Genre) String() string {
	return fmt.Sprintf("Genre<%+v>", g.Name)
}

func GetSongs() []Song {
	var song Song
	var songs []Song
	sqlStatment := `SELECT * FROM songs`
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
		err := rows.Scan(&song.Id, &song.Title, &song.Track, &song.Comment, &song.AlbumId, &song.ArtistId, &song.Genre.Name, &song.Path)
		if err != nil {
			log.Fatal(err)
		}
		// Artist{Id: id, Name: name}
		songs = append(songs, song)
	}
	return songs
}
