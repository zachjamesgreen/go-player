package models

import (
	// "fmt"
	. "github.com/stretchr/testify/assert"
	// "music/database"
	"testing"
	"gorm.io/datatypes"
)

func TestAlbumString(t *testing.T) {
	expected := "Album<ID: 1 | Title: Test | Artist ID: 1 | Spotify ID: test | Spotify Link: test | Images: {\"image\": \"test\"}>\n"
	album := Album{
		ID:          1,
		Title:       "Test",
		ArtistId:    1,
		SpotifyId:   "test",
		SpotifyLink: "test",
		Images:      datatypes.JSON([]byte(`{"image": "test"}`)),
	}
	Equal(t, expected, album.String())
}

// func TestUpsert(t *testing.T) {
// 	db := database.GetTestDB(false)
// 	album := Album{
// 		Title:       "Test",
// 		ArtistId:    1,
// 		SpotifyId:   "test",
// 		SpotifyLink: "test",
// 		Images:      datatypes.JSON([]byte(`{"image": "test"}`)),
// 	}

// 	err := album.Upsert()
// 	NoError(t, err)
// 	result := db.First(&album)
// 	NoError(t, result.Error)
// 	fmt.Print(result)
// 	// Equal(t, album, result.Value)
// }
