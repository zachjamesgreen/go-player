package models

import (
	"fmt"
	"testing"

	. "github.com/stretchr/testify/assert"
	"github.com/zachjamesgreen/go-player/database"
	"gorm.io/datatypes"
)

func TestAlbumString(t *testing.T) {
	expected := "Album<ID: 1 | Title: Test | Artist ID: 1 | Spotify ID: test | Spotify Link: test | Images: {\"image\": \"test\"}>\n"
	album := Album{ID: 1, Title: "Test", ArtistId: 1, SpotifyId: "test", SpotifyLink: "test", Images: datatypes.JSON([]byte(`{"image": "test"}`))}
	Equal(t, expected, album.String())
}

func TestFirstOrCreateAlbum(t *testing.T) {
	db := database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := Artist{Name: "Test"}
	err := artist.FirstOrCreate()
	NoError(t, err)

	album := Album{Title: "Test", Artist: &artist}
	err = album.FirstOrCreate()
	NoError(t, err)

	expected := Album{}
	result := db.First(&expected)
	NoError(t, result.Error)
	Equal(t, album.Title, expected.Title)
	Equal(t, album.ID, expected.ID)
	Equal(t, album.ArtistId, expected.ArtistId)
}

func TestGetAlbums(t *testing.T) {
	db := database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := Artist{Name: "Test"}
	err := artist.FirstOrCreate()
	NoError(t, err)

	var albums []Album
	for i := 0; i < 5; i++ {
		album := Album{Title: fmt.Sprintf("Test %d", i), Artist: &artist}
		err = album.FirstOrCreate()
		NoError(t, err)
		albums = append(albums, album)
	}

	expected := []Album{}
	err = db.Find(&expected).Error
	NoError(t, err)
	for i := 0; i < 5; i++ {
		Equal(t, albums[i].Title, expected[i].Title)
		Equal(t, albums[i].ID, expected[i].ID)
	}
}

func TestSaveAlbum(t *testing.T) {
	database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := Artist{Name: "Test"}
	err := artist.FirstOrCreate()
	NoError(t, err)

	album := Album{Title: "Test", Artist: &artist}
	err = album.FirstOrCreate()
	NoError(t, err)

	album.Title = "Test 2"
	err = album.Save()
	NoError(t, err)
	Equal(t, album.Title, "Test 2")
}

func TestGetAlbum(t *testing.T) {
	db := database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := Artist{Name: "Test"}
	err := artist.FirstOrCreate()
	NoError(t, err)
	var albums []Album

	for i := 0; i < 5; i++ {
		album := Album{Title: fmt.Sprintf("Test %d", i), Artist: &artist}
		err = album.FirstOrCreate()
		NoError(t, err)
		albums = append(albums, album)
	}

	var expected Album
	err = db.Find(&expected, albums[3].ID).Error
	NoError(t, err)

	Equal(t, albums[3].Title, expected.Title)
	Equal(t, albums[3].ID, expected.ID)
}

func TestGetAlbumSongs(t *testing.T) {
	database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := Artist{Name: "Test"}
	err := artist.FirstOrCreate()
	NoError(t, err)
	album := Album{Title: "Test", Artist: &artist}
	err = album.FirstOrCreate()
	NoError(t, err)

	var songs []Song
	for i := 0; i < 5; i++ {
		title := fmt.Sprintf("Test %d", i)
		song := Song{Title: title, Album: &album, Artist: &artist}
		err = song.FirstOrCreate()
		NoError(t, err)
		songs = append(songs, song)
	}

	expected, err := GetAlbumSongs(album.ID)
	NoError(t, err)
	Equal(t, len(expected), 5)
	for i := 0; i < 5; i++ {
		Equal(t, songs[i].Title, expected[i].Title)
		Equal(t, songs[i].ID, expected[i].ID)
		Equal(t, songs[i].ArtistId, expected[i].ArtistId)
	}
}

func TestDeleteAlbum(t *testing.T) {
	db := database.GetTestDB(false)
	defer database.CleanTestDB()
	artist := Artist{Name: "Test"}
	album := Album{Title: "Test", Artist: &artist}
	song := Song{Title: "Test", Album: &album, Artist: &artist}
	err := artist.FirstOrCreate()
	NoError(t, err)
	err = album.FirstOrCreate()
	NoError(t, err)
	err = song.FirstOrCreate()
	NoError(t, err)

	err = album.Delete()
	NoError(t, err)

	var expectedAlbum Album
	var expectedSong Song
	err = db.Find(&expectedAlbum, album.ID).Error
	NoError(t, err)
	Equal(t, Album{}, expectedAlbum)
	err = db.Find(&expectedSong, song.ID).Error
	NoError(t, err)
	Equal(t, Song{}, expectedSong)
}
