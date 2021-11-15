package models

import (
	"fmt"
	"log"
	db "github.com/zachjamesgreen/go-player/database"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Album struct {
	ID          int
	Title       string
	ArtistId    int
	Image       bool
	SpotifyId   string
	SpotifyLink string
	Images      datatypes.JSON
	Artist      *Artist
	Songs       []Song // on delete cascade
}

func (a Album) String() string {
	return fmt.Sprintf(
		"Album<ID: %+v | Title: %+v | Artist ID: %+v | Spotify ID: %+v | Spotify Link: %+v | Images: %+v>\n", 
		a.ID, a.Title, a.ArtistId, a.SpotifyId, a.SpotifyLink, a.Images)
}

func (album *Album) FirstOrCreate() (err error) {
	return db.DB.Where(Album{Title: album.Title, ArtistId: album.Artist.ID}).FirstOrCreate(&album).Error
}

func (album *Album) Save() (err error) {
	return db.DB.Save(&album).Error
}

func GetAlbums() (albums []Album) {
	log.Println("Getting Albums")
	err := db.DB.Preload("Artist").Find(&albums).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Zero rows")
		} else {
			panic(err)
		}
	}
	return
}

func GetAlbum(id int) (album Album) {
	err := db.DB.Find(&album, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("ErrRecordNotFound")
		} else {
			panic(err)
		}
	}
	return
}

func GetAlbumSongs(album_id int) (songs []Song, err error) {
	err = db.DB.Where("album_id = ?", album_id).Preload("Album").Preload("Artist").Order("track asc").Find(&songs).Error
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (album Album) Delete() (err error){
	return db.DB.Delete(&album).Error
}
