package models

import (
	"database/sql"
	"fmt"
	"time"

	db "github.com/zachjamesgreen/go-player/database"
)

type Song struct {
	ID         int
	Title      string
	Track      int
	Comment    string
	Year       int
	LastPlayed sql.NullTime
	ArtistId   int
	AlbumId    int
	Path       string
	Genre      string
	Duration   uint64
	CreatedAt  time.Time
	Liked      bool
	LikedDate  time.Time
	Artist     *Artist
	Album      *Album
	// Plays 		int
}

func (s Song) String() string {
	return fmt.Sprintf("Song<ID: %+v\n Title: %+v\n Track: %+v\n Comment: %+v\n AlbumId: %+v\n ArtistId: %+v\n Path: %+v\n Genre: %+v>\n", s.ID, s.Title, s.Track, s.Comment, s.AlbumId, s.ArtistId, s.Path, s.Genre)
}

func (song *Song) Create() (err error) {
	return db.DB.Create(&song).Error
}

func (song *Song) FirstOrCreate() (err error) {
	return db.DB.Where(Song{Title: song.Title, ArtistId: song.Artist.ID, AlbumId: song.Album.ID}).FirstOrCreate(&song).Error
}

func (song *Song) Save() (err error) {
	return db.DB.Save(song).Error
}

func (song Song) Delete() (err error) {
	return db.DB.Delete(&song).Error
}

func GetSongs() (songs []Song, err error) {
	err = db.DB.Preload("Artist").Preload("Album").Find(&songs).Error
	return
}

func GetSong(id int) (song Song, err error) {
	err = db.DB.First(&song, id).Error
	return
}

func GetLikedSongs() (songs []Song, err error) {
	err = db.DB.Where("liked = ?", true).Find(&songs).Error
	return
}

func (song *Song) AddLike() (err error) {
	return db.DB.First(&song).Update("liked", true).Error
}

func (song *Song) RemoveLike() (err error) {
	return db.DB.First(&song).Update("liked", false).Error
}
