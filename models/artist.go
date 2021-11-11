package models

import (
	"fmt"
	db "music/database"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Artist struct {
	ID        int
	Name      string
	Albums    []Album
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
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Zero Rows")
		} else {
			return
		}
	}
	return
}

func GetArtistAlbumsById(artist_id int) (albums []Album, err error) {
	err = db.DB.Where("artist_id = ?", artist_id).Preload("Artist").Find(&albums).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Zero Rows")
		} else {
			return
		}
	}
	return
}

func (artist Artist) Delete() (err error) {
	// TODO: figure out has many through
	var albums []Album
	err = db.DB.Where("artist_id = ?", artist.ID).Find(&albums).Error
	if err != nil {
		return
	}
	if len(albums) > 0 {
		err = db.DB.Select("Songs").Where("").Delete(&albums).Error
		if err != nil {
			return
		}
	}
	err = db.DB.Delete(&artist).Error
	if err != nil {
		return
	}
	return
}
