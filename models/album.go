package models

import (
	"fmt"
	"log"
	db "music/database"

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
	Songs       []Song
}

func (a Album) String() string {
	return fmt.Sprintf(
		"Album<ID: %+v | Title: %+v | Artist ID: %+v | Spotify ID: %+v | Spotify Link: %+v | Images: %+v>\n", 
		a.ID, a.Title, a.ArtistId, a.SpotifyId, a.SpotifyLink, a.Images)
}

func (album *Album) Upsert() (err error) {
	err = db.DB.Where(Album{Title: album.Title}).FirstOrInit(&album).Error
	if err != nil {
		return err
	}
	if album.ID == 0 {
		err = db.DB.Create(&album).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (album *Album) Save() (err error) {
	err = db.DB.Save(&album).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("ErrRecordNotFound")
		} else {
			return
		}
	}
	return
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

func GetAlbumSongs(album_id int) (songs []Song) {
	err := db.DB.Where("album_id = ?", album_id).Preload("Album").Preload("Artist").Order("track asc").Find(&songs).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("ErrRecordNotFound")
		} else {
			panic(err)
		}
	}
	return
}

func (album Album) Delete() {
	err := db.DB.Select("Songs").Delete(&album).Error
	if err != nil {
		panic(err)
	}
}
