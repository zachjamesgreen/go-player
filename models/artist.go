package models

import (
	"fmt"
	"log"
	db "music/database"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Artist struct {
	ID            int
	Name          string
	Albums        []Album
	SpotifyId     string
	Images datatypes.JSON
}

func (a Artist) String() string {
	return fmt.Sprintf(
		"Artist<ID: %+v | Name: %+v>\n", a.ID, a.Name)
}

func (artist *Artist) Upsert() {
	err := db.DB.FirstOrCreate(&artist, Artist{Name: artist.Name}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Zero rows")
		} else {
			panic(err)
		}
	}
}

func (artist *Artist) Save() {
	err := db.DB.Save(&artist).Error
	if err != nil {
		panic(err)
	}
}

func GetArtists() (artists []Artist) {
	err := db.DB.Find(&artists).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Zero rows")
		} else {
			panic(err)
		}
	}
	return
}

func GetArtist(id int) (artist Artist) {
	err := db.DB.Find(&artist, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Zero rows")
		} else {
			log.Fatal(err)
		}
	}
	return
}

func GetArtistSongs(artist_id int) (songs []Song) {
	err := db.DB.Where("artist_id = ?", artist_id).Preload("Artist").Preload("Album").Find(&songs).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Zero Rows")
		} else {
			panic(err)
		}
	}
	return
}

func GetArtistAlbums(artist_id int) (albums []Album) {
	err := db.DB.Where("artist_id = ?", artist_id).Preload("Artist").Find(&albums).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Zero Rows")
		} else {
			panic(err)
		}
	}
	return
}

func (artist Artist) Delete() {
	var err error
	// TODO: figure out has many through
	var albums []Album
	err = db.DB.Where("artist_id = ?", artist.ID).Find(&albums).Error
	if err != nil {
		panic(err)
	}
	if len(albums) > 0 {
		err = db.DB.Select("Songs").Where("").Delete(&albums).Error
		if err != nil {
			panic(err)
		}
	}
	err = db.DB.Delete(&artist).Error
	if err != nil {
		panic(err)
	}
}