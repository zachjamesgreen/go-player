package models

import (
	"database/sql"
	"fmt"
	"log"
	db "music/database"
	"time"
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
	Duration   uint64       `json:"duration"`
	CreatedAt  time.Time    `json:"created_at"`
	Liked      bool         `json:"liked"`
	LikedDate  time.Time    `json:"liked_date"`
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

func (song Song) Create(artist_id, album_id, genre, path string) (err error) {
	sqlStatment := `
	INSERT INTO songs (title, track, comment, album_id, artist_id, genre, path, year, duration) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) 
	ON CONFLICT (title, artist_id, album_id) DO UPDATE SET title=EXCLUDED.title 
	returning id`
	_, err = db.DB.Exec(sqlStatment, song.Title, song.Track, song.Comment, album_id, artist_id, genre, path, song.Year, song.Duration, time.Now(), time.Now())
	return
}

func GetSongs() []Song {
	var song Song
	var songs []Song
	sqlStatment := `
	SELECT s.id, s.title, s.track, s.comment, s.year, s.last_played, s.path, s.genre, s.album_id, s.artist_id, s.created_at, s.duration, al.title as album_title, ar.name
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
		err := rows.Scan(&song.Id, &song.Title, &song.Track, &song.Comment, &song.Year, &song.LastPlayed, &song.Path, &song.Genre.Name, &song.AlbumId, &song.ArtistId, &song.CreatedAt, &song.Duration, &song.Album, &song.Artist)
		if err != nil {
			log.Fatal(err)
		}
		songs = append(songs, song)
	}
	return songs
}

func GetLikedSongs() (songs []Song, err error) {
	song := Song{}
	sqlStatment := `
	SELECT s.id, s.title, s.track, 
				s.comment, s.year, s.last_played, 
				s.path, s.genre, s.album_id, 
				s.artist_id, s.created_at, s.duration, 
				s.liked, s.liked_date, al.title, ar.name
	FROM songs AS s
	JOIN albums AS al ON s.album_id = al.id
	JOIN artists AS ar ON s.artist_id = ar.id
	WHERE s.liked = true
	ORDER BY s.liked_date;`
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
		err := rows.Scan(
			&song.Id, &song.Title, &song.Track,
			&song.Comment, &song.Year, &song.LastPlayed,
			&song.Path, &song.Genre.Name, &song.AlbumId,
			&song.ArtistId, &song.CreatedAt, &song.Duration,
			&song.Liked, &song.LikedDate, &song.Album, &song.Artist)
		if err != nil {
			log.Fatal(err)
		}
		songs = append(songs, song)
	}
	return
}

func AddLike(id int) (err error) {
	sqlStatment := `
	UPDATE songs set liked = true WHERE id = $1`
	_, err = db.DB.Exec(sqlStatment, id)
	return
}

func RemoveLike(id int) (err error) {
	sqlStatment := `
	UPDATE songs set liked = false, liked_date = NULL WHERE id = $1`
	_, err = db.DB.Exec(sqlStatment, id)
	return
}
