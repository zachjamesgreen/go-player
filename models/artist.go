package models

import (
	"fmt"

	db "github.com/zachjamesgreen/go-player/database"

	"gorm.io/datatypes"
)

type Artist struct {
	ID        int
	Name      string
	Albums    []Album // on delete casade
	SpotifyId string
	Images    datatypes.JSON
}

func (a Artist) String() string {
	return fmt.Sprintf("Artist<ID: %+v | Name: %+v>\n", a.ID, a.Name)
}

func (artist *Artist) FirstOrCreate() (err error) {
	return db.DB.FirstOrCreate(&artist, Artist{Name: artist.Name}).Error
}

func (artist *Artist) Save() error {
	return db.DB.Save(&artist).Error
}

func GetAllArtists() (artists []Artist, err error) {
	err = db.DB.Find(&artists).Error
	return
}

func GetArtistById(id int) (artist Artist, err error) {
	result := db.DB.Find(&artist, id)
	return artist, result.Error
}

func GetArtistSongsById(artist_id int) (songs []Song, err error) {
	err = db.DB.Where("artist_id = ?", artist_id).Preload("Artist").Preload("Album").Find(&songs).Error
	return
}

func GetArtistAlbumsById(artist_id int) (albums []Album, err error) {
	err = db.DB.Where("artist_id = ?", artist_id).Preload("Artist").Find(&albums).Error
	return
}

func (artist Artist) Delete() (err error) {
	return db.DB.Delete(&artist).Error
}
