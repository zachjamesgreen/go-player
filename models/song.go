package models

import (
	"database/sql"
	"fmt"
	db "github.com/zachjamesgreen/go-player/database"
	"time"

	"gorm.io/gorm"
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
	return fmt.Sprintf("Song<%+v %+v %+v %+v %+v %+v %+v %+v>\n", s.ID, s.Title, s.Track, s.Comment, s.AlbumId, s.ArtistId, s.Path, s.Genre)
}

func (song *Song) Upsert() (err error){
	err = db.DB.Where(song).FirstOrInit(&song).Error
	if err != nil {
		return
	}
	if song.ID == 0 {
		err = db.DB.Create(&song).Error
		if err != nil {
			return
		}
	}
	return
}

func (song *Song) Save() {
	err := db.DB.Save(song).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("ErrRecordNotFound")
		} else {
			panic(err)
		}
	}
}

func (song Song) Delete() {
	err := db.DB.Delete(&song).Error
	if err != nil {
		panic(err)
	}
}

func GetSongs() (songs []Song) {
	err := db.DB.Preload("Artist").Preload("Album").Find(&songs).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("ErrRecordNotFound")
		} else {
			panic(err)
		}
	}
	return
}

func GetSong(id int) (song Song) {
	err := db.DB.First(&song, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("ErrRecordNotFound")
		} else {
			panic(err)
		}
	}
	return
}

func GetLikedSongs() (songs []Song) {
	err := db.DB.Where("liked = ?", true).Find(&songs).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("ErrRecordNotFound")
		} else {
			panic(err)
		}
	}
	return
}

func AddLike(id int) (err error) {
	err = db.DB.Model(&Song{}).Where("id = ?", id).Update("liked", true).Error
	return
}

func RemoveLike(id int) (err error) {
	err = db.DB.Model(&Song{}).Where("id = ?", id).Update("liked", false).Error
	return

}
